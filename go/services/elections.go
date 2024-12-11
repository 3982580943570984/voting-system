package services

import (
	"context"
	"fmt"
	"log"
	"voting-system/database"
	"voting-system/ent/generated"
	"voting-system/ent/generated/election"
	filters "voting-system/ent/generated/electionfilters"
	"voting-system/ent/generated/electionsettings"
	"voting-system/ent/generated/user"
	"voting-system/scheduler"

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

	_, err = tx.Candidate.MapCreateBulk(ec.Candidates, func(cc *generated.CandidateCreate, i int) {
		candidate := ec.Candidates[i]
		cc.SetElection(election).
			SetName(candidate.Name).
			SetDescription(candidate.Description)
	}).Save(ctx)
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
		Select(election.Columns...).
		All(ctx)
}

func (e *Elections) GetAllWithDuration(ctx context.Context) ([]*generated.Election, error) {
	return e.DB.Query().
		Select(election.Columns...).
		WithSettings(func(esq *generated.ElectionSettingsQuery) {
			esq.Select(electionsettings.FieldDuration)
		}).
		All(ctx)
}

func (e *Elections) GetAllFiltered(ctx context.Context, id int) ([]*generated.Election, error) {
	profile, err := NewProfiles().GetByUserId(ctx, id)
	if err != nil {
		return nil, err
	}

	query := e.DB.Query().WithFilters()

	if profile.FirstName == "" {
		query = query.Where(
			election.HasFiltersWith(
				filters.HasFirstName(false),
			),
		)
	}

	if profile.LastName == "" {
		query = query.Where(
			election.HasFiltersWith(
				filters.HasLastName(false),
			),
		)
	}

	if profile.Birthdate.IsZero() {
		query = query.Where(
			election.HasFiltersWith(
				filters.HasBirthdate(false),
			),
		)
	}

	if profile.PhoneNumber == "" {
		query = query.Where(
			election.HasFiltersWith(
				filters.HasPhoneNumber(false),
			),
		)
	}

	if profile.Bio == "" {
		query = query.Where(
			election.HasFiltersWith(
				filters.HasBio(false),
			),
		)
	}

	if profile.Address == "" {
		query = query.Where(
			election.HasFiltersWith(
				filters.HasAddress(false),
			),
		)
	}

	if profile.PhotoURL == "" {
		query = query.Where(
			election.HasFiltersWith(
				filters.HasPhotoURL(false),
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
		Select(election.Columns...).
		Where(election.ID(id)).
		Only(ctx)
}

func (e *Elections) GetByUserId(ctx context.Context, id int) ([]*generated.Election, error) {
	return e.DB.Query().
		Select(election.Columns...).
		Where(election.HasUserWith(user.ID(id))).
		All(ctx)
}

func (e *Elections) GetCandidates(ctx context.Context, id int) ([]*generated.Candidate, error) {
	return e.DB.Query().Where(election.ID(id)).QueryCandidates().All(ctx)
}

func (e *Elections) GetSettings(ctx context.Context, id int) (*generated.ElectionSettings, error) {
	return e.DB.Query().Where(election.ID(id)).QuerySettings().Only(ctx)
}

func (e *Elections) GetFilters(ctx context.Context, id int) (*generated.ElectionFilters, error) {
	return e.DB.Query().Where(election.ID(id)).QueryFilters().Only(ctx)
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
