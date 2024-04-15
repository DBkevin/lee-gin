package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "PONG")
	})

	router.GET("/test", func(ctx *gin.Context) {
		time.Sleep(5 * time.Second)
		ctx.String(http.StatusOK, "success")
	})
}
