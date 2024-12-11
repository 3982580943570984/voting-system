package services

import (
	"context"
	"errors"
	"voting-system/database"
	"voting-system/ent/generated"
	"voting-system/ent/generated/election"
	"voting-system/ent/generated/tag"
)

type Tags struct {
	DB *generated.TagClient
}

type TagCreate struct {
	ElectionId int
	Name       string `json:"name"`
}

type TagsUpdate struct {
	ElectionId int
	Names      []string `json:"names"`
}

type TagRemove struct {
	ElectionId int
	Name       string `json:"name"`
}

func NewTags() *Tags {
	return &Tags{
		DB: database.Client.Tag,
	}
}

func (t *Tags) Create(ctx context.Context, tc *TagCreate) (*generated.Tag, error) {
	exists, err := t.DB.Query().
		Where(tag.Name(tc.Name)).
		Exist(ctx)
	if err != nil {
		return nil, err
	}

	if exists {
		err = t.DB.Update().
			Where(tag.Name(tc.Name)).
			AddElectionIDs(tc.ElectionId).
			Exec(ctx)
		if err != nil {
			return nil, err
		}

		return t.DB.Query().
			Where(tag.Name(tc.Name)).
			Only(ctx)
	}

	return t.DB.Create().
		SetName(tc.Name).
		AddElectionIDs(tc.ElectionId).
		Save(ctx)
}

func (t *Tags) GetAll(ctx context.Context) ([]*generated.Tag, error) {
	return t.DB.Query().All(ctx)
}

func (t *Tags) GetByElectionId(ctx context.Context, id int) ([]*generated.Tag, error) {
	return t.DB.Query().
		Where(tag.HasElectionsWith(election.ID(id))).
		All(ctx)
}

func (t *Tags) UpdateByElectionId(ctx context.Context, tu *TagsUpdate) error {
	existingTags, err := t.GetByElectionId(ctx, tu.ElectionId)
	if err != nil {
		return err
	}

	for _, existindTag := range existingTags {
		err := t.DB.Update().
			Where(tag.Name(existindTag.Name)).
			RemoveElectionIDs(tu.ElectionId).
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	for _, name := range tu.Names {
		existingTag, err := t.DB.Query().
			Where(tag.Name(name)).
			Only(ctx)

		if err != nil {
			if errors.Is(err, &generated.NotFoundError{}) {
				existingTag, err = t.DB.Create().
					SetName(name).
					Save(ctx)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}

		_, err = t.DB.UpdateOne(existingTag).
			AddElectionIDs(tu.ElectionId).
			Save(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *Tags) RemoveFromElectionById(ctx context.Context, tr *TagRemove) (int, error) {
	return t.DB.Update().
		Where(tag.Name(tr.Name)).
		RemoveElectionIDs(tr.ElectionId).
		Save(ctx)
}
