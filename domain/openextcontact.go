package domain

import (
	"fmt"
	"github.com/iissy/gooa/types"
	"github.com/iissy/gooa/utils"
	"github.com/juju/errors"
	"time"
)

var (
	openExtContactColumn = "uuid,title,share_dept_ids,label_ids,remark,address,name,follower_user_id,state_code,company_name,share_user_ids,mobile,userid,email"
)

func SaveOpenExtContact(o *types.OpenExtContact) (int64, error) {
	id := utils.UUID()
	sql := "INSERT INTO open_ext_contact(%s) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE " +
		"title = ?,share_dept_ids = ?,label_ids = ?,remark = ?,address = ?,name = ?,follower_user_id = ?,state_code = ?,company_name = ?,share_user_ids = ?,mobile = ?,email = ?,update_time=?;"
	sql = fmt.Sprintf(sql, openExtContactColumn)
	result, err := dingDB.Exec(sql, id, o.Title, utils.JsonToString(o.ShareDeptIds), utils.JsonToString(o.LabelIds), o.Remark, o.Address, o.Name, o.FollowerUserId, o.StateCode, o.CompanyName, utils.JsonToString(o.ShareUserIds), o.Mobile, o.UserId, o.Email,
		o.Title, utils.JsonToString(o.ShareDeptIds), utils.JsonToString(o.LabelIds), o.Remark, o.Address, o.Name, o.FollowerUserId, o.StateCode, o.CompanyName, utils.JsonToString(o.ShareUserIds), o.Mobile, o.Email, time.Now())
	if err != nil {
		return 0, errors.Trace(err)
	}

	return result.RowsAffected()
}
