package model

import "gorm.io/gorm"

// Entity
type Movie struct {
	gorm.Model
	ID int `gorm:"column:id;type:int"`
	Title string `gorm:"column:title;type:text(50)"`
	Year int `gorm:"column:year;type:int"`
	Director string `gorm:"column:director;type:text(30)"`
}
