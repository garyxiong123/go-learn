package orm

import (
	"fmt"
	"github.com/garyxiong123/go-learn/4_db/basic"
	"testing"
)

// 使用 Unscoped 找到被软删除的记录
func Test_unscoped_select_Person(t *testing.T) {
	Db.Create(&basic.Person{Name: "toni"})
	Db.Where("name = ?", "toni").Delete(&basic.Person{})
	person := &basic.Person{}
	Db.Unscoped().Where("name = ?", "toni").Find(person)
	fmt.Println(person)

}
