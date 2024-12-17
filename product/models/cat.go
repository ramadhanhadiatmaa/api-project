package models

type Cat struct {
	ID    int    `gorm:"type:int(11);primaryKey;autoIncrement" json:"id"`
	Title string `gorm:"type:varchar(30);not null" json:"title"`
}

func (Cat) TableName() string {
	return "cat"
}
