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

type Author struct {
	Name  string
	Email string
}

type Blog struct {
	ID      int
	Author  Author `gorm:"embedded;embeddedPrefix:base_"`
	Upvotes int32
}
