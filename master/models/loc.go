package models

type Loc struct {
	ID    int    `gorm:"type:int(11);primaryKey;" json:"id"`
	Name  string `gorm:"type:varchar(60);not null" json:"name"`
	Alias string `gorm:"type:varchar(60);not null" json:"alias"`
	Image string `gorm:"type:varchar(100);not null" json:"image"`
}

func (Loc) TableName() string {
	return "loc"
}
