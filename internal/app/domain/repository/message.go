package repository

import (
	"github.com/orensimple/trade-chat-app/internal/app/adapter/mongodb/model"
)

// Message is interface of message repository
type Message interface {
	Create(u *model.Message) error
	Get(f *model.Message) (*model.Message, error)
	Search(f *model.Message) ([]*model.Message, error)
}
