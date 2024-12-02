package services

import (
	"context"
	"voting-system/database"
	"voting-system/ent/generated"
	"voting-system/ent/generated/candidate"
	"voting-system/ent/generated/user"
	"voting-system/ent/generated/vote"
)

type Votes struct {
	DB *generated.VoteClient
}

type VoteCreate struct {
	UserId      int `json:"user_id"`
	CandidateId int `json:"candidate_id"`
}

type VoteDelete struct {
	UserId      int `json:"user_id"`
	CandidateId int `json:"candidate_id"`
}

func NewVotes() *Votes {
	return &Votes{
		DB: database.Client.Vote,
	}
}

func (v *Votes) Create(ctx context.Context, vc *VoteCreate) (*generated.Vote, error) {
	return v.DB.Create().
		SetUserID(vc.UserId).
		SetCandidateID(vc.CandidateId).
		Save(ctx)
}

func (v *Votes) Delete(ctx context.Context, vd *VoteDelete) error {
	voteId, err := v.DB.Query().Where(
		vote.HasUserWith(user.ID(vd.UserId)),
		vote.HasCandidateWith(candidate.ID(vd.CandidateId)),
	).OnlyID(ctx)

	if err != nil {
		return err
	}

	return v.DB.DeleteOneID(voteId).Exec(ctx)
}
