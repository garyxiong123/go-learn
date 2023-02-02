package orm

import (
	"github.com/garyxiong123/go-learn/4_db/basic"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	//setup()
	println("exec TestMain---------")
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)

}

var Db *gorm.DB
var Uuser basic.User
var person basic.Person

func setup() {
	initDbConnection()

	intDbScheme()

	Uuser = basic.User{
		Name: "gary",
		Age:  32,
	}
	person = basic.Person{
		Age:   20,
		Email: "50222@qq.com",
		Name:  "gary"}

}

// gorm事务默认是开启的。为了确保数据一致性，Gorm会在事务里执行写入操作（增删改）。
// 如果对数据一致性要求不高的话，可以在初始化时禁用它，性能将提升大约30%。
// 一般不推荐禁用。
func initDbConnection() {
	dsn := "root:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	Db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	println(Db)
}

func intDbScheme() {
	//delete table
	Db.Migrator().DropTable(basic.Person{})
	//create table
	Db.AutoMigrate(basic.Person{})

	Db.Migrator().DropTable(basic.User{})
	Db.AutoMigrate(basic.User{})
}

func shutdown() {

}
