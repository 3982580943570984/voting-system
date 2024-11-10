package schema

import (
	"entgo.io/ent"
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
			MaxLen(255).
			NotEmpty(),

		field.String("description").
			Optional(),

		field.Time("start_date").
			Optional().
			SchemaType(map[string]string{
				"postgres": "date",
			}),

		field.Time("end_date").
			Optional().
			SchemaType(map[string]string{
				"postgres": "date",
			}),

		field.Bool("is_active").
			Default(true),
	}
}

// Edges of the Election.
func (Election) Edges() []ent.Edge {
	return nil
}
