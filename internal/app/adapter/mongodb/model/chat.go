package model

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Chat is the model of Chat
type Chat struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
	UserIDs   []uuid.UUID         `bson:"users,omitempty" json:"users,omitempty"`
	Label     string              `bson:"label,omitempty" json:"label,omitempty"`
	CreatedAt primitive.Timestamp `bson:"create_at,omitempty" json:"create_at,omitempty"`
}
