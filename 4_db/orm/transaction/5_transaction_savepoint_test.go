package orm

import (
	"github.com/garyxiong123/go-learn/4_db/basic"
	"testing"
)

func Test_point_back(t *testing.T) {
	tx := Db.Begin()
	tx.Create(&basic.User{Name: "toni"})

	tx.SavePoint("point")
	tx.Create(&basic.User{Name: "ton2"})
	tx.RollbackTo("point") // Rollback user2

	tx.Commit() // Commit user1
}
