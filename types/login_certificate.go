package types

import "time"

type LoginCertificate struct {
	UUID       string     `db:"uuid"`
	AppID      string     `db:"app_id"`
	AppSecret  string     `db:"app_secret"`
	Name       string     `db:"name"`
	CreateTime *time.Time `db:"create_time"`
	UpdateTime *time.Time `db:"update_time"`
}
