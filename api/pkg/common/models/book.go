package models

import "gorm.io/gorm"

type Book struct {
	gorm.models
	Title string `json:"title"`
	Author string `json:"author"`
	Description string `json:"description"`
}