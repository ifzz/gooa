package jobs

import (
	"github.com/kataras/golog"
)

type Exec func()

func (e Exec) Before() {
	golog.Infof("------开始------")
}

func (e Exec) After() {
	golog.Infof("------结束------\n")
}

func Dispatch(e Exec) Exec {
	return func() {
		e.Before()
		defer e.After()
		e()
	}
}
