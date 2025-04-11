package models

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	Alias string `gorm:"unique"`
	Root  string `gorm:"unique"`
}
