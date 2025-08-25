package schema

import (
	"entgo.io/ent/dialect"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/zaytcevcom/golang-ninja/internal/types"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.Other("id", types.MessageID{}).
			SchemaType(map[string]string{
				dialect.Postgres: "uuid",
				dialect.SQLite:   "blob",
			}).
			Default(func() types.MessageID { return types.NewMessageID() }).
			Immutable(),
		field.Other("chat_id", types.ChatID{}).
			SchemaType(map[string]string{
				dialect.Postgres: "uuid",
				dialect.SQLite:   "blob",
			}),
		field.Other("problem_id", types.ProblemID{}).
			SchemaType(map[string]string{
				dialect.Postgres: "uuid",
				dialect.SQLite:   "blob",
			}),
		field.Other("author_id", types.UserID{}).
			SchemaType(map[string]string{
				dialect.Postgres: "uuid",
				dialect.SQLite:   "blob",
			}),
		field.Bool("is_visible_for_client"),
		field.Bool("is_visible_for_manager"),
		field.Text("body"),
		field.Time("checked_at").
			Nillable().
			Optional(),
		field.Bool("is_blocked"),
		field.Bool("is_service"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("chat", Chat.Type).
			Ref("messages").
			Unique().
			Required().
			Field("chat_id"),
		edge.From("problem", Problem.Type).
			Ref("messages").
			Unique().
			Required().
			Field("problem_id"),
	}
}
