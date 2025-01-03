package models

type Address struct {
	ID       int    `gorm:"type:int(11);primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Street   string `gorm:"type:varchar(250);not null" json:"street"`
	City     string `gorm:"type:varchar(60);not null" json:"city"`
	State    string `gorm:"type:varchar(60);not null" json:"state"`
	Zipcode  string `gorm:"type:varchar(30);not null" json:"zipcode"`
	Username string `gorm:"type:varchar(100);not null" json:"username"`
	Loc      string `gorm:"type:varchar(20);null" json:"loc"`
	User     User   `gorm:"foreignKey:Username;references:Username" json:"user"`
	LocInfo  Loc    `gorm:"foreignKey:Loc;references:ID" json:"loc_info"`
}

func (Address) TableName() string {
	return "address"
}

type User struct {
	Username string `gorm:"type:varchar(100);primaryKey;" json:"username"`
}

func (User) TableName() string {
	return "user"
}
