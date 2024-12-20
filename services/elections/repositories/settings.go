package repositories

import (
	"context"
	"shared/database"
	"shared/ent/generated"
	"shared/ent/generated/election"
	"shared/ent/generated/electionsettings"
	"time"
)

type ElectionSettings struct {
	DB *generated.ElectionSettingsClient
}

type ElectionSettingsUpdate struct {
	ElectionID    int
	IsActive      *bool      `json:"is_active,omitempty"`
	IsAnonymous   *bool      `json:"is_anonymous,omitempty"`
	AllowComments *bool      `json:"allow_comments,omitempty"`
	MaxVotes      *int       `json:"max_votes,omitempty"`
	Duration      *time.Time `json:"duration,omitempty"`
}

func NewElectionSettings() *ElectionSettings {
	return &ElectionSettings{
		DB: database.Client.ElectionSettings,
	}
}

func (ef *ElectionSettings) Update(ctx context.Context, esu *ElectionSettingsUpdate) error {
	return ef.DB.Update().
		Where(electionsettings.HasElectionWith(election.ID(esu.ElectionID))).
		SetNillableIsActive(esu.IsActive).
		SetNillableIsAnonymous(esu.IsAnonymous).
		SetNillableAllowComments(esu.AllowComments).
		SetNillableMaxVotes(esu.MaxVotes).
		SetNillableDuration(esu.Duration).
		Exec(ctx)
}
