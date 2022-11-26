package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ylinyang/gin-web/logger"
	"net/http"
)

func SetUp() (r *gin.Engine) {
	r = gin.New()

	//	将zap日志中间间注册到gin中
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//	 注册路由信息
	r.GET("ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return
}
