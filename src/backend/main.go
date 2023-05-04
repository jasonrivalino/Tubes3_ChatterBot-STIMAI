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
