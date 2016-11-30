package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (ctrl HealthController) Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
