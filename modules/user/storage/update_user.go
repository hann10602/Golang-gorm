package storage

import (
	"context"
	"gin_mysql/common"
	"gin_mysql/enum"
	"gin_mysql/modules/user/model"
)

func (s *sqlStore) UpdateUser(ctx context.Context, cond map[string]interface{}, dataUpdate *model.UpdateUserDTO) error {
	if err := s.db.Table(enum.USERS_TABLE).Where(cond).Updates(&dataUpdate).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
