package fetch

import (
	"github.com/iissy/gooa/domain"
	"github.com/iissy/gooa/types"
	"github.com/iissy/gooa/utils"
	"github.com/juju/errors"
	"strconv"
)

// GetOpenExtContactList 获取外部联系人列表
func GetOpenExtContactList(offset, size int) {
	dict := map[string]string{
		"offset": strconv.Itoa(offset),
		"size":   strconv.Itoa(size),
	}

	body, err := fetch(dict, &types.OpenExtContactDTO{}, "https://oapi.dingtalk.com/topapi/extcontact/list")
	if utils.WriteErrorLog("获取外部联系人列表", err) {
		return
	}
	hasMore := false
	result := body.(*types.OpenExtContactDTO)
	if result.ErrCode == 0 {
		if len(result.Results) >= size {
			hasMore = true
		}
		for _, u := range result.Results {
			_, err := domain.SaveOpenExtContact(&u)
			utils.WriteErrorLog("同步外部联系人列表入库", err)
		}
	} else {
		utils.WriteErrorLog("获取外部联系人列表", errors.New(result.ErrMsg))
	}

	if hasMore {
		offset = offset + size
		GetOpenExtContactList(offset, size)
	}
}
