package biz

import (
	"context"
	"fmt"
	"gin_mysql/modules/user/model"
)

type UpdateUserStorage interface {
	GetUser(ctx context.Context, conds map[string]interface{}) (data *model.Users, err error)
	UpdateUser(ctx context.Context, cond map[string]interface{}, data *model.UpdateUserDTO) error
}

type updateUserBiz struct {
	store UpdateUserStorage
}

func NewUpdateUserBiz(store UpdateUserStorage) *updateUserBiz {
	return &updateUserBiz{store: store}
}

func (biz *updateUserBiz) UpdateUserById(ctx context.Context, id int, dataUpdate *model.UpdateUserDTO) error {
	data, err := biz.store.GetUser(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	fmt.Printf("Xu ly day ne %v", data)

	if err := biz.store.UpdateUser(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return err
	}

	return nil
}
