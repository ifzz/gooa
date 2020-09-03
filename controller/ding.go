package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iissy/gooa/types"
	"github.com/iissy/gooa/utils"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func Getuserinfo(ctx *gin.Context) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1000000, 10) // 毫秒时间戳
	signature := utils.EncodeSHA256(timestamp, "k2JAcZPLBMQo9etoQe5mcKnUngzop7WXyFDaL_Fpkpcqky6EvBLp8OXPL4K7GMNi")
	code := ctx.Query("code")

	body := fmt.Sprintf(`{"tmp_auth_code": "%s"}`, code)
	url := fmt.Sprintf("https://oapi.dingtalk.com/sns/getuserinfo_bycode?accessKey=dingoalb2yow51ilsojddu&timestamp=%s&signature=%s", timestamp, signature)
	resp, err := client.R().SetHeader("Accept", "application/json").SetBody(body).SetResult(&types.UserInfoDTO{}).Post(url)
	if err != nil {
		logrus.Fatal("fail")
	}

	result := resp.Result().(*types.UserInfoDTO)
	if result.ErrCode == 0 {
		fmt.Println(result)
	}
}
