package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Times struct {
	ent.Schema
}

func (Times) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("userId"),
		field.String("timeSpace"),
		field.String("startTime"),
	}
}

func (Times) Edges() []ent.Edge {
	return nil
}
