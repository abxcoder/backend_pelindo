package main

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine, controller *Controller) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/negaras", controller.GetNegaras)

		v1.GET("/pelabuhans", controller.GetPelabuhans)

		v1.GET("/barangs", controller.GetBarangs)
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Server is running",
			"version": "1.0.0",
		})
	})
	
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Pelabuhan Nusantara API Server",
			"version": "1.0.0",
			"endpoints": gin.H{
				"countries": "/api/v1/negaras",
				"ports":     "/api/v1/pelabuhans?id_negara={id}",
				"goods":     "/api/v1/barangs?id_pelabuhan={id}",
				"health":    "/health",
			},
		})
	})
}
