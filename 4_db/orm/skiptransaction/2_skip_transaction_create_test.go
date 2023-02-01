package orm

import (
	"errors"
	"github.com/garyxiong123/go-learn/4_db/basic"
	"gorm.io/gorm"
	"testing"
)

func Test_skip_autocommit(t *testing.T) {
	Db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(&basic.User{Name: "toni1"}).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		if err := tx.Create(&basic.User{Name: "toni2"}).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
}

func Test_skip_autocommit_error(t *testing.T) {
	Db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(&basic.User{Name: "toni1"}).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		if err := tx.Create(&basic.User{Name: "toni2"}).Error; err == nil {
			return errors.New("rollback toni1")
		}

		// 返回 nil 提交事务
		return nil
	})
}
