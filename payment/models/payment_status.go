package models

type StatusPayment struct {
	ID         int    `gorm:"primaryKey;autoIncrement" json:"id"`
	StatusName string `gorm:"type:varchar(50);not null" json:"status_name"`
}

func (StatusPayment) TableName() string {
	return "status_payment"
}
