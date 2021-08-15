package model

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Message is the model of Message
type Message struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
	ChatID    primitive.ObjectID  `bson:"chat_id,omitempty" json:"chat_id,omitempty"`
	UserID    uuid.UUID           `bson:"user,omitempty" json:"user,omitempty"`
	Body      string              `bson:"body,omitempty" json:"body,omitempty"`
	CreatedAt primitive.Timestamp `bson:"create_at,omitempty" json:"create_at,omitempty"`
}
