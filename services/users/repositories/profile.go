package repositories

import (
	"context"
	"shared/database"
	"shared/ent/generated"
	"shared/ent/generated/profile"
	"shared/ent/generated/user"
	"time"
)

type Profiles struct {
	DB *generated.ProfileClient
}

type ProfileUpdate struct {
	UserID      int
	FirstName   *string    `json:"first_name,omitempty"`
	LastName    *string    `json:"last_name,omitempty"`
	Birthdate   *time.Time `json:"birthdate,omitempty"`
	PhoneNumber *string    `json:"phone_number,omitempty"`
	Bio         *string    `json:"bio,omitempty"`
	Address     *string    `json:"address,omitempty"`
	PhotoURL    *string    `json:"photo_url,omitempty"`
}

func NewProfiles() *Profiles {
	return &Profiles{
		DB: database.Client.Profile,
	}
}

func (p *Profiles) GetByUserId(ctx context.Context, id int) (*generated.Profile, error) {
	return p.DB.Query().
		Select(profile.Columns...).
		Where(profile.HasUserWith(user.ID(id))).
		Only(ctx)
}

func (p *Profiles) Update(ctx context.Context, pu *ProfileUpdate) (int, error) {
	return p.DB.Update().
		Where(profile.HasUserWith(user.ID(pu.UserID))).
		SetNillableFirstName(pu.FirstName).
		SetNillableLastName(pu.LastName).
		SetNillableBirthdate(pu.Birthdate).
		SetNillablePhoneNumber(pu.PhoneNumber).
		SetNillableBio(pu.Bio).
		SetNillableAddress(pu.Address).
		SetNillablePhotoURL(pu.PhotoURL).
		Save(ctx)
}
