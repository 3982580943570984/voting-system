package schema

import (
	"context"
	"errors"
	"net/mail"
	"time"

	"shared/ent/generated"
	"shared/ent/generated/hook"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").
			NotEmpty().
			Unique().
			Validate(func(s string) error {
				_, err := mail.ParseAddress(s)
				if err != nil {
					return errors.New("invalid email format")
				}
				return nil
			}),

		field.String("password").
			NotEmpty().
			Sensitive(),

		field.Time("last_login").
			Default(time.Now).
			UpdateDefault(time.Now),

		field.Bool("is_active").
			Default(true),

		field.Bool("is_organizer").
			Default(false),
	}
}

// Mixins of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", Profile.Type).
			Unique().
			Annotations(entsql.OnDelete(entsql.Cascade)),

		edge.To("comments", Comment.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),

		edge.To("elections", Election.Type),

		edge.To("votes", Vote.Type),
	}
}

// Hooks of the User.
func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(m ent.Mutator) ent.Mutator {
			return hook.UserFunc(func(ctx context.Context, um *generated.UserMutation) (generated.Value, error) {
				v, err := m.Mutate(ctx, um)
				if err != nil {
					return nil, err
				}

				user, ok := v.(*generated.User)
				if !ok {
					return nil, errors.New("unexpected type: expected *generated.User")
				}

				_, err = um.Client().Profile.
					Create().
					SetUser(user).
					Save(ctx)

				if err != nil {
					return nil, err
				}

				return v, nil
			})
		}, ent.OpCreate),
	}
}
