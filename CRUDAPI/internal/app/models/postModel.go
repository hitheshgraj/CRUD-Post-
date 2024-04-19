package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string
	Body string
}
var Post1 struct{
	Body  string
	Title string
}