package fetch

import (
	"github.com/iissy/gooa/domain"
	"github.com/iissy/gooa/types"
	"github.com/iissy/gooa/utils"
	"github.com/juju/errors"
	"strconv"
)

// GetDepartmentList 获取部门列表
func GetDepartmentList(offset, size int) {
	dict := map[string]string{}

	body, err := fetch(dict, &types.DepartmentListDTO{}, "https://oapi.dingtalk.com/department/list")
	if utils.WriteErrorLog("获取部门列表", err) {
		return
	}

	result := body.(*types.DepartmentListDTO)
	if result.ErrCode == 0 {
		for _, u := range result.Department {
			GetDepartmentDetail(u.ID)
			GetUserList(u.ID, offset, size)
		}
	} else {
		utils.WriteErrorLog("获取部门列表", errors.New(result.ErrMsg))
	}
}

// GetDepartmentDetail 获取部门详情
func GetDepartmentDetail(id int64) {
	dict := map[string]string{
		"id": strconv.FormatInt(id, 10),
	}

	body, err := fetch(dict, &types.DepartmentDTO{}, "https://oapi.dingtalk.com/department/get")
	if utils.WriteErrorLog("获取部门详情", err) {
		return
	}

	result := body.(*types.DepartmentDTO)
	if result.ErrCode == 0 {
		_, err := domain.SaveDepartment(&result.Department)
		utils.WriteErrorLog("同步部门详情入库", err)
	} else {
		utils.WriteErrorLog("获取部门详情", errors.New(result.ErrMsg))
	}
}
