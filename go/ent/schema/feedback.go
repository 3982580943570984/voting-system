package schema

import "entgo.io/ent"

// Feedback holds the schema definition for the Feedback entity.
type Feedback struct {
	ent.Schema
}

// Fields of the Feedback.
func (Feedback) Fields() []ent.Field {
	return nil
}

// Edges of the Feedback.
func (Feedback) Edges() []ent.Edge {
	return nil
}
