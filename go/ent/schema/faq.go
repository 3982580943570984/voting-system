package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// FAQ holds the schema definition for the FAQ entity.
type FAQ struct {
	ent.Schema
}

// Fields of the FAQ.
func (FAQ) Fields() []ent.Field {
	return []ent.Field{
		field.String("question"),

		field.String("answer"),
	}
}

// Edges of the FAQ.
func (FAQ) Edges() []ent.Edge {
	return []ent.Edge{}
}
