package main

import (
	"backend/controllers/listquestioncontroller"
	"backend/controllers/riwayatcontroller"
	"backend/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	// Add CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.GET("/api/histories", riwayatcontroller.Index)
	r.GET("/api/history/:id", riwayatcontroller.Show)
	r.POST("/api/history", riwayatcontroller.Create)
	r.PUT("/api/history/:id", riwayatcontroller.Update)
	r.DELETE("/api/history", riwayatcontroller.Delete)

	r.GET("/api/list-question", listquestioncontroller.Index)
	r.GET("/api/question/:id", listquestioncontroller.Show)
	r.POST("/api/question", listquestioncontroller.Create)
	r.PUT("/api/question/:id", listquestioncontroller.Update)
	r.DELETE("/api/question", listquestioncontroller.Delete)

	r.Run()
}
