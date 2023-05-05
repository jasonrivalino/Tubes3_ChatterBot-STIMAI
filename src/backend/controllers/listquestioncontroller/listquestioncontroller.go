package listquestioncontroller

import (
	"encoding/json"
	"net/http"

	"backend/models"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var listquestion []models.ListQuestion

	models.DB.Find(&listquestion)
	c.JSON(http.StatusOK, gin.H{"listquestion": listquestion})

}

func Show(c *gin.Context) {
	var listquestion []models.ListQuestion
	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).Find(&listquestion).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"listquestion": listquestion})
}

func GetAllListQuestion() ([]models.ListQuestion, error) {
	var listquestion []models.ListQuestion

	if err := models.DB.Find(&listquestion).Error; err != nil {
		return listquestion, err
	}

	return listquestion, nil
}

func SearchQuestionOnDatabase(question string) (models.ListQuestion, error) {
	var listquestion models.ListQuestion

	if err := models.DB.Where("pertanyaan = ?", question).Find(&listquestion).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return listquestion, err
		default:
			return listquestion, err
		}
	}

	return listquestion, nil

}

func Create(c *gin.Context) {

	var listquestion models.ListQuestion

	if err := c.ShouldBindJSON(&listquestion); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&listquestion)
	c.JSON(http.StatusOK, gin.H{"listquestion": listquestion})
}

func AddNewQuestion(question string, answer string) error {

	var listquestion models.ListQuestion

	listquestion.Pertanyaan = question
	listquestion.Jawaban = answer

	return models.DB.Create(&listquestion).Error

}

func Update(c *gin.Context) {
	var listquestion models.ListQuestion
	id := c.Param("id")

	if err := c.ShouldBindJSON(&listquestion); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&listquestion).Where("id = ?", id).Updates(&listquestion).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate listquestion"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})

}

func UpdateAnswer(data models.ListQuestion, ans string) error {

	data.Jawaban = ans

	// Update data ke database
	return models.DB.Model(&data).Where("id = ?", data.Id).Updates(&data).Error
}

func Delete(c *gin.Context) {
	var listquestion models.ListQuestion

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&listquestion, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus listquestion"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}

func DeleteQuestion(question string) error {

	var listquestion models.ListQuestion

	err := models.DB.Where("pertanyaan = ?", question).Delete(&listquestion).Error

	switch err {
	case gorm.ErrRecordNotFound:
		return err
	default:
		return err
	}

}
