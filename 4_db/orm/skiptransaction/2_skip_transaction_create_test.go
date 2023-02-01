package orm

import (
	"fmt"
	"github.com/garyxiong123/go-learn/4_db/basic"
	"testing"
)

func Test_skip_autocommit(t *testing.T) {
	var value []*basic.User
	value = append(value, &basic.User{Name: "toni1"})
	value = append(value, &basic.User{Name: "toni2"})
	Db.CreateInBatches(value, len(value))
}

func Test_skip_autocommit_error(t *testing.T) {
	var value []*basic.User
	value = append(value, &basic.User{Name: "toni1"})
	value = append(value, &basic.User{Name: "错误"})
	Db.CreateInBatches(value, len(value))
}

func Test_skip_autocommit_error_size(t *testing.T) {
	var value []*basic.User
	value = append(value, &basic.User{Name: "toni1"})
	value = append(value, &basic.User{Name: "toni2"})
	value = append(value, &basic.User{Name: "错误"})
	tx := Db.CreateInBatches(value, 1)
	fmt.Println(tx.Error)
}
