package storage

import (
	"context"
	"gin_mysql/common"
	"gin_mysql/enum"
	"gin_mysql/modules/user/model"
)

func (s *sqlStore) GetUsers(ctx context.Context, filter *model.Filter, paging *common.Paging, moreKeys ...string) (*[]model.Users, error) {
	var data []model.Users

	db := s.db.Where("status <> ?", "DELETED")

	if f := filter; f != nil {
		if v := f.Status; v != "" {
			db = db.Where("status = ?", v)
		}
	}

	if err := db.Table(enum.USERS_TABLE).Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := db.Table(enum.USERS_TABLE).Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
