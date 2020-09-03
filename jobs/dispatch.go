package jobs

import (
	"github.com/sirupsen/logrus"
)

type Exec func()

func (e Exec) Before() {
	logrus.Info("------开始------")
}

func (e Exec) After() {
	logrus.Info("------结束------\n")
}

func Dispatch(e Exec) Exec {
	return func() {
		e.Before()
		defer e.After()
		e()
	}
}
