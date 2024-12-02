package services

import (
	"context"
	"voting-system/database"
	"voting-system/ent/generated"
	"voting-system/ent/generated/election"
)

// Elections представляет сервис для работы с выборами
// Этот сервис предоставляет методы для создания, получения всех и получения выбора по ID.
type Elections struct {
	DB *generated.ElectionClient
}

// ElectionCreate представляет структуру данных для создания нового выбора
// Используется при создании новых выборов через API.
type ElectionCreate struct {
	// UserID - идентификатор пользователя, которые создает выборы
	// Это обязательное поле.
	UserID int `json:"user_id"`

	// Title — название выборов
	// Это обязательное поле, минимальная длина 8 символов, максимальная длина 64 символа.
	Title string `json:"title" validate:"min=8,max=64"`

	// Description — описание выборов
	// Это обязательное поле, максимальная длина 1000 символов.
	Description string `json:"description" validate:"max=1000"`
}

// ElectionUpdate представляет структуру данных для обновления информации о существующих выборах
// Используется при обновлении данных выборов.
type ElectionUpdate struct {
	// ID — уникальный идентификатор выбора, который необходимо обновить
	// Это обязательное поле для обновления информации о конкретном выборе.
	ID int `json:"id"`

	// Title — новое название выборов
	// Это обязательное поле при обновлении данных.
	Title *string `json:"title,omitempty"`

	// Description — новое описание выборов
	// Это необязательное поле при обновлении данных.
	Description *string `json:"description,omitempty"`
}

func NewElections() *Elections {
	return &Elections{
		DB: database.Client.Election,
	}
}

func (e *Elections) Create(ctx context.Context, ec *ElectionCreate) (*generated.Election, error) {
	return e.DB.Create().
		SetUserID(ec.UserID).
		SetTitle(ec.Title).
		SetDescription(ec.Description).
		Save(ctx)
}

func (e *Elections) GetAll(ctx context.Context) ([]*generated.Election, error) {
	return e.DB.Query().Select(election.Columns...).All(ctx)
}

func (e *Elections) GetById(ctx context.Context, id int) (*generated.Election, error) {
	return e.DB.Query().Select(election.Columns...).Where(election.ID(id)).Only(ctx)
}

func (e *Elections) GetCandidates(ctx context.Context, id int) ([]*generated.Candidate, error) {
	return e.DB.Query().Where(election.ID(id)).QueryCandidates().All(ctx)
}

func (e *Elections) GetSettings(ctx context.Context, id int) (*generated.ElectionSettings, error) {
	return e.DB.Query().Where(election.ID(id)).QuerySettings().Only(ctx)
}

func (e *Elections) Update(ctx context.Context, eu *ElectionUpdate) (*generated.Election, error) {
	return e.DB.UpdateOneID(eu.ID).
		SetNillableTitle(eu.Title).
		SetNillableDescription(eu.Description).
		Save(ctx)
}

func (e *Elections) DeleteById(ctx context.Context, id int) error {
	return e.DB.DeleteOneID(id).Exec(ctx)
}
