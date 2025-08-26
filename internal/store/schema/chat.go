package schema

import (
	"github.com/zaytcevcom/golang-ninja/internal/types"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Chat holds the schema definition for the Chat entity.
type Chat struct {
	ent.Schema
}

// Fields of the Chat.
func (Chat) Fields() []ent.Field {
	return []ent.Field{
		field.Other("id", types.ChatID{}).
			SchemaType(map[string]string{
				dialect.Postgres: "uuid",
				dialect.SQLite:   "blob",
			}).
			Default(func() types.ChatID { return types.NewChatID() }).
			Immutable(),
		field.Other("client_id", types.UserID{}).
			SchemaType(map[string]string{
				dialect.Postgres: "uuid",
				dialect.SQLite:   "blob",
			}).
			Unique(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the Chat.
func (Chat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("messages", Message.Type),
		edge.To("problems", Problem.Type),
	}
}
