package models

type OrderStatus struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id_status"`
	StatusName string `gorm:"type:varchar(50);not null" json:"status_name"`
}

func (OrderStatus) TableName() string {
	return "order_status"
}
