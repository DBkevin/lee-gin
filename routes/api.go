package routes

import (
	"lee-gin/app/common/request"
	"lee-gin/app/controllers/app"
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

	router.POST("/user/register", func(ctx *gin.Context) {
		var form request.Register
		if err := ctx.ShouldBindJSON(&form); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"error": request.GetErrorMsg(form, err),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	router.POST("/auth/register", app.Register)
}
