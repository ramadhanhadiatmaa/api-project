package models

type Seller struct {
	IdType int    `gorm:"type:int(11);primaryKey;autoIncrement" json:"id_type"`
	Type   string `gorm:"type:varchar(60)" json:"type"`
}

func (Seller) TableName() string {
	return "seller"
}
