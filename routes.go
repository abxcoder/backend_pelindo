package main

import "github.com/gin-gonic/gin"

// SetupRoutes - Setup all API routes
func SetupRoutes(router *gin.Engine, controller *Controller) {
	// API v1 group
	v1 := router.Group("/api/v1")
	{
		// Countries endpoints
		v1.GET("/negaras", controller.GetNegaras)

		// Ports endpoints - dengan query parameter id_negara
		v1.GET("/pelabuhans", controller.GetPelabuhans)

		// Goods endpoints - dengan query parameter id_pelabuhan
		v1.GET("/barangs", controller.GetBarangs)
	}

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Server is running",
			"version": "1.0.0",
		})
	})

	// Root endpoint
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
