package model

type District struct {
	ID       uint   	`gorm:"primarykey;autoIncrement"`
	Name     string 	`json:"name" gorm:"size:200;not null;"`
	CityID   uint 		`json:"city_id" gorm:"not null"`
}