package orm

import (
	"github.com/garyxiong123/go-learn/4_db/basic"
	"log"
	"testing"
)

// 手动事务
func Test_manual(t *testing.T) {
	// 开始事务
	tx := Db.Begin()

	// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
	err := tx.Create(&basic.User{Name: "toni"}).Error

	if err != nil {
		// 遇到错误时回滚事务
		tx.Rollback()
		log.Fatal(err)
	}
	// 否则，提交事务
	tx.Commit()
}

func Test_manual_error(t *testing.T) {
	// 开始事务
	tx := Db.Begin()

	// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
	err := tx.Create(&basic.User{Name: "toni"}).Error

	if err != nil {
		// 遇到错误时回滚事务
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Create(&basic.User{Name: "错误"}).Error
	if err != nil {
		// 遇到错误时回滚事务
		tx.Rollback()
		log.Fatal(err)
	}
	// 否则，提交事务
	tx.Commit()
}
