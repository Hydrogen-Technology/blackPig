package schema

import (
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
		field.Int32("id"),
		field.String("username"),
		field.String("phone"),
		field.Strings("carPictures"),
		field.String("openid"),
		field.String("gender"),
		field.String("description"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
