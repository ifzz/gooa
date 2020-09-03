package fetch

import (
	"github.com/go-resty/resty/v2"
	"github.com/iissy/gooa/domain"
	"github.com/iissy/gooa/types"
	"github.com/juju/errors"
)

var client *resty.Client

func init() {
	client = resty.New()
}

// fetch 获取钉钉数据统一入口
func fetch(dict map[string]string, t interface{}, url string) (interface{}, error) {
	token, err := getToken()
	if err != nil {
		return nil, errors.Trace(err)
	}

	if token == "" {
		return nil, errors.New("token is null")
	}

	dict["access_token"] = token
	resp, err := client.R().SetQueryParams(dict).
		SetHeader("Accept", "application/json").
		SetResult(t).
		Get(url)

	if err != nil {
		return nil, errors.Trace(err)
	}

	return resp.Result(), err
}

// getToken 获取凭证
func getToken() (string, error) {
	dict := map[string]string{
		"appkey":    domain.AppKey,
		"appsecret": domain.AppSecret,
	}

	resp, err := client.R().SetQueryParams(dict).
		SetHeader("Accept", "application/json").
		SetResult(&types.AccessToken{}).
		Get("https://oapi.dingtalk.com/gettoken")

	if err != nil {
		return "", errors.Trace(err)
	}

	result, ok := resp.Result().(*types.AccessToken)
	if ok && result.ErrCode == 0 {
		return result.AccessToken, nil
	}

	return "", errors.New(result.ErrMsg)
}
