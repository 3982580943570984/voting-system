package services

import (
	"context"
	"time"
	"voting-system/database"
	ent "voting-system/ent/generated"
	"voting-system/ent/generated/user"

	"golang.org/x/crypto/bcrypt"
)

type UserCreate struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdate struct {
	ID        int        `json:"id,omitempty"`
	Email     *string    `json:"email,omitempty"`
	Password  *string    `json:"password,omitempty"`
	LastLogin *time.Time `json:"last_login,omitempty"`
}

type Users struct {
	DB *ent.UserClient
}

func NewUsers() *Users {
	return &Users{
		DB: database.Users,
	}
}

func (u *Users) Create(ctx context.Context, uc *UserCreate) (*ent.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(uc.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, nil
	}

	user, err := u.DB.
		Create().
		SetEmail(uc.Email).
		SetPassword(string(hashedPassword)).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *Users) GetAll(ctx context.Context) ([]*ent.User, error) {
	return u.DB.
		Query().
		All(ctx)
}

func (u *Users) GetById(ctx context.Context, id int) (*ent.User, error) {
	return u.DB.
		Get(ctx, id)
}

func (u *Users) GetByEmail(ctx context.Context, email string) (*ent.User, error) {
	return u.DB.
		Query().
		Where(user.EmailEQ(email)).
		Only(context.Background())
}

func (u *Users) Update(ctx context.Context, uu *UserUpdate) (*ent.User, error) {
	builder := u.DB.UpdateOneID(uu.ID)

	if uu.Email != nil {
		builder.SetEmail(*uu.Email)
	}

	if uu.Password != nil {
		builder.SetPassword(*uu.Password)
	}

	if uu.LastLogin != nil {
		builder.SetLastLogin(*uu.LastLogin)
	}

	return builder.Save(ctx)
}

func (u *Users) Delete(ctx context.Context, id int) error {
	return u.DB.
		DeleteOneID(id).
		Exec(ctx)
}
