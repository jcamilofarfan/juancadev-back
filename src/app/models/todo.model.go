package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Task      string `gorm:"not null"`
	Completed bool   `gorm:"default:false"`
	User      *uint  `gorm:"not null" gorm:"index"`
}