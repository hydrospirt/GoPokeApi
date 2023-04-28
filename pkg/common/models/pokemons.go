package models

import "gorm.io/gorm"

type Pokemon struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
