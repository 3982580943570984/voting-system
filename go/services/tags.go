package services

import (
	"context"
	"voting-system/database"
	"voting-system/ent/generated"
	"voting-system/ent/generated/tag"
)

type Tags struct {
	DB *generated.TagClient
}

type TagCreate struct {
	ElectionId int    `json:"election_id"`
	Name       string `json:"name"`
}

type TagRemove struct {
	ElectionId int    `json:"election_id"`
	Name       string `json:"name"`
}

func NewTags() *Tags {
	return &Tags{
		DB: database.Client.Tag,
	}
}

func (t *Tags) Create(ctx context.Context, tc *TagCreate) (*generated.Tag, error) {
	return t.DB.Create().
		SetName(tc.Name).
		AddElectionIDs(tc.ElectionId).
		Save(ctx)
}

func (t *Tags) RemoveFromElectionById(ctx context.Context, tr *TagRemove) (int, error) {
	return t.DB.Update().
		Where(tag.Name(tr.Name)).
		RemoveElectionIDs(tr.ElectionId).
		Save(ctx)
}
