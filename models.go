package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
	Token    string
}

type Music struct {
	gorm.Model
	code     uint64
	Name     string
	FileName string
}
