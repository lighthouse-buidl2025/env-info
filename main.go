package main

import (
	"env-info/routes"
	"env-info/services"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	err := services.LoadCategoryMap()
	if err != nil {
		log.Fatalf("failed to load category mapping: %v", err)
	}

	r := gin.Default()

	routes.RegisterEnvInfoRoutes(r)

	err = r.Run(":8023")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		return
	}
}
