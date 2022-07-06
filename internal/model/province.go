package model

type Province struct {
	ID       uint   	`gorm:"primarykey;autoIncrement"`
	Name     string 	`json:"name" gorm:"size:200;not null;"`
}