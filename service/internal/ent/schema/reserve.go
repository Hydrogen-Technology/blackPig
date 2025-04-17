package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Reserve struct {
	ent.Schema
}

func (Reserve) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Time("created_at"),
		field.Time("updated_at"),
		field.Int32("user_id"), // 用户id
		field.String("start_time"),
		field.String("ordered_id"), // 被预约的导员的id
		field.String("day"),
		field.String("detail"),
	}
}
func (Reserve) Edges() []ent.Edge {
	return nil
}
