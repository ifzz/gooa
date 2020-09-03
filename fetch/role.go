package fetch

import (
	"github.com/iissy/gooa/domain"
	"github.com/iissy/gooa/types"
	"github.com/iissy/gooa/utils"
	"github.com/juju/errors"
	"strconv"
)

// GetRoleGroupList 获取角色列表
func GetRoleGroupList(offset, size int) {
	dict := map[string]string{
		"offset": strconv.Itoa(offset),
		"size":   strconv.Itoa(size),
	}

	body, err := fetch(dict, &types.RoleGroupResultDTO{}, "https://oapi.dingtalk.com/topapi/role/list")
	if utils.WriteErrorLog("获取角色列表", err) {
		return
	}
	hasMore := false
	result := body.(*types.RoleGroupResultDTO)
	if result.ErrCode == 0 {
		hasMore = result.Result.HasMore
		for _, u := range result.Result.List {
			_, err := domain.SaveRoleGroup(&u)
			utils.WriteErrorLog("同步角色组入库", err)

			for _, r := range u.Roles {
				r.GroupId = u.ID
				_, err := domain.SaveRole(&r)
				utils.WriteErrorLog("同步角色入库", err)

				GetUserListByRole(&r, utils.Offset, utils.Size)
			}
		}
	} else {
		utils.WriteErrorLog("获取角色列表", errors.New(result.ErrMsg))
	}

	if hasMore {
		offset = offset + size
		GetRoleGroupList(offset, size)
	}
}

// GetUserListByRole 获取角色下的员工列表
func GetUserListByRole(r *types.Role, offset, size int) {
	dict := map[string]string{
		"role_id": strconv.FormatInt(r.ID, 10),
		"offset":  strconv.Itoa(offset),
		"size":    strconv.Itoa(size),
	}

	body, err := fetch(dict, &types.RoleUserListResultDTO{}, "https://oapi.dingtalk.com/topapi/role/simplelist")
	if utils.WriteErrorLog("获取角色下的员工列表", err) {
		return
	}
	hasMore := false
	result := body.(*types.RoleUserListResultDTO)
	if result.ErrCode == 0 {
		hasMore = result.Result.HasMore
		for _, u := range result.Result.List {
			_, err := domain.SaveUserRole(&types.UserRole{UserId: u.UserId, RoleId: r.ID, UserName: u.Name, RoleName: r.Name})
			utils.WriteErrorLog("同步用户角色入库", err)
		}
	} else {
		utils.WriteErrorLog("获取角色下的员工列表", errors.New(result.ErrMsg))
	}

	if hasMore {
		offset = offset + size
		GetUserListByRole(r, offset, size)
	}
}
