package model

import (
	"time"

	"gorm.io/gorm"
)

type UserToken struct {
	ID       uint   	`gorm:"primarykey;autoIncrement"`
	Email    string `json:"email" gorm:"size:200;not null;unique"`
	Token    string 	`json:"token" gorm:"size:200;not null"`
	Type 	 string 	`json:"type" gorm:"size:50;not null"`
	CreatedAt time.Time
}

// BeforeCreate is a method for struct User
// gorm call this method before they execute query
func (u *UserToken) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	return
}
