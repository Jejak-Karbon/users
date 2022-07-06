package model

type Province struct {
	ID       string   	`gorm:"primarykey;size:50;"`
	Name     string 	`json:"name" gorm:"size:200;not null;"`
}