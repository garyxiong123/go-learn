package orm

import (
	"errors"
	"github.com/garyxiong123/go-learn/4_db/basic"
	"gorm.io/gorm"
	"testing"
)

// 嵌套事务
// SavePoint会触发
// 生成point名称规则 fmt.Sprintf("sp%p", fc)
// 默认情况下会对嵌套事务的SavePoint采用自动命名，例如github.com/garyxiong123/go-learn/4_db/orm/transaction.Test_autocommit_nested.func1.2
// 其中的N表示嵌套的层级数量，如果您看到日志中出现func1.2 表示当前嵌套层级为1（从0开始计算）

func Test_autocommit_nested_0_error(t *testing.T) {
	Db.Transaction(func(tx *gorm.DB) error {
		tx.Create(&basic.User{Name: "toni1"})

		//1.1
		tx.Transaction(func(tx2 *gorm.DB) error {
			tx2.Create(&basic.User{Name: "toni2"})
			return nil
		})

		//1.2
		tx.Transaction(func(tx2 *gorm.DB) error {
			tx2.Create(&basic.User{Name: "toni3"})
			return errors.New("rollback toni2") // Rollback user2
		})

		return nil
	})
}

func Test_autocommit_nested_1_1_error(t *testing.T) {
	Db.Transaction(func(tx *gorm.DB) error {
		tx.Create(&basic.User{Name: "toni1"})

		//1.1
		tx.Transaction(func(tx2 *gorm.DB) error {
			tx2.Create(&basic.User{Name: "toni2"})
			return errors.New("rollback toni2") // Rollback user2
		})

		//1.2
		tx.Transaction(func(tx2 *gorm.DB) error {
			tx2.Create(&basic.User{Name: "toni3"})
			return nil
		})

		return nil
	})
}

func Test_autocommit_nested_1_2_error(t *testing.T) {
	Db.Transaction(func(tx *gorm.DB) error {
		tx.Create(&basic.User{Name: "toni1"})

		//1.1
		tx.Transaction(func(tx2 *gorm.DB) error {
			tx2.Create(&basic.User{Name: "toni2"})
			return nil
		})

		//1.2
		tx.Transaction(func(tx2 *gorm.DB) error {
			tx2.Create(&basic.User{Name: "toni3"})
			return errors.New("rollback toni2") // Rollback user2
		})

		return nil
	})
}
