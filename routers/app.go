package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/ylinyang/gin-web/logger"
	"net/http"
)

func SetUp() (r *gin.Engine) {

	// gin默认的日志为debug模式 如果设置为release模式 将在不在打印日志  但是zap的日志还是会输出到文件中
	if viper.GetString("app.mode") == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r = gin.New()

	//	将zap日志中间间注册到gin中
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//	 注册路由信息
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("%s", viper.GetString("app.name")))
	})

	return
}
