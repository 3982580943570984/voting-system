package services

import (
	"context"
	"voting-system/database"
	ent "voting-system/ent/generated"
)

type Elections struct {
	DB *ent.ElectionClient
}

type ElectionCreate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ElectionUpdate struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewElections() *Elections {
	return &Elections{
		DB: database.Client.Election,
	}
}

func (e *Elections) Create(ctx context.Context, ec *ElectionCreate) (*ent.Election, error) {
	election, err := e.DB.Create().
		SetTitle(ec.Title).
		SetDescription(ec.Description).
		Save(ctx)

	if err != nil {
		return nil, nil
	}

	return election, nil
}

func (e *Elections) GetAll(ctx context.Context) ([]*ent.Election, error) {
	return e.DB.Query().All(ctx)
}

func (e *Elections) GetById(ctx context.Context, id int) (*ent.Election, error) {
	return e.DB.Get(ctx, id)
}
