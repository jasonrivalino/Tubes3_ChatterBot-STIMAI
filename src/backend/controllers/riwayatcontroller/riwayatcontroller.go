package riwayatcontroller

import (
	"encoding/json"
	"net/http"

	"backend/models"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var riwayat []models.Riwayat

	models.DB.Find(&riwayat)
	c.JSON(http.StatusOK, gin.H{"riwayat": riwayat})

}

func Show(c *gin.Context) {
	var riwayat []models.Riwayat
	id := c.Param("id")

	// Ambil semua data riwayat yang memiliki id yang sama dengan parameter id_riwayat
	if err := models.DB.Where("id_riwayat = ?", id).Find(&riwayat).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"riwayat": riwayat})
}

func Create(c *gin.Context) {

	// var product models.Product

	// if err := c.ShouldBindJSON(&product); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }

	// models.DB.Create(&product)
	// c.JSON(http.StatusOK, gin.H{"product": product})

	var riwayat models.Riwayat

	if err := c.ShouldBindJSON(&riwayat); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&riwayat)
	c.JSON(http.StatusOK, gin.H{"riwayat": riwayat})
}

func Update(c *gin.Context) {
	// var product models.Product
	// id := c.Param("id")

	// if err := c.ShouldBindJSON(&product); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }

	// if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate product"})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})

	var riwayat models.Riwayat
	id := c.Param("id")

	if err := c.ShouldBindJSON(&riwayat); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&riwayat).Where("id_riwayat = ?", id).Updates(&riwayat).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate riwayat"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})

}

func Delete(c *gin.Context) {

	// var product models.Product

	// var input struct {
	// 	Id json.Number
	// }

	// if err := c.ShouldBindJSON(&input); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }

	// id, _ := input.Id.Int64()
	// if models.DB.Delete(&product, id).RowsAffected == 0 {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus product"})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})

	var riwayat models.Riwayat

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&riwayat, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus riwayat"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
