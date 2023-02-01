package basic

import "gorm.io/gorm"

type User struct {
	Name string
	Age  uint64
}

// 模型包含了一个 gorm.deletedat 字段（gorm.Model 已经包含了该字段)，它将自动获得软删除的能力
type Person struct {
	gorm.Model
	Name  string
	Email string
	Age   uint8
}

// 引入 gorm.Model，您也可以这样启用软删除特性
type PersonDeleted struct {
	ID      int
	Deleted gorm.DeletedAt
	Name    string
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
