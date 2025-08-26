package schema

import (
	"entgo.io/ent/dialect"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/zaytcevcom/golang-ninja/internal/types"
)

// Problem holds the schema definition for the Problem entity.
type Problem struct {
	ent.Schema
}

// Fields of the Problem.
func (Problem) Fields() []ent.Field {
	return []ent.Field{
		field.Other("id", types.ProblemID{}).
			SchemaType(map[string]string{
				dialect.Postgres: "uuid",
				dialect.SQLite:   "blob",
			}).
			Default(func() types.ProblemID { return types.NewProblemID() }).
			Immutable(),
		field.Other("chat_id", types.ChatID{}).
			SchemaType(map[string]string{
				dialect.Postgres: "uuid",
				dialect.SQLite:   "blob",
			}),
		field.Other("manager_id", types.UserID{}).
			SchemaType(map[string]string{
				dialect.Postgres: "uuid",
				dialect.SQLite:   "blob",
			}).
			Nillable().
			Optional(),
		field.Time("resolved_at").
			Nillable().
			Optional(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the Problem.
func (Problem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("chat", Chat.Type).
			Ref("problems").
			Unique().
			Required().
			Field("chat_id"),
		edge.To("messages", Message.Type),
	}
}
