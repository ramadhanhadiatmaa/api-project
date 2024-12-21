package models

import "time"

type Message struct {
	ID             int       `gorm:"type:int(11);primaryKey;" json:"id"`
	ConversationID int       `gorm:"type:int(11);not null" json:"conversation_id"`
	Sender         string    `gorm:"type:varchar(100);not null" json:"sender"`
	Content        string    `gorm:"type:varchar(300);not null" json:"content"`
	IsRead         bool      `json:"is_read"`
	CreatedAt      time.Time `json:"created_at"`
}

func (Message) TableName() string {
	return "message"
}
