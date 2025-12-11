package controllers

import (
	"net/http"
	"order_service/internal/utils"

	"github.com/gin-gonic/gin"
)

func GetTestToken(c *gin.Context) {
	token, _ := utils.GenerateToken("test-user")
	c.JSON(http.StatusOK, gin.H{"token": token})
}
