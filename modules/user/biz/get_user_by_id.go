package biz

import (
	"context"
	"gin_mysql/modules/user/model"
)

type GetUserStorage interface {
	GetUser(ctx context.Context, conds map[string]interface{}) (data *model.Users, err error)
}

type getUserBiz struct {
	store GetUserStorage
}

func NewGetUserBiz(store GetUserStorage) *getUserBiz {
	return &getUserBiz{store: store}
}

func (biz *getUserBiz) GetUserById(ctx context.Context, id int) (*model.Users, error) {
	data, err := biz.store.GetUser(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	return data, nil
}
