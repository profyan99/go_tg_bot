package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("create_at").
			Immutable().
			Default(time.Now),
		field.Time("update_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}
