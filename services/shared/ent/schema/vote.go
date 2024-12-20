package schema

import (
	"context"
	"errors"

	"shared/ent/generated"
	"shared/ent/generated/candidate"
	"shared/ent/generated/hook"
	"shared/ent/generated/vote"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Vote holds the schema definition for the Vote entity.
type Vote struct {
	ent.Schema
}

// Fields of the Vote.
func (Vote) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("is_active").
			Default(true),
	}
}

func (Vote) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
	}
}

// Edges of the Vote.
func (Vote) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("candidate", Candidate.Type).
			Ref("votes").
			Unique(),

		edge.From("user", User.Type).
			Ref("votes").
			Unique(),
	}
}

// Hooks of the Vote.
func (Vote) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(m ent.Mutator) ent.Mutator {
			return hook.VoteFunc(func(ctx context.Context, vm *generated.VoteMutation) (generated.Value, error) {
				v, err := m.Mutate(ctx, vm)
				if err != nil {
					return nil, err
				}

				vote, ok := v.(*generated.Vote)
				if !ok {
					return nil, errors.New("unexpected type: expected *generated.Vote")
				}

				candidateID, err := vote.QueryCandidate().OnlyID(ctx)
				if err != nil {
					return nil, err
				}

				_, err = vm.Client().Candidate.
					UpdateOneID(candidateID).
					AddVotesCount(1).
					Save(ctx)

				if err != nil {
					return nil, err
				}

				return v, nil
			})
		}, ent.OpCreate),

		hook.On(func(m ent.Mutator) ent.Mutator {
			return hook.VoteFunc(func(ctx context.Context, vm *generated.VoteMutation) (generated.Value, error) {
				voteId, exists := vm.ID()
				if !exists {
					return nil, errors.New("unexpected error: vote has no id")
				}

				candidateId, err := vm.Client().Candidate.Query().
					Where(candidate.HasVotesWith(vote.ID(voteId))).
					OnlyID(ctx)

				if err != nil {
					return nil, err
				}

				_, err = vm.Client().Candidate.
					UpdateOneID(candidateId).
					AddVotesCount(-1).
					Save(ctx)

				if err != nil {
					return nil, err
				}

				return m.Mutate(ctx, vm)
			})
		}, ent.OpDeleteOne),
	}
}
