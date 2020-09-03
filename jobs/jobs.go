package jobs

import (
	"github.com/iissy/gooa/fetch"
	"github.com/iissy/gooa/utils"
)

func Run() {
	fetch.GetRoleGroupList(utils.Offset, utils.Size)
	fetch.GetDepartmentList(utils.Offset, utils.Size)
	fetch.GetOpenExtContactList(utils.Offset, utils.Size)
}
