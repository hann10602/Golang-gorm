package storage

import (
	"context"
	"gin_mysql/enum"
	"gin_mysql/modules/user/model"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *model.CreateUserDTO) error {
	if err := s.db.Table(enum.USERS_TABLE).Create(&data).Error; err != nil {
		return err
	}

	return nil
}
