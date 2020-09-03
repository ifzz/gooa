package fetch

import (
	"github.com/iissy/gooa/domain"
	"github.com/iissy/gooa/types"
	"github.com/iissy/gooa/utils"
	"github.com/juju/errors"
	"strconv"
)

// GetUserList 获取部门用户详情
func GetUserList(depId int64, offset, size int) {
	dict := map[string]string{
		"department_id": strconv.FormatInt(depId, 10),
		"offset":        strconv.Itoa(offset),
		"size":          strconv.Itoa(size),
	}

	body, err := fetch(dict, &types.UserListDTO{}, "https://oapi.dingtalk.com/user/listbypage")
	if utils.WriteErrorLog("获取部门用户详情", err) {
		return
	}
	hasMore := false
	result := body.(*types.UserListDTO)
	if result.ErrCode == 0 {
		hasMore = result.HasMore
		for _, u := range result.UserList {
			_, err := domain.SaveUser(&u)
			utils.WriteErrorLog("同步部门用户详情入库", err)

			for _, d := range u.Department {
				_, err := domain.SaveUserDepartment(&types.UserDepartment{UserID: u.UserId, DepartmentId: d})
				utils.WriteErrorLog("同步用户部门入库", err)
			}
		}
	} else {
		utils.WriteErrorLog("获取部门用户详情", errors.New(result.ErrMsg))
	}

	if hasMore {
		offset = offset + size
		GetUserList(depId, offset, size)
	}
}
