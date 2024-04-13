package biz

import (
	"context"
	"gin_mysql/common"
	"gin_mysql/modules/user/model"
)

type GetUsersStorage interface {
	GetUsers(ctx context.Context, filter *model.Filter, paging *common.Paging, moreKeys ...string) (*[]model.Users, error)
}

type getUsersBiz struct {
	store GetUsersStorage
}

func NewGetUsersBiz(store GetUsersStorage) *getUsersBiz {
	return &getUsersBiz{store: store}
}

func (biz *getUsersBiz) GetUsers(ctx context.Context, filter *model.Filter, paging *common.Paging) (*[]model.Users, error) {
	data, err := biz.store.GetUsers(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(model.EntityName, err)
	}

	return data, nil
}
