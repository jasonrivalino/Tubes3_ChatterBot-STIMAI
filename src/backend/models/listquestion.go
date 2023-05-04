package models

type ListQuestion struct {
	Id         int64  `gorm:"primaryKey" json:"id"`
	Pertanyaan string `gorm:"type:varchar(500)" json:"pertanyaan"`
	Jawaban    string `gorm:"type:varchar(500)" json:"jawaban"`
}

func (ListQuestion) TableName() string {
	return "listquestion"
}
