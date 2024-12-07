package models

type TypeSale struct {
	IdSale int    `gorm:"type:int(11);primaryKey;autoIncrement" json:"id_sale"`
	Title  string `gorm:"type:varchar(30);not null" json:"title"`
	Count  string `gorm:"type:varchar(30);not null" json:"count"`
}

func (TypeSale) TableName() string {
	return "type_sale"
}
