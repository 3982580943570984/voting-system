package schema

import "entgo.io/ent"

// Candidate holds the schema definition for the Candidate entity.
type Candidate struct {
	ent.Schema
}

// Fields of the Candidate.
func (Candidate) Fields() []ent.Field {
	return nil
}

// Edges of the Candidate.
func (Candidate) Edges() []ent.Edge {
	return nil
}
