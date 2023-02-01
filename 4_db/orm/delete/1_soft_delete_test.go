package orm

import (
	"fmt"
	"github.com/garyxiong123/go-learn/4_db/basic"
	"gorm.io/gorm"
	"testing"
)

// 如果您的模型包含了一个 gorm.deletedat 字段（gorm.Model 已经包含了该字段)，它将自动获得软删除的能力！
// 拥有软删除能力的模型调用 Delete 时，记录不会从数据库中被真正删除。
// 但 GORM 会将 DeletedAt 置为当前时间， 并且你不能再通过普通的查询方法找到该记录。
func Test_soft_delete_Person(t *testing.T) {
	Db.Create(&basic.Person{Name: "toni"})
	Db.Where("name = ?", "toni").Delete(&basic.Person{})

	stmt := Db.Session(&gorm.Session{DryRun: true}).Where("name = ?", "toni").Delete(&basic.Person{}).Statement
	stmt.SQL.String()
	fmt.Println(stmt.SQL.String(), stmt.Vars)
}

func Test_soft_delete_PersonDeleted(t *testing.T) {
	Db.Create(&basic.PersonDeleted{Name: "toni"})
	Db.Where("name = ?", "toni").Delete(&basic.PersonDeleted{})

	stmt := Db.Session(&gorm.Session{DryRun: true}).Where("name = ?", "toni").Delete(&basic.PersonDeleted{}).Statement
	stmt.SQL.String()
	fmt.Println(stmt.SQL.String(), stmt.Vars)
}
