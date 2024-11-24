package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// TODO: custom validator for email value
		field.String("email").
			NotEmpty().
			Unique(),

		field.String("password").
			MinLen(8).
			MaxLen(32).
			Sensitive(),

		field.Time("created_at").
			Immutable().
			Default(time.Now),

		field.Time("last_login").
			Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
