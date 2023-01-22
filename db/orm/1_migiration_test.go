package orm

import (
	"testing"
)

var person = &Person{
	Name:  "gary",
	Email: "502238415@qq.com",
	Age:   33,
}

func Test_AutoMigrate(t *testing.T) {

	// Get the first record ordered by primary key
	Db.Migrator().DropTable(Person{})
	Db.AutoMigrate(Person{})

	tx := Db.Save(person)
	println(tx)
	println(person)

}

func Test_update(t *testing.T) {
	Db.Save(person)
	person.Name = "zhangsan"
	Db.Updates(person)

	//Db.Save(person)
	println(person)

}

func Test_delete(t *testing.T) {

}
