package main

import (
	"lee-gin/bootstrap"
	"lee-gin/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//初始化配置
	bootstrap.InitializeConfig()
	r := gin.Default()
	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")
	//测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//启动服务器

	r.Run(global.App.Config.App.AppLocal + ":" + global.App.Config.App.Port)
}
