package cron_ser

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func CronInit() {
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	crontab := cron.New(cron.WithSeconds(), cron.WithLocation(timezone))
	crontab.AddFunc("0 0 2 3 * *", timeFunc)

	crontab.Start()
}

func timeFunc() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
