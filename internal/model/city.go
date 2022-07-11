package model

type City struct {
	ID       	 string   		`gorm:"primarykey;size:50;"`
	Name     	 string 		`json:"name" gorm:"size:200;not null;"`
	ProvinceID   string 		`json:"province_id" gorm:"not null;size:50"`
}