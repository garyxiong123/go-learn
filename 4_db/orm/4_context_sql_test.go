package orm

import (
	"context"
	"fmt"
	"github.com/garyxiong123/go-learn/4_db/basic"
	"testing"
	"time"
)

// 对于长Sql查询，你可以传入一个带超时的 context 给 db.WithContext 来设置超时时间，例如：
func Test_context_sql(t *testing.T) {

	Db.Create(&basic.User{Name: "toni"})
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	user := &basic.User{}
	Db.WithContext(ctx).Find(user)
	fmt.Println(user)
}

func Test_context_sql_timeout(t *testing.T) {

	Db.Create(&basic.User{Name: "toni"})
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Microsecond)
	defer cancel()
	user := &basic.User{}
	tx := Db.WithContext(ctx).Find(user)
	fmt.Println(tx.Error)
	fmt.Println(user)
}
