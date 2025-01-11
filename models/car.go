package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	Name string
}
