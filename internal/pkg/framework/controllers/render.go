package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/mirzakhany/pkg/logger"
	"github.com/mirzakhany/gorender/internal/global"
	"net/http"
)

// RenderHandler start render progress
func IssueToken(c *gin.Context) {

	var form global.RenderReq
	if err := c.ShouldBindWith(&form, binding.JSON); err != nil {
		logger.Info(err)
		abortWithError(c, http.StatusBadRequest, "invalid render data")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result":   "ok",
	})
}
