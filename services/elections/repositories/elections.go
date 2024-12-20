package repositories

import (
	"context"
	"elections/scheduler"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"shared/database"
	"shared/ent/generated"
	genele "shared/ent/generated/election"
	genfil "shared/ent/generated/electionfilters"
	"shared/ent/generated/electionsettings"
	"shared/ent/generated/user"
	"strconv"
	"time"

	"github.com/go-co-op/gocron/v2"
)

type Elections struct {
	DB *generated.ElectionClient
}

type ElectionCreate struct {
	UserID      int    `json:"user_id"`
	Title       string `json:"title" validate:"min=8,max=64"`
	Description string `json:"description" validate:"max=1000"`
	Candidates  []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"candidates" validate:"required,min=1"`
	Settings struct {
		IsActive      *bool      `json:"is_active,omitempty"`
		IsAnonymous   *bool      `json:"is_anonymous,omitempty"`
		AllowComments *bool      `json:"allow_comments,omitempty"`
		MaxVotes      *int       `json:"max_votes,omitempty"`
		Duration      *time.Time `json:"duration,omitempty"`
	}
	Filters struct {
		HasFirstName   *bool `json:"has_first_name,omitempty"`
		HasLastName    *bool `json:"has_last_name,omitempty"`
		HasBirthdate   *bool `json:"has_birthdate,omitempty"`
		HasPhoneNumber *bool `json:"has_phone_number,omitempty"`
		HasBio         *bool `json:"has_bio,omitempty"`
		HasAddress     *bool `json:"has_address,omitempty"`
		HasPhotoURL    *bool `json:"has_photo_url,omitempty"`
	} `json:"filters"`
	Tags []string `json:"tags,omitempty"`
}

type ElectionUpdate struct {
	ID          int     `json:"id"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Completed   *bool   `json:"completed,omitempty"`
}

type AgeGroup struct {
	Under12        int `json:"under_12"`
	Between12And18 int `json:"between_12_and_18"`
	Between18And25 int `json:"between_18_and_25"`
	Between25And40 int `json:"between_25_and_40"`
	Between40And60 int `json:"between_40_and_60"`
	Over60         int `json:"over_60"`
	NoAge          int `json:"no_age"`
}

type CommentInfo struct {
	TotalComments    int     `json:"total_comments"`
	AvgCommentLength int     `json:"avg_comment_length"`
	CommentsPerUser  float64 `json:"comments_per_user"`
}

type Statistics struct {
	AgeGroups   AgeGroup    `json:"age_groups"`
	CommentInfo CommentInfo `json:"comment_info"`
}

func NewElections() *Elections {
	return &Elections{
		DB: database.Client.Election,
	}
}

func (e *Elections) Create(ctx context.Context, ec *ElectionCreate) (*generated.Election, error) {
	rollback := func(tx *generated.Tx, err error) error {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return err
	}

	tx, err := database.Client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	election, err := tx.Election.Create().
		SetUserID(ec.UserID).
		SetTitle(ec.Title).
		SetDescription(ec.Description).
		Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	settings, err := election.QuerySettings().Only(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	err = settings.Update().
		SetNillableIsActive(ec.Settings.IsActive).
		SetNillableIsAnonymous(ec.Settings.IsAnonymous).
		SetNillableAllowComments(ec.Settings.AllowComments).
		SetNillableMaxVotes(ec.Settings.MaxVotes).
		Exec(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	filters, err := election.QueryFilters().Only(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	err = filters.Update().
		SetNillableHasFirstName(ec.Filters.HasFirstName).
		SetNillableHasLastName(ec.Filters.HasLastName).
		SetNillableHasBirthdate(ec.Filters.HasBirthdate).
		SetNillableHasPhoneNumber(ec.Filters.HasPhoneNumber).
		SetNillableHasBio(ec.Filters.HasBio).
		SetNillableHasAddress(ec.Filters.HasAddress).
		SetNillableHasPhotoURL(ec.Filters.HasPhotoURL).
		Exec(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	_, err = tx.Candidate.MapCreateBulk(ec.Candidates, func(cc *generated.CandidateCreate, i int) {
		candidate := ec.Candidates[i]
		cc.SetElection(election).
			SetName(candidate.Name).
			SetDescription(candidate.Description)
	}).Save(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	_, err = scheduler.Scheduler.NewJob(
		gocron.OneTimeJob(
			gocron.OneTimeJobStartDateTime(
				settings.Duration,
			),
		),
		gocron.NewTask(
			func() {
				_, err := NewElections().UpdateCompleted(
					context.Background(),
					election.ID,
					true,
				)
				if err != nil {
					log.Fatalf("Error during election update: %v", err)
				}
			},
		),
	)
	if err != nil {
		return nil, rollback(tx, err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, rollback(tx, err)
	}

	tagsService := NewTags()
	for _, tag := range ec.Tags {
		_, err := tagsService.Create(ctx, &TagCreate{
			ElectionId: election.ID,
			Name:       tag,
		})
		if err != nil {
			return nil, rollback(tx, err)
		}
	}

	return election, nil
}

func (e *Elections) GetAll(ctx context.Context) ([]*generated.Election, error) {
	return e.DB.Query().
		Select(genele.Columns...).
		All(ctx)
}

func (e *Elections) GetAllWithDuration(ctx context.Context) ([]*generated.Election, error) {
	return e.DB.Query().
		Select(genele.Columns...).
		WithSettings(func(esq *generated.ElectionSettingsQuery) {
			esq.Select(
				electionsettings.FieldCreateTime,
				electionsettings.FieldDuration,
			)
		}).
		All(ctx)
}

func (e *Elections) GetAllFiltered(ctx context.Context, id int) ([]*generated.Election, error) {
	profile, err := e.fetchUserProfile(ctx, id)
	if err != nil {
		return nil, err
	}

	query := e.DB.Query().WithFilters()

	if profile.FirstName == "" {
		query = query.Where(
			genele.HasFiltersWith(
				genfil.HasFirstName(false),
			),
		)
	}

	if profile.LastName == "" {
		query = query.Where(
			genele.HasFiltersWith(
				genfil.HasLastName(false),
			),
		)
	}

	if profile.Birthdate.IsZero() {
		query = query.Where(
			genele.HasFiltersWith(
				genfil.HasBirthdate(false),
			),
		)
	}

	if profile.PhoneNumber == "" {
		query = query.Where(
			genele.HasFiltersWith(
				genfil.HasPhoneNumber(false),
			),
		)
	}

	if profile.Bio == "" {
		query = query.Where(
			genele.HasFiltersWith(
				genfil.HasBio(false),
			),
		)
	}

	if profile.Address == "" {
		query = query.Where(
			genele.HasFiltersWith(
				genfil.HasAddress(false),
			),
		)
	}

	if profile.PhotoURL == "" {
		query = query.Where(
			genele.HasFiltersWith(
				genfil.HasPhotoURL(false),
			),
		)
	}

	elections, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	return elections, nil
}

func (e *Elections) GetById(ctx context.Context, id int) (*generated.Election, error) {
	return e.DB.Query().
		Select(genele.Columns...).
		Where(genele.ID(id)).
		Only(ctx)
}

func (e *Elections) GetByUserId(ctx context.Context, id int) ([]*generated.Election, error) {
	return e.DB.Query().
		Select(genele.Columns...).
		Where(genele.HasUserWith(user.ID(id))).
		All(ctx)
}

func (e *Elections) GetCandidates(ctx context.Context, id int) ([]*generated.Candidate, error) {
	return e.DB.Query().Where(genele.ID(id)).QueryCandidates().All(ctx)
}

func (e *Elections) GetStatistics(ctx context.Context, id int) (any, error) {
	calculateAge := func(birthdate time.Time) int {
		now := time.Now()
		age := now.Year() - birthdate.Year()
		if now.YearDay() < birthdate.YearDay() {
			age--
		}
		return age
	}

	elections, err := e.GetByUserId(ctx, id)
	if err != nil {
		return nil, err
	}

	var ageGroups AgeGroup
	commentInfo := CommentInfo{}
	var totalCommentLength int

	for _, election := range elections {
		votes, err := election.QueryCandidates().QueryVotes().All(ctx)
		if err != nil {
			return nil, err
		}

		var totalUsers int
		for _, vote := range votes {
			user, err := vote.QueryUser().Only(ctx)
			if err != nil {
				return nil, err
			}

			profile, err := user.QueryProfile().Only(ctx)
			if err != nil {
				return nil, err
			}

			if profile.Birthdate.IsZero() {
				ageGroups.NoAge++
			} else {
				age := calculateAge(profile.Birthdate)
				switch {
				case age < 12:
					ageGroups.Under12++
				case age >= 12 && age < 18:
					ageGroups.Between12And18++
				case age >= 18 && age < 25:
					ageGroups.Between18And25++
				case age >= 25 && age < 40:
					ageGroups.Between25And40++
				case age >= 40 && age < 60:
					ageGroups.Between40And60++
				default:
					ageGroups.Over60++
				}
			}
			totalUsers++
		}

		comments, err := election.QueryComments().All(ctx)
		if err != nil {
			return nil, err
		}

		commentInfo.TotalComments += len(comments)

		for _, comment := range comments {
			totalCommentLength += len(comment.Contents)
		}

		if commentInfo.TotalComments > 0 {
			commentInfo.AvgCommentLength = totalCommentLength / commentInfo.TotalComments
		}

		if totalUsers > 0 && commentInfo.TotalComments > 0 {
			commentInfo.CommentsPerUser = float64(commentInfo.TotalComments) / float64(totalUsers)
		}
	}

	statistics := Statistics{
		AgeGroups:   ageGroups,
		CommentInfo: commentInfo,
	}

	return statistics, nil
}

func (e *Elections) GetSettings(ctx context.Context, id int) (*generated.ElectionSettings, error) {
	return e.DB.Query().Where(genele.ID(id)).QuerySettings().Only(ctx)
}

func (e *Elections) GetFilters(ctx context.Context, id int) (*generated.ElectionFilters, error) {
	return e.DB.Query().Where(genele.ID(id)).QueryFilters().Only(ctx)
}

func (e *Elections) Update(ctx context.Context, eu *ElectionUpdate) (*generated.Election, error) {
	return e.DB.UpdateOneID(eu.ID).
		SetNillableTitle(eu.Title).
		SetNillableDescription(eu.Description).
		SetNillableCompleted(eu.Completed).
		Save(ctx)
}

func (e *Elections) UpdateCompleted(ctx context.Context, id int, completed bool) (*generated.Election, error) {
	return e.DB.UpdateOneID(id).
		SetCompleted(completed).
		Save(ctx)
}

func (e *Elections) DeleteById(ctx context.Context, id int) error {
	return e.DB.DeleteOneID(id).Exec(ctx)
}

func (e *Elections) fetchUserProfile(ctx context.Context, userID int) (*generated.Profile, error) {
	usersURL, err := url.JoinPath("http://localhost:3001", strconv.Itoa(userID), "profile")
	if err != nil {
		return nil, fmt.Errorf("failed to construct users service url: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, usersURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}
	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform http request to users service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("user service request failed with status code %d, and body %s", resp.StatusCode, string(body))
	}

	var profile generated.Profile
	if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
		return nil, fmt.Errorf("failed to decode profile data from user service response: %w", err)
	}
	return &profile, nil
}
