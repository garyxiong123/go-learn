package orm

import (
	"github.com/garyxiong123/go-learn/4_db/basic"
	"testing"
)

// 开始事务
//BeforeSave
//BeforeUpdate
// 关联前的 save
// 更新 db
// 关联后的 save
//AfterUpdate
//AfterSave
// 提交或回滚事务

func Test_hook_update(t *testing.T) {
	Db.Create(&basic.Apple{Address: "bahrain"})
	Db.Model(&basic.Apple{}).Where("Address=?", "bahrain").Update("weight", 10)
}
