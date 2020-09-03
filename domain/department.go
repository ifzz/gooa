package domain

import (
	"fmt"
	"github.com/iissy/gooa/types"
	"github.com/iissy/gooa/utils"
	"github.com/juju/errors"
	"time"
)

var (
	departmentColumn = "uuid,id,name,parentid,order_no,createDeptGroup,autoAddUser,deptHiding,deptPermits,userPermits,outerDept,outerPermitDepts,outerPermitUsers,orgDeptOwner,deptManagerUseridList,sourceIdentifier,groupContainSubDept,ext"
)

func SaveDepartment(o *types.Department) (int64, error) {
	id := utils.UUID()
	sql := "INSERT INTO department(%s) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE " +
		"name = ?,parentid = ?,order_no = ?,createDeptGroup = ?,autoAddUser = ?,deptHiding = ?,deptPermits = ?,userPermits = ?,outerDept = ?,outerPermitDepts = ?,outerPermitUsers = ?,orgDeptOwner = ?,deptManagerUseridList = ?,sourceIdentifier = ?,groupContainSubDept = ?,ext = ?,update_time = ?;"
	sql = fmt.Sprintf(sql, departmentColumn)
	result, err := dingDB.Exec(sql, id, o.ID, o.Name, o.ParentID, o.Order, o.CreateDeptGroup, o.AutoAddUser, o.DeptHiding, o.DeptPermits, o.UserPermits, o.OuterDept, o.OuterPermitDepts,
		o.OuterPermitUsers, o.OrgDeptOwner, o.DeptManagerUseridList, o.SourceIdentifier, o.GroupContainSubDept, o.Ext, o.Name, o.ParentID, o.Order, o.CreateDeptGroup, o.AutoAddUser,
		o.DeptHiding, o.DeptPermits, o.UserPermits, o.OuterDept, o.OuterPermitDepts, o.OuterPermitUsers, o.OrgDeptOwner, o.DeptManagerUseridList, o.SourceIdentifier, o.GroupContainSubDept, o.Ext, time.Now())
	if err != nil {
		return 0, errors.Trace(err)
	}

	return result.RowsAffected()
}
