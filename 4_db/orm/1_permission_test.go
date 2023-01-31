package orm

import (
	"fmt"
	"testing"
)

func Test_create11(t *testing.T) {

	type User struct {
		Name1  string `gorm:"<-:create"`          // allow read and create
		Name2  string `gorm:"<-:update"`          // allow read and update
		Name3  string `gorm:"<-"`                 // allow read and write (create and update)
		Name4  string `gorm:"<-:false"`           // allow read, disable write permission
		Name5  string `gorm:"->"`                 // readonly (disable write permission unless it configured)
		Name6  string `gorm:"->;<-:create"`       // allow read and create
		Name7  string `gorm:"->:false;<-:create"` // createonly (disabled read from db)
		Name8  string `gorm:"-"`                  // ignore this field when write and read with struct
		Name9  string `gorm:"-:all"`              // ignore this field when write, read and migrate with struct
		Name10 string `gorm:"-:migration"`        // ignore this field when migrate with struct
	}
	// Get the first record ordered by primary key
	var user User
	fmt.Println(user)
	//println(user)
	//
	//user = User{
	//	Name: "gary",
	//	Age:  32,
	//}
	//db.Save(user)
	//
	//db.First(&user)
	//// SELECT * FROM users ORDER BY id LIMIT 1;
	//
	//// Get one record, no specified order
	//db.Take(&user)
	//// SELECT * FROM users LIMIT 1;
	//// Get last record, ordered by primary key desc
	//db.Last(&user)
	//// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	//
	//result := db.First(&user)
	//
	//// check error ErrRecordNotFound
	//errors.Is(result.Error, gorm.ErrRecordNotFound)

}
