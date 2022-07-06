package model

type City struct {
	ID       uint   	`gorm:"primarykey;autoIncrement"`
	Name     string 	`json:"name" gorm:"size:200;not null;"`
	ProvinceID   uint 	`json:"province_id" gorm:"not null"`
}