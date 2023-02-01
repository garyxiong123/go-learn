package orm

import (
	"errors"
	"fmt"
	"github.com/garyxiong123/go-learn/4_db/basic"
	"gorm.io/gorm"
	"testing"
)

func Test_select(t *testing.T) {
	Db.Create(&basic.User{Name: "toni"})
	user := &basic.User{}
	tx := Db.Where("name = ?", "toni").First(user)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}
	if tx.RowsAffected == 0 {
		fmt.Println(0)
	}
	fmt.Println(user)
}

// 当 First、Last、Take 方法找不到记录时，GORM 会返回 ErrRecordNotFound 错误。
// 如果发生了多个错误，你可以通过 errors.Is 判断错误是否为 ErrRecordNotFound，例如：
func Test_exp_first(t *testing.T) {
	user := &basic.User{}
	tx := Db.Table("users").Where("name = ?", "toni").First(user)
	fmt.Println(errors.Is(tx.Error, gorm.ErrRecordNotFound))
}

func Test_exp_find(t *testing.T) {
	user := &basic.User{}
	tx := Db.Table("users").Where("name = ?", "toni").Find(user)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	} else if tx.RowsAffected == 0 {
		fmt.Println(0)
	}
	fmt.Println(user)
}
