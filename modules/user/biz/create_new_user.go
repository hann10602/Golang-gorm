package biz

import (
	"context"
	"gin_mysql/modules/user/model"
	"strings"
)

type CreateUserStorage interface {
	CreateUser(ctx context.Context, data *model.CreateUserDTO) error
}

type createUserBiz struct {
	store CreateUserStorage
}

func NewCreateUserBiz(store CreateUserStorage) *createUserBiz {
	return &createUserBiz{store: store}
}

func (biz *createUserBiz) CreateNewUser(ctx context.Context, data *model.CreateUserDTO) error {
	username := strings.TrimSpace(data.Username)
	password := strings.TrimSpace(data.Password)

	if len(username) == 0 {
		return model.ErrUsernameIsBlank
	}

	if len(password) == 0 {
		return model.ErrPasswordIsBlank
	}

	if err := biz.store.CreateUser(ctx, data); err != nil {
		return err
	}

	return nil
}
