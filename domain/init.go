package domain

import (
	"github.com/go-gorp/gorp/v3"
	"github.com/iissy/gooa/utils"
)

var dingDB, iissyDB *gorp.DbMap
var AppKey = ""
var AppSecret = ""

func init() {
	dingDB = utils.InitDingDb()
	iissyDB = utils.InitIissyDb()

	app, err := GetAppCertificate()
	utils.CheckErr(err, "serialized AppKeySecret fail")

	AppKey = app.AppKey
	AppSecret = app.AppSecret
}
