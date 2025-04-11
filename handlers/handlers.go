package handlers

import (
	"env-info/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func EnvInfoPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"module":  "env-info",
		"message": "pong",
	})
}

func GetContractInfo(c *gin.Context) {
	address := c.Param("address")
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "address is required",
		})
		return
	}

	info, err := services.GetContractInfo(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, info)
}
