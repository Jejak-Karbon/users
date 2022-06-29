package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	CreatedAt time.Time
	CreatedBy uint
	UpdatedAt time.Time
	UpdatedBy uint
	DeletedAt gorm.DeletedAt `gorm:"index"`
	DeletedBy uint
}
