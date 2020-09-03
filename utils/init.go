package utils

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
	"github.com/sirupsen/logrus"
	"log"
)

const (
	defaultConfigPath = "conf/config.json"
)

func init() {
	loadConfig()
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     false,
		TimestampFormat: "2006/01/02 15:04:05",
		FullTimestamp:   true,
	})
}

func loadConfig() {
	if err := config.Load(file.NewSource(
		file.WithPath(defaultConfigPath),
	)); err != nil {
		log.Panic(err)
	}
}
