package repositories

import (
	"context"
	"shared/database"
	"shared/ent/generated"
	"shared/ent/generated/user"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	DB *generated.UserClient
}

type UserCreate struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsOrganizer bool   `json:"is_organizer"`
}

type UserUpdate struct {
	ID          int        `json:"id"`
	Email       *string    `json:"email,omitempty"`
	Password    *string    `json:"password,omitempty"`
	LastLogin   *time.Time `json:"last_login,omitempty"`
	IsActive    *bool      `json:"is_active,omitempty"`
	IsOrganizer *bool      `json:"is_organizer,omitempty"`
}

func NewUsers() *Users {
	return &Users{
		DB: database.Client.User,
	}
}

func (u *Users) Create(ctx context.Context, uc *UserCreate) (*generated.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(uc.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, nil
	}

	return u.DB.
		Create().
		SetEmail(uc.Email).
		SetPassword(string(hashedPassword)).
		SetIsOrganizer(uc.IsOrganizer).
		Save(ctx)
}

func (u *Users) GetAll(ctx context.Context) ([]*generated.User, error) {
	return u.DB.Query().All(ctx)
}

func (u *Users) GetById(ctx context.Context, id int) (*generated.User, error) {
	return u.DB.Get(ctx, id)
}

func (u *Users) GetByEmail(ctx context.Context, email string) (*generated.User, error) {
	return u.DB.Query().Where(user.Email(email)).Only(ctx)
}

func (u *Users) Update(ctx context.Context, uu *UserUpdate) (*generated.User, error) {
	return u.DB.UpdateOneID(uu.ID).
		SetNillableEmail(uu.Email).
		SetNillablePassword(uu.Password).
		SetNillableIsActive(uu.IsActive).
		SetNillableIsOrganizer(uu.IsOrganizer).
		Save(ctx)
}

func (u *Users) Delete(ctx context.Context, id int) error {
	return u.DB.DeleteOneID(id).Exec(ctx)
}
