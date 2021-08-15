package usecase

import (
	"github.com/orensimple/trade-chat-app/internal/app/adapter/mongodb/model"
	"github.com/orensimple/trade-chat-app/internal/app/domain/repository"
)

// CreateMessage create new message
func CreateMessage(r repository.Message, u *model.Message) (*model.Message, error) {
	err := r.Create(u)

	return u, err
}

// GetMessage find message by filter
func GetMessage(r repository.Message, f *model.Message) (*model.Message, error) {
	res, err := r.Get(f)

	return res, err
}

// SearchMessage find message by filter
func SearchMessage(r repository.Message, f *model.Message) ([]*model.Message, error) {
	res, err := r.Search(f)

	return res, err
}
