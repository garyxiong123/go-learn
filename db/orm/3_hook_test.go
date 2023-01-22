package orm

import (
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func Test_hook(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	type User struct {
		Name string
		Age  uint64
	}
	// Get the first record ordered by primary key
	var user User

	user = User{
		Name: "gary",
		Age:  32,
	}
	db.Save(user)

	db.First(&user)
	// SELECT * FROM users ORDER BY id LIMIT 1;

	// Get one record, no specified order
	db.Take(&user)
	// SELECT * FROM users LIMIT 1;
	// Get last record, ordered by primary key desc
	db.Last(&user)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	result := db.First(&user)

	// check error ErrRecordNotFound
	errors.Is(result.Error, gorm.ErrRecordNotFound)

}

//func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
//	u.UUID = uuid.New()
//
//	if u.Role == "admin" {
//		return errors.New("invalid role")
//	}
//	return
//}
