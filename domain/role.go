package domain

import (
	"github.com/iissy/gooa/types"
	"github.com/iissy/gooa/utils"
	"github.com/juju/errors"
	"time"
)

func SaveRoleGroup(o *types.RoleGroup) (int64, error) {
	id := utils.UUID()
	sql := "INSERT INTO role_group(uuid,id,name) values (?,?,?) ON DUPLICATE KEY UPDATE name = ?, update_time = ?;"
	result, err := dingDB.Exec(sql, id, o.ID, o.Name, o.Name, time.Now())
	if err != nil {
		return 0, errors.Trace(err)
	}

	return result.RowsAffected()
}

func SaveRole(o *types.Role) (int64, error) {
	id := utils.UUID()
	sql := "INSERT INTO role(uuid,id,name,group_id) values (?,?,?,?) ON DUPLICATE KEY UPDATE name = ?, group_id = ?, update_time = ?;"
	result, err := dingDB.Exec(sql, id, o.ID, o.Name, o.GroupId, o.Name, o.GroupId, time.Now())
	if err != nil {
		return 0, errors.Trace(err)
	}

	return result.RowsAffected()
}

func SaveUserRole(o *types.UserRole) (int64, error) {
	id := utils.UUID()
	sql := "INSERT INTO user_role(uuid,user_id,role_id) values (?,?,?) ON DUPLICATE KEY UPDATE update_time = ?;"
	result, err := dingDB.Exec(sql, id, o.UserId, o.RoleId, time.Now())
	if err != nil {
		return 0, errors.Trace(err)
	}

	return result.RowsAffected()
}
