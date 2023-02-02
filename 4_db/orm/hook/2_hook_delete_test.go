package orm

import (
	"github.com/garyxiong123/go-learn/4_db/basic"
	"testing"
)

// 开始事务
//BeforeDelete
// 删除 db 中的数据
//AfterDelete
// 提交或回滚事务

func Test_delete_update(t *testing.T) {
	Db.Create(&basic.Apple{Address: "bahrain"})
	Db.Model(&basic.Apple{}).Where("Address=?", "bahrain").Delete(&basic.Apple{})
}
