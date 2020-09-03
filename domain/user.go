package domain

import (
	"fmt"
	"github.com/iissy/gooa/types"
	"github.com/iissy/gooa/utils"
	"github.com/juju/errors"
	"time"
)

var (
	userColumn = "uuid,userid,unionid,name,tel,workPlace,remark,mobile,email,orgEmail,active,orderInDepts,isAdmin,isBoss,isLeaderInDepts,isHide,position,avatar,hiredDate,jobnumber,extattr,isSenior,stateCode,realAuthed"
)

func SaveUser(o *types.User) (int64, error) {
	id := utils.UUID()
	sql := "INSERT INTO user(%s) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE " +
		"name = ?,tel = ?,workPlace = ?,remark = ?,mobile = ?,email = ?,orgEmail = ?,active = ?,orderInDepts = ?,isAdmin = ?,isBoss = ?,isLeaderInDepts = ?,isHide = ?,position = ?,avatar = ?,hiredDate = ?,jobnumber = ?,extattr = ?,isSenior = ?,stateCode = ?,realAuthed = ?,update_time = ?;"
	sql = fmt.Sprintf(sql, userColumn)
	result, err := dingDB.Exec(sql, id, o.UserId, o.UnionId, o.Name, o.Tel, o.WorkPlace, o.Remark, o.Mobile, o.Email, o.OrgEmail, o.Active, o.OrderInDepts, o.IsAdmin, o.IsBoss, o.IsLeaderInDepts,
		o.IsHide, o.Position, o.Avatar, o.HiredDate, o.JobNumber, o.Extattr, o.IsSenior, o.StateCode, o.RealAuthed, o.Name, o.Tel, o.WorkPlace, o.Remark, o.Mobile, o.Email, o.OrgEmail, o.Active, o.OrderInDepts, o.IsAdmin, o.IsBoss, o.IsLeaderInDepts,
		o.IsHide, o.Position, o.Avatar, o.HiredDate, o.JobNumber, o.Extattr, o.IsSenior, o.StateCode, o.RealAuthed, time.Now())
	if err != nil {
		return 0, errors.Trace(err)
	}

	return result.RowsAffected()
}

func SaveUserDepartment(o *types.UserDepartment) (int64, error) {
	id := utils.UUID()
	sql := "INSERT INTO user_department(uuid,user_id,department_id) values (?,?,?) ON DUPLICATE KEY UPDATE update_time = ?;"
	result, err := dingDB.Exec(sql, id, o.UserID, o.DepartmentId, time.Now())
	if err != nil {
		return 0, errors.Trace(err)
	}

	return result.RowsAffected()
}
