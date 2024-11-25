package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ElectionSettings holds the schema definition for the ElectionSettings entity.
type ElectionSettings struct {
	ent.Schema
}

// Fields of the ElectionSettings.
func (ElectionSettings) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("is_active").
			Default(true),

		field.Bool("is_anonymous").
			Default(false),

		field.Bool("allow_comments").
			Default(true),

		field.Int("max_votes").
			Positive().
			Default(1),

		field.Time("start_date").
			Default(time.Now),

		field.Time("end_date").
			Default(time.Now().AddDate(0, 0, 5)),
	}
}

// Edges of the ElectionSettings.
func (ElectionSettings) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("election", Election.Type).
			Ref("settings").
			Unique().
			Required(),
	}
}
