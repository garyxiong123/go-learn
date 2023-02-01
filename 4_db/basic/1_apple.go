package basic

import (
	"fmt"
	"gorm.io/gorm"
)

type Apple struct {
	Address string
	Weight  uint64
}

func (u *Apple) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeCreate")
	return
}

func (u *Apple) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("AfterCreate")
	return
}

func (u *Apple) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeUpdate")
	return
}

// 在同一个事务中更新数据
func (u *Apple) AfterUpdate(tx *gorm.DB) (err error) {
	fmt.Println("AfterUpdate")
	return
}

// 在同一个事务中更新数据
func (u *Apple) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println("AfterDelete")
	return
}

func (u *Apple) AfterFind(tx *gorm.DB) (err error) {
	fmt.Println("AfterFind")
	return
}
