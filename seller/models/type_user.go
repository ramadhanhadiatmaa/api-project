package models

type TypeUser struct {
	IdType int    `gorm:"type:int(11);primaryKey;autoIncrement" json:"id_type"`
	Type   string `gorm:"type:varchar(30);not null" json:"type"`
}

func (TypeUser) TableName() string {
	return "type_user"
}
