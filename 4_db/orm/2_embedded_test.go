package orm

import (
	"testing"
)

type Author struct {
	Name  string
	Email string
}

type Blog struct {
	ID      int
	Author  Author `gorm:"embedded"`
	Upvotes int32
}

func Test_embedded(t *testing.T) {

	Db.Save(Uuser)

}
