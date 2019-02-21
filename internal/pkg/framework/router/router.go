package router

import (
	"fmt"
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/mirzakhany/pkg/logger"
	"github.com/mirzakhany/pkg/version"
	"github.com/mirzakhany/gorender/internal/global"
	"github.com/mirzakhany/gorender/internal/pkg/framework/controllers"
)

const (
	apiVersionOne = "/api/v1"
)

// initRoutes application routes
func initRoutes(r *gin.Engine) {
	v1 := r.Group(apiVersionOne)
	{
		v1.POST("/render", controllers.IssueToken)
	}
	r.GET("/api-status", controllers.APIStatusHandler)
	r.GET("/healthz", controllers.HeartbeatHandler)
	r.GET("/version", controllers.VersionHandler)
	r.GET("/", controllers.RootHandler)
}

// InitRouter init router
func InitRouter() error {
	if !global.AppConf.Core.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	gin.DefaultWriter = logger.LogAccess.Out
	gin.DefaultErrorWriter = logger.LogError.Out

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(version.HeaderVersionMiddleware())
	r.Use(logger.LogMiddleware())

	initRoutes(r)
	addr := fmt.Sprintf("%s:%d", global.AppConf.Core.Address, global.AppConf.Core.Port)
	logger.Infof("Start server on: %s", addr)

	err := gracehttp.Serve(&http.Server{
		Addr:    addr,
		Handler: r,
	})
	return err
}
