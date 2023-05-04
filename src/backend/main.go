package main

import (
	"backend/controllers/riwayatcontroller"
	"backend/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/api/histories", riwayatcontroller.Index)
	r.GET("/api/history/:id", riwayatcontroller.Show)
	r.POST("/api/history", riwayatcontroller.Create)
	r.PUT("/api/history/:id", riwayatcontroller.Update)
	r.DELETE("/api/history", riwayatcontroller.Delete)

	r.Run()
}
