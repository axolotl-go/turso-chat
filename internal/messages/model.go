package messages

import "gorm.io/gorm"

type Message struct {
	gorm.Model

	Sender  string `json:"sender" gorm:"not null"`
	RoomID  uint   `json:"room_id" gorm:"index;not null"`
	Content string `json:"content" gorm:"type:text;not null"`
}
