package storage

import (
	"context"
	"gin_mysql/common"
	"gin_mysql/enum"
)

func (s *sqlStore) DeleteUser(ctx context.Context, conds map[string]interface{}) error {
	if err := s.db.Table(enum.USERS_TABLE).Where(conds).Updates(map[string]interface{}{
		"status": "DELETED",
	}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
