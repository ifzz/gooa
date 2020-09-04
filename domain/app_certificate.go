package domain

import (
	"github.com/iissy/gooa/types"
	"github.com/juju/errors"
)

func GetAppCertificate() (*types.AppCertificate, error) {
	result := new(types.AppCertificate)
	sql := "select * from app_certificate limit 1"
	err := dingDB.SelectOne(&result, sql)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return result, err
}
