package domain

import (
	"github.com/iissy/gooa/types"
	"github.com/juju/errors"
)

func GetAccountAppKey() (*types.AccountConfig, error) {
	result := new(types.AccountConfig)
	sql := "select * from account_config limit 1"
	err := dingDB.SelectOne(&result, sql)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return result, err
}
