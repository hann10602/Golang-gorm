package storage

import (
	"context"
	"gin_mysql/common"
	"gin_mysql/enum"
	"gin_mysql/modules/user/model"

	"gorm.io/gorm"
)

func (s *sqlStore) GetUser(ctx context.Context, cond map[string]interface{}) (*model.Users, error) {
	var data model.Users

	if err := s.db.Table(enum.USERS_TABLE).Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, err
	}

	return &data, nil
}
