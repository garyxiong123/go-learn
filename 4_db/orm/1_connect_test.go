package orm

import (
	"errors"
	"github.com/garyxiong123/go-learn/4_db/basic"
	"gorm.io/gorm"
	"testing"
)

func Test_connect(t *testing.T) {

	user := basic.User{
		Name: "gary",
		Age:  32,
	}
	result2 := Db.Create(&user)
	println(result2)

	result := Db.Save(user)

	result1 := Db.First(&user)
	//result1.Statement.Model
	println(result1)
	// SELECT * FROM users ORDER BY id LIMIT 1;

	// Get one record, no specified order
	Db.Take(&user)
	// SELECT * FROM users LIMIT 1;
	// Get last record, ordered by primary key desc
	Db.Last(&user)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	//result := db.First(&user)

	// check error ErrRecordNotFound
	errors.Is(result.Error, gorm.ErrRecordNotFound)

}
