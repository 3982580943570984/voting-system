package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ElectionFilters holds the schema definition for the ElectionFilters entity.
type ElectionFilters struct {
	ent.Schema
}

// Fields of the ElectionFilters.
func (ElectionFilters) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("has_first_name").
			Default(false),

		field.Bool("has_last_name").
			Default(false),

		field.Bool("has_birthdate").
			Default(false),

		field.Bool("has_phone_number").
			Default(false),

		field.Bool("has_bio").
			Default(false),

		field.Bool("has_address").
			Default(false),

		field.Bool("has_photo_url").
			Default(false),
	}
}

// Edges of the ElectionFilters.
func (ElectionFilters) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("election", Election.Type).
			Ref("filters").
			Unique().
			Required(),
	}
}
