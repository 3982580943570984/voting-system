package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Profile holds the schema definition for the Profile entity.
type Profile struct {
	ent.Schema
}

// Fields of the Profile.
func (Profile) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").
			Optional().
			NotEmpty().
			MaxLen(50),

		field.String("last_name").
			Optional().
			NotEmpty().
			MaxLen(50),

		field.Time("birthdate").
			Optional(),

		field.String("phone_number").
			Optional().
			Match(regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)),

		field.String("bio").
			Optional().
			MaxLen(500),

		field.String("address").
			Optional().
			MaxLen(255),

		field.String("photo_url").
			Optional(),
	}
}

// Edges of the Profile.
func (Profile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("profile").
			Unique().
			Required(),
	}
}
