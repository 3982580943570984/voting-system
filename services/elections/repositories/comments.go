package repositories

import (
	"context"
	"shared/database"
	"shared/ent/generated"
	"shared/ent/generated/comment"
	"shared/ent/generated/election"
)

type Comments struct {
	DB *generated.CommentClient
}

type CommentCreate struct {
	UserId     int
	ElectionId int
	ParentId   *int   `json:"parent_id,omitempty"`
	Contents   string `json:"contents"`
}

func NewComments() *Comments {
	return &Comments{
		DB: database.Client.Comment,
	}
}

func (c *Comments) Create(ctx context.Context, cc *CommentCreate) (*generated.Comment, error) {
	return c.DB.Create().
		SetContents(cc.Contents).
		SetUserID(cc.UserId).
		SetElectionID(cc.ElectionId).
		SetNillableParentID(cc.ParentId).
		Save(ctx)
}

func (c *Comments) GetAllByElectionId(ctx context.Context, electionId int) ([]*generated.Comment, error) {
	return c.DB.Query().
		Where(comment.HasElectionWith(election.ID(electionId))).
		WithUser().
		WithParent().
		All(ctx)
}
