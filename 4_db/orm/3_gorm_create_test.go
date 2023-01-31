package orm

import (
	"errors"
	"gorm.io/gorm"
	"testing"
)

func Test_gorm_create(t *testing.T) {

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
