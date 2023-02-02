package orm

import (
	"fmt"
	"github.com/garyxiong123/go-learn/4_db/basic"
	"gorm.io/gorm"
	"testing"
)

// 使用 Unscoped 永久删除匹配的记录
func Test_unscoped_delete_Person(t *testing.T) {
	Db.Create(&basic.Person{Name: "toni"})
	Db.Unscoped().Where("name = ?", "toni").Delete(&basic.Person{})

	stmt := Db.Session(&gorm.Session{DryRun: true}).Unscoped().Where("name = ?", "toni").Delete(&basic.Person{}).Statement
	stmt.SQL.String()
	fmt.Println(stmt.SQL.String(), stmt.Vars)
}
