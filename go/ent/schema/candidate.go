package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Candidate holds the schema definition for the Candidate entity.
type Candidate struct {
	ent.Schema
}

// Fields of the Candidate.
func (Candidate) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(100),

		field.String("description").
			NotEmpty().
			MaxLen(1000),

		field.String("photo_url").
			Optional(),

		field.Int("votes_count").
			Default(0).
			NonNegative(),
	}
}

// Mixins of the Candidate.
func (Candidate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Candidate.
func (Candidate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("election", Election.Type).
			Ref("candidates").
			Unique(),

		edge.To("votes", Vote.Type),
	}
}
