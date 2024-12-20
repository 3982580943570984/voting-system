package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("contents").
			NotEmpty().
			MaxLen(2000),
	}
}

// Mixins of the Comment.
func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Comment.Type).
			From("parent").
			Unique(),

		edge.From("user", User.Type).
			Ref("comments").
			Unique(),

		edge.From("election", Election.Type).
			Ref("comments").
			Unique(),
	}
}
