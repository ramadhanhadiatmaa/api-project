package models

import "time"

type Conversation struct {
	ID         int       `gorm:"type:int(11);primaryKey;" json:"id"`
	CustUser   string    `gorm:"type:varchar(100);not null" json:"cust_user"`
	SellerUser string    `gorm:"type:varchar(100);not null" json:"seller_user"`
	LastId     int       `gorm:"type:int(11);" json:"last_id"`
	CreatedAt  time.Time `json:"created_at"`

	Cust   User `gorm:"foreignKey:CustUser"`
	Seller User `gorm:"foreignKey:SellerUser"`
}

func (Conversation) TableName() string {
	return "conversation"
}
