package models

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"publication"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}
