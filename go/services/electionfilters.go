package services

import (
	"context"
	"voting-system/database"
	"voting-system/ent/generated"
	"voting-system/ent/generated/election"
	"voting-system/ent/generated/electionfilters"
)

type ElectionFilters struct {
	DB *generated.ElectionFiltersClient
}

type ElectionFiltersUpdate struct {
	ElectionID     int
	HasFirstName   *bool `json:"has_first_name,omitempty"`
	HasLastName    *bool `json:"has_last_name,omitempty"`
	HasBirthdate   *bool `json:"has_birthdate,omitempty"`
	HasPhoneNumber *bool `json:"has_phone_number,omitempty"`
	HasBio         *bool `json:"has_bio,omitempty"`
	HasAddress     *bool `json:"has_address,omitempty"`
	HasPhotoURL    *bool `json:"has_photo_url,omitempty"`
}

func NewElectionFilters() *ElectionFilters {
	return &ElectionFilters{
		DB: database.Client.ElectionFilters,
	}
}

func (ef *ElectionFilters) Update(ctx context.Context, efu *ElectionFiltersUpdate) error {
	return ef.DB.Update().
		Where(electionfilters.HasElectionWith(election.ID(efu.ElectionID))).
		SetNillableHasFirstName(efu.HasFirstName).
		SetNillableHasLastName(efu.HasLastName).
		SetNillableHasBirthdate(efu.HasBirthdate).
		SetNillableHasPhoneNumber(efu.HasPhoneNumber).
		SetNillableHasBio(efu.HasBio).
		SetNillableHasAddress(efu.HasAddress).
		SetNillableHasPhotoURL(efu.HasPhotoURL).
		Exec(ctx)
}
