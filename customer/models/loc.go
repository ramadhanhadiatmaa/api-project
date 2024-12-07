package models

type Loc struct {
	IdLoc string `gorm:"type:varchar(25);primaryKey;" json:"id_loc"`
	Name  string `gorm:"type:varchar(60);not null" json:"name"`
	Image string `gorm:"type:varchar(100);not null" json:"image"`
}

func (Loc) TableName() string {
	return "loc"
}
