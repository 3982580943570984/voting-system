package schema

import "entgo.io/ent"

// FAQ holds the schema definition for the FAQ entity.
type FAQ struct {
	ent.Schema
}

// Fields of the FAQ.
func (FAQ) Fields() []ent.Field {
	return nil
}

// Edges of the FAQ.
func (FAQ) Edges() []ent.Edge {
	return nil
}
