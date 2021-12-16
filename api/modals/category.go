package modals

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name  string `json:"name"`
	Year  uint   `json:"year"`
	Price uint   `json:"price"`
}
