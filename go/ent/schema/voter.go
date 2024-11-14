package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Voter holds the schema definition for the Voter entity.
type Voter struct {
	ent.Schema
}

// Fields of the Voter.
func (Voter) Fields() []ent.Field {
	return nil
}

// Edges of the Voter.
func (Voter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("voter").
			Unique().
			Required(),
	}
}
