package usecase

import (
	"github.com/orensimple/trade-chat-app/internal/app/adapter/mongodb/model"
	"github.com/orensimple/trade-chat-app/internal/app/domain/repository"
)

// CreateChat create new chat
func CreateChat(r repository.Chat, u *model.Chat) (*model.Chat, error) {
	err := r.Create(u)

	return u, err
}

// GetChat find chat by filter
func GetChat(r repository.Chat, f *model.Chat) (*model.Chat, error) {
	res, err := r.Get(f)

	return res, err
}

// SearchChat find chats by filter
func SearchChat(r repository.Chat, f *model.Chat) ([]*model.Chat, error) {
	res, err := r.Search(f)

	return res, err
}
