package types

type AppCertificate struct {
	UUID       string `db:"uuid"`
	AppKey     string `db:"app_key"`
	AppSecret  string `db:"app_secret"`
	CreateTime string `db:"create_time"`
	Status     bool   `db:"status"`
}
