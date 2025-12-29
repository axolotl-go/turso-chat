package room

import "gorm.io/gorm"

type Room struct {
	gorm.Model

	Name string `json:"name" gorm:"uniqueIndex;not null"`
}
