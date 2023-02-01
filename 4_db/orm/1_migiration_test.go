package orm

import (
	"fmt"
	"github.com/garyxiong123/go-learn/4_db/basic"
	"testing"
)

func Test_AutoMigrate(t *testing.T) {

	// Get the first record ordered by primary key
	Db.Migrator().DropTable(basic.Person{})
	Db.AutoMigrate(basic.Person{})

	tx := Db.Save(person)
	fmt.Println(tx)
	fmt.Println(person)

}

func Test_update(t *testing.T) {
	Db.Save(person)
	person.Name = "zhangsan"
	Db.Updates(person)

	//Db.Save(person)
	fmt.Println(person)

}

func Test_delete(t *testing.T) {

}
