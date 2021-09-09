package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(),
		field.String("name").
			Default("unknown"),
		field.String("description"),
	}
}
//ID          uuid.UUID      `gorm:primaryKey" json:id`
//Name        string         `json:name`
//Description string         `json:description`
//UpdatedAt   time.Time      `json:updated_at`
//CreatedAt   time.Time      `json:created_at`



// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return nil
}
