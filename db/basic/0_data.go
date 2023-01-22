package basic

import "gorm.io/gorm"

type User struct {
	Name string
	Age  uint64
}

type Person struct {
	gorm.Model
	Name  string
	Email string
	Age   uint8
}
