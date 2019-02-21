package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mirzakhany/pkg/version"
	api "gopkg.in/fukata/golang-stats-api-handler.v1"
)

func abortWithError(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

// RootHandler handle root URL
func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"text": "Welcome to gorender server.",
	})
}

// HeartbeatHandler handle heartbeat requests
func HeartbeatHandler(c *gin.Context) {
	c.AbortWithStatus(http.StatusOK)
}

// VersionHandler return version
func VersionHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": version.GetVersion(),
	})
}

// APIStatusHandler return api status
func APIStatusHandler(c *gin.Context) {
	c.JSON(http.StatusOK, api.GetStats())
}
