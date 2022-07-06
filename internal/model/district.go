package model

type District struct {
	ID       string   	`gorm:"primarykey;size:50;"`
	Name     string 	`json:"name" gorm:"size:200;not null;"`
	CityID   string 	`json:"city_id" gorm:"not null;size:50"`
}