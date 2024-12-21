package models

type User struct {
	Username string `gorm:"type:varchar(100);primaryKey;" json:"username"`
}

func (User) TableName() string {
	return "user"
}
