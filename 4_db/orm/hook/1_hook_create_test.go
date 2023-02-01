package orm

import (
	"github.com/garyxiong123/go-learn/4_db/basic"
	"testing"
)

// 开始事务
//BeforeSave
//BeforeCreate
// 关联前的 save
// 插入记录至 db
// 关联后的 save
//AfterCreate
//AfterSave
// 提交或回滚事务

func Test_hook_create(t *testing.T) {
	Db.Create(&basic.Apple{Address: "bahrain"})
}
