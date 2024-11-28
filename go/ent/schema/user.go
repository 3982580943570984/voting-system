package schema

import (
	"context"
	"errors"
	"net/mail"
	"time"

	"voting-system/ent/generated"
	"voting-system/ent/generated/hook"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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

		field.Time("created_at").
			Immutable().
			Default(time.Now),

		field.Time("last_login").
			Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", Profile.Type).
			Unique(),

		edge.To("role", Role.Type).
			Unique(),

		edge.To("elections", Election.Type),

		edge.To("comments", Comment.Type),

		edge.To("votes", Vote.Type),
	}
}

// Hooks of the User.
func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				v, err := next.Mutate(ctx, m)
				if err != nil {
					return nil, err
				}

				user, ok := v.(*generated.User)
				if !ok {
					return nil, errors.New("unexpected type: expected *ent.User")
				}

				userMutation, ok := m.(*generated.UserMutation)
				if !ok {
					return nil, errors.New("unexpected mutation type: expected *ent.UserMutation")
				}

				_, err = userMutation.Client().Profile.
					Create().
					SetUser(user).
					Save(ctx)

				if err != nil {
					return nil, err
				}

				return v, nil
			})
		}, ent.OpCreate),

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

				_, err = um.Client().Role.
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
