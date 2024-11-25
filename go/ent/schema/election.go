package schema

import (
	"entgo.io/ent"
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

		edge.To("comments", Comment.Type),

		edge.To("candidates", Candidate.Type),

		edge.To("settings", ElectionSettings.Type).
			Unique(),
	}
}
