package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Organizer holds the schema definition for the Organizer entity.
type Organizer struct {
	ent.Schema
}

// Fields of the Organizer.
func (Organizer) Fields() []ent.Field {
	return nil
}

// Edges of the Organizer.
func (Organizer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("organizer").
			Unique().
			Required(),
	}
}
