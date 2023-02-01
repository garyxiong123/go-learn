package orm

import (
	"fmt"
	"github.com/garyxiong123/go-learn/4_db/basic"
	"testing"
	"time"
)

func Test_skip_transaction(t *testing.T) {
	start := time.Now()
	for i := 0; i < 5000; i++ {
		Db.Create(&basic.User{Name: "toni"})
	}
	fmt.Print(float64(time.Since(start).Milliseconds()))
}
