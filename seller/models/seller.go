package models

type Seller struct {
	Id int    `gorm:"type:int(11);primaryKey;autoIncrement" json:"kd_kel"`
	NmKel string `gorm:"type:varchar(60)" json:"nm_kel"`
}

func (Seller) TableName() string {
	return "seller"
}
