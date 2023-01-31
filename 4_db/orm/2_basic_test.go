package orm

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"testing"
)

func Test_basic(t *testing.T) {

	Db.Save(Uuser)

	Db.First(&Uuser)
	// SELECT * FROM users ORDER BY id LIMIT 1;

	// Get one record, no specified order
	Db.Take(&Uuser)
	// SELECT * FROM users LIMIT 1;
	// Get last record, ordered by primary key desc
	Db.Last(&Uuser)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	result := Db.First(&Uuser)

	// check error ErrRecordNotFound
	errors.Is(result.Error, gorm.ErrRecordNotFound)

}

func Test_double_create(t *testing.T) {
	Db.Save(Uuser)
	Db.Save(Uuser)
	//sql input will be two  as no uniqe key
}
func Test_double_create_with_point(t *testing.T) {
	Db.Save(&Uuser)
	Db.Save(&Uuser)
	//sql input will be two  as no uniqe key
}

func Test_double_create_with_unique_key_without_point(t *testing.T) {
	Db.Save(person)
	//error happen
}

func Test_double_create_with_unique_key(t *testing.T) {
	Db.Save(&person)
	fmt.Println(person)
	person.Name = "mark"
	Db.Save(&person)
	//sql input will be two  as no uniqe key
}
func Test_create_update(t *testing.T) {
	Db.Save(Uuser)
	Uuser.Age = 31
	Db.Save(Uuser)
	//sql input must be one
}

func Test_create_with_selected_fields(t *testing.T) {
	Db.Select("Name", "Age", "CreatedAt").Create(&Uuser)
	// INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")

}

//Create a record and ignore the values for fields passed to omit.
func Test_create_with_omit(t *testing.T) {

	Db.Omit("Name", "Age").Create(&Uuser)
	// INSERT INTO `users` (`birthday`,`updated_at`) VALUES ("2020-01-01 00:00:00.000", "2020-07-04 11:05:21.775")

}
