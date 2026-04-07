package routers

import (
	"fast_gin/global"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Run() {
	//精简日志
	gin.SetMode(global.Config.System.Mode)
	r := gin.Default()
	//静态映射
	r.Static("/uploads", "uploads")

	//路由组
	g := r.Group("api")

	UserRouter(g)
	ImageRouter(g)
	CaptchaRouter(g)

	addr := global.Config.System.Addr()
	if global.Config.System.Mode == "release" {
		logrus.Infof("服务运行在%s", addr)
	}

	r.Run(addr)
}
