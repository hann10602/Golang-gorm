package storage

import (
	"context"
	"gin_mysql/enum"
	"gin_mysql/modules/user/model"
)

func (s *sqlStore) GetUser(ctx context.Context, cond map[string]interface{}) (*model.Users, error) {
	var data model.Users

	if err := s.db.Table(enum.USERS_TABLE).Where(cond).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
