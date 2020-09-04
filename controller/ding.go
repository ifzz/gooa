package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iissy/gooa/domain"
	"github.com/iissy/gooa/types"
	"github.com/iissy/gooa/utils"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

func GetUserInfo(ctx *gin.Context) {
	id := ctx.Param("id")
	cert, err := domain.GetLoginCertificate(id)
	if err != nil {
		logrus.Fatal(err)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/1000000, 10) // 毫秒时间戳
	signature := utils.EncodeSHA256(timestamp, cert.AppSecret)
	code := ctx.Query("code")

	body := fmt.Sprintf(`{"tmp_auth_code": "%s"}`, code)
	url := fmt.Sprintf("https://oapi.dingtalk.com/sns/getuserinfo_bycode?accessKey=%s&timestamp=%s&signature=%s", cert.AppID, timestamp, signature)
	resp, err := client.R().SetHeader("Accept", "application/json").SetBody(body).SetResult(&types.UserInfoDTO{}).Post(url)
	if err != nil {
		logrus.Fatal("fail")
	}

	result := resp.Result().(*types.UserInfoDTO)
	fmt.Println(result)
	// TODO 验证返回，实现cookie保持会话状态
	ctx.Redirect(http.StatusFound, "http://localhost:8080/main")
}
