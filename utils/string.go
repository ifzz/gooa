package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"github.com/iissy/gooa/utils/uuid"
	"net/url"
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

func EncodeSHA256(message, secret string) string {
	// 钉钉签名算法实现
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	sum := h.Sum(nil) // 二进制流
	message1 := base64.StdEncoding.EncodeToString(sum)

	uv := url.Values{}
	uv.Add("0", message1)
	message2 := uv.Encode()[2:]
	return message2
}
