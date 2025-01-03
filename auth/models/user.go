package models

import "time"

type User struct {
	Username  string    `gorm:"type:varchar(100);primaryKey;" json:"username"`
	Password  string    `gorm:"type:varchar(255);" json:"password"`
	Email     string    `gorm:"type:varchar(100);unique" json:"email"`
	FirstName string    `gorm:"column:first_name;type:varchar(255);null" json:"first_name"`
	LastName  string    `gorm:"column:last_name;type:varchar(255);null" json:"last_name"`
	ImagePath string    `gorm:"type:varchar(255);null" json:"image_path"`
	Desc      string    `gorm:"type:varchar(255);null" json:"desc"`
	Hp        string    `gorm:"type:varchar(20);null" json:"hp"`
	Type      int       `gorm:"type:int(11)" json:"type"`
	TypeInfo  TypeUser  `gorm:"foreignKey:Type;references:ID" json:"type_info"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "user"
}
