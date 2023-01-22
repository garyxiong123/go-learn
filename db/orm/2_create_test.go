package orm

import (
	"errors"
	"gorm.io/gorm"
	"testing"
)

func Test_create(t *testing.T) {

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

func Test_create_with_selected_fields(t *testing.T) {
	Db.Select("Name", "Age", "CreatedAt").Create(&Uuser)
	// INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")

}

//Create a record and ignore the values for fields passed to omit.
func Test_create_with_omit(t *testing.T) {

	Db.Omit("Name", "Age").Create(&Uuser)
	// INSERT INTO `users` (`birthday`,`updated_at`) VALUES ("2020-01-01 00:00:00.000", "2020-07-04 11:05:21.775")

}
