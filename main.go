package main

import (
	"github.com/iissy/gooa/api"
	"github.com/iissy/gooa/jobs"
	"github.com/robfig/cron/v3"
)

func main() {
	go jobs.Run()

	contact := cron.New(cron.WithSeconds(),
		cron.WithChain(cron.Recover(cron.DefaultLogger)),
		cron.WithChain(cron.DelayIfStillRunning(cron.DefaultLogger)))

	contact.AddFunc("0 0 0/1 * * ?", jobs.Dispatch(jobs.Run))
	contact.Start()

	api.Start()
}
