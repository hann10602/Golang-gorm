package biz

import (
	"context"
	"gin_mysql/common"
	"gin_mysql/modules/user/model"
)

type DeleteUserStorage interface {
	DeleteUser(ctx context.Context, conds map[string]interface{}) error
}

type deleteUserBiz struct {
	store DeleteUserStorage
}

func NewDeleteUserBiz(store DeleteUserStorage) *deleteUserBiz {
	return &deleteUserBiz{store: store}
}

func (biz *deleteUserBiz) DeleteUser(ctx context.Context, id int) error {
	if err := biz.store.DeleteUser(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotDeleteEntity(model.EntityName, err)
	}

	return nil
}
