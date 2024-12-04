package services

import (
	"context"
	"voting-system/database"
	"voting-system/ent/generated"
	"voting-system/ent/generated/candidate"
	"voting-system/ent/generated/election"
)

type Candidates struct {
	DB *generated.CandidateClient
}

type CandidateCreate struct {
	ElectionId  int
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CandidateUpdate struct {
	ID          int     `json:"id"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	PhotoURL    *string `json:"photo_url,omitempty"`
}

type CandidateDelete struct {
	ID         int `json:"id"`
	ElectionId int
}

func NewCandidates() *Candidates {
	return &Candidates{
		DB: database.Client.Candidate,
	}
}

func (c *Candidates) Create(ctx context.Context, cc *CandidateCreate) (*generated.Candidate, error) {
	return c.DB.Create().
		SetElectionID(cc.ElectionId).
		SetName(cc.Name).
		SetDescription(cc.Description).
		Save(ctx)
}

func (c *Candidates) GetAll(ctx context.Context) ([]*generated.Candidate, error) {
	return c.DB.Query().Select(
		candidate.FieldName,
		candidate.FieldDescription,
		candidate.FieldPhotoURL,
		candidate.FieldVotesCount,
	).All(ctx)
}

func (c *Candidates) GetById(ctx context.Context, id int) (*generated.Candidate, error) {
	return c.DB.Query().Select(
		candidate.FieldName,
		candidate.FieldDescription,
		candidate.FieldPhotoURL,
		candidate.FieldVotesCount,
	).Where(candidate.ID(id)).Only(ctx)
}

func (c *Candidates) GetByElectionId(ctx context.Context, id int) ([]*generated.Candidate, error) {
	return c.DB.Query().Select(
		candidate.FieldName,
		candidate.FieldDescription,
		candidate.FieldPhotoURL,
		candidate.FieldVotesCount,
	).Where(candidate.HasElectionWith(election.ID(id))).All(ctx)
}

func (c *Candidates) Update(ctx context.Context, cu *CandidateUpdate) (*generated.Candidate, error) {
	return c.DB.UpdateOneID(cu.ID).
		SetNillableName(cu.Name).
		SetNillableDescription(cu.Description).
		SetNillablePhotoURL(cu.PhotoURL).
		Save(ctx)
}

func (c *Candidates) Delete(ctx context.Context, cd *CandidateDelete) error {
	return c.DB.DeleteOneID(cd.ID).
		Where(candidate.HasElectionWith(election.ID(cd.ElectionId))).
		Exec(ctx)
}
