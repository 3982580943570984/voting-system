package schema

import "entgo.io/ent"

// Complaint holds the schema definition for the Complaint entity.
type Complaint struct {
	ent.Schema
}

// Fields of the Complaint.
func (Complaint) Fields() []ent.Field {
	return nil
}

// Edges of the Complaint.
func (Complaint) Edges() []ent.Edge {
	return nil
}
