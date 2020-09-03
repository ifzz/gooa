package utils

import (
	"github.com/juju/errors"
	"github.com/kataras/golog"
	"log"
)

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func WriteErrorLog(kind string, err error) bool {
	if err != nil {
		golog.Errorf("[%s] %s", kind, errors.ErrorStack(err))
		return true
	} else {
		return false
	}
}
