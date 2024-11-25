package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
			Positive(),
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
