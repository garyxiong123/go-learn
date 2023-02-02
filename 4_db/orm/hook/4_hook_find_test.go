package orm

import (
	"github.com/garyxiong123/go-learn/4_db/basic"
	"testing"
)

// 从 db 中加载数据
// Preloading (eager loading)
//AfterFind

func Test_find_update(t *testing.T) {
	Db.Create(&basic.Apple{Address: "bahrain"})
	Db.Model(&basic.Apple{}).Where("Address=?", "bahrain").Find(&basic.Apple{})
}
