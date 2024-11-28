package schema

import (
	"context"
	"errors"
	"voting-system/ent/generated"
	"voting-system/ent/generated/hook"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Election holds the schema definition for the Election entity.
type Election struct {
	ent.Schema
}

// Fields of the Election.
func (Election) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			MinLen(8).
			MaxLen(64),

		field.String("description").
			NotEmpty().
			MaxLen(1000),
	}
}

// Edges of the Election.
func (Election) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("elections").
			Unique(),

		edge.From("tags", Tag.Type).
			Ref("elections"),

		edge.To("comments", Comment.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),

		edge.To("candidates", Candidate.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),

		edge.To("settings", ElectionSettings.Type).
			Unique().
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

// Hooks of the Election.
func (Election) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.ElectionFunc(func(ctx context.Context, em *generated.ElectionMutation) (generated.Value, error) {
				v, err := next.Mutate(ctx, em)
				if err != nil {
					return nil, err
				}

				election, ok := v.(*generated.Election)
				if !ok {
					return nil, errors.New("unexpected type: expected *ent.Election")
				}

				_, err = em.Client().ElectionSettings.
					Create().
					SetElection(election).
					Save(ctx)

				if err != nil {
					return nil, err
				}

				return v, nil
			})
		}, ent.OpCreate),
	}
}
