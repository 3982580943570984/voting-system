package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// TODO: add email validation
		field.String("email").
			Unique().
			Validate(func(s string) error {
				return nil
			}),

		field.String("password").
			Sensitive(),

		field.Bool("is_active").
			Default(true),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", Profile.Type).
			Unique().
			Annotations(entsql.OnDelete(entsql.Cascade)),

		edge.To("voter", Voter.Type).
			Unique().
			Annotations(entsql.OnDelete(entsql.Cascade)),

		edge.To("organizer", Organizer.Type).
			Unique().
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email").
			Unique(),
	}
}
