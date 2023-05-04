package models

type Riwayat struct {
	IdRiwayat    int64  `json:"id_riwayat"`
	IdPertanyaan int64  `gorm:"primaryKey;autoIncrement;unsigned" json:"id_pertanyaan"`
	Pertanyaan   string `gorm:"type:varchar(500)" json:"pertanyaan"`
	Jawaban      string `gorm:"type:varchar(500)" json:"jawaban"`
}

func (Riwayat) TableName() string {
	return "riwayat"
}
