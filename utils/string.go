package utils

import (
	"encoding/json"
	"github.com/iissy/gooa/utils/uuid"
)

func UUID() string {
	return uuid.V4Compressed()[0:8]
}

func JsonToString(obj interface{}) string {
	data, err := json.Marshal(obj)
	if WriteErrorLog("json 系列化出错", err) {
		return ""
	}

	return string(data)
}
