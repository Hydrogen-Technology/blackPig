package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type TimeList struct {
	ent.Schema
}

func (TimeList) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("created_at"),
		field.String("updated_at"),
		field.Int64("times_id"),
	}
}
func (TimeList) Edges() []ent.Edge {
	return nil
}
