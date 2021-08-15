package repository

import (
	"github.com/orensimple/trade-chat-app/internal/app/adapter/mongodb/model"
)

// Chat is interface of chat repository
type Chat interface {
	Create(u *model.Chat) error
	Get(f *model.Chat) (*model.Chat, error)
	Search(f *model.Chat) ([]*model.Chat, error)
}
