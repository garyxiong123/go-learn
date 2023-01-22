package orm

import (
	"github.com/garyxiong123/go-learn/db/basic"
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
var Uuser User

func setup() {
	initDbConnection()

	intDbScheme()

	Uuser = User{
		Name: "gary",
		Age:  32,
	}

}

func initDbConnection() {
	dsn := "root:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	Db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	println(Db)
}

func intDbScheme() {
	Db.Migrator().DropTable(Person{})
	Db.AutoMigrate(Person{})

	Db.Migrator().DropTable(User{})
	Db.AutoMigrate(User{})
}

func shutdown() {

}
