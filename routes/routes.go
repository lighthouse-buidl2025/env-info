package routes

import (
	"env-info/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterEnvInfoRoutes(r *gin.Engine) {
	env := r.Group("/api/env-info")
	{
		env.GET("/ping", handlers.EnvInfoPing)
		env.GET("/contract/:address", handlers.GetContractInfo)
	}
}
