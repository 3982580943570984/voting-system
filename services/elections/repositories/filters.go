package repositories

import (
	"context"
	"shared/database"
	"shared/ent/generated"
	"shared/ent/generated/election"
	"shared/ent/generated/electionfilters"
)

type Filters struct {
	DB *generated.ElectionFiltersClient
}

type FiltersUpdate struct {
	ElectionID     int
	HasFirstName   *bool `json:"has_first_name,omitempty"`
	HasLastName    *bool `json:"has_last_name,omitempty"`
	HasBirthdate   *bool `json:"has_birthdate,omitempty"`
	HasPhoneNumber *bool `json:"has_phone_number,omitempty"`
	HasBio         *bool `json:"has_bio,omitempty"`
	HasAddress     *bool `json:"has_address,omitempty"`
	HasPhotoURL    *bool `json:"has_photo_url,omitempty"`
}

func NewFilters() *Filters {
	return &Filters{
		DB: database.Client.ElectionFilters,
	}
}

func (f *Filters) Update(ctx context.Context, efu *FiltersUpdate) error {
	return f.DB.Update().
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
