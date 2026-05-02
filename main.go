package main

import (
	"fast_gin/core"
	"fast_gin/flags"
	"fast_gin/global"
	"fast_gin/routers"
	"fast_gin/service/cron_ser"
)

func main() {
	core.InitLogger()
	flags.Parse()
	global.Config = core.ReadConfig()
	global.DB = core.InitGorm()
	global.Redis = core.InitRedis()
	global.Es = core.EsConnect()
	cron_ser.CronInit()

	flags.Run()
	routers.Run()
}
