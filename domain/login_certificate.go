package domain

import (
	"github.com/iissy/gooa/types"
	"github.com/juju/errors"
)

func GetLoginCertificate(id string) (*types.LoginCertificate, error) {
	result := new(types.LoginCertificate)
	sql := "select * from login_certificate where app_id = ?"
	err := dingDB.SelectOne(&result, sql, id)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return result, err
}
