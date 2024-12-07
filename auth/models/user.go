package models

type User struct {
	Username string   `gorm:"type:varchar(60);primaryKey;" json:"username"`
	Password string   `gorm:"type:varchar(60)" json:"password"`
	Type     int      `gorm:"type:int(11)" json:"type"`
	TypeInfo TypeUser `gorm:"foreignKey:Type;references:IdType" json:"type_info"`
}

type TypeUser struct {
	IdType int    `gorm:"type:int(11);primaryKey;autoIncrement" json:"id_type"`
	Type   string `gorm:"type:varchar(30);not null" json:"type"`
}

func (User) TableName() string {
	return "user"
}

func (TypeUser) TableName() string {
	return "type_user"
}
