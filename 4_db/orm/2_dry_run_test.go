package orm

import (
	"fmt"
	"github.com/garyxiong123/go-learn/4_db/basic"
	"gorm.io/gorm"
	"testing"
)

// Generate SQL and its arguments without executing, can be used to prepare or test generated SQL,
func Test_dry_run(t *testing.T) {

	user := basic.User{}
	stmt := Db.Session(&gorm.Session{DryRun: true}).First(&user, 1).Statement
	stmt.SQL.String() //=> SELECT * FROM `users` WHERE `id` = $1 ORDER BY `id`
	//=> []interface{}{1}
	println(stmt.Vars)

}

func Test_dry_run_create(t *testing.T) {

	stmt := Db.Session(&gorm.Session{DryRun: true}).Create(&basic.User{Name: "toni"}).Statement
	stmt.SQL.String() //=> SELECT * FROM `users` WHERE `id` = $1 ORDER BY `id`
	fmt.Println(stmt.SQL.String(), stmt.Vars)
}

func Test_to_sql(t *testing.T) {
	sql := Db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(&basic.User{}).Where("id = ?", 100).Limit(10).Order("age desc").Find(&[]basic.User{})
	})

	//=> SELECT * FROM "users" WHERE id = 100 AND "users"."deleted_at" IS NULL ORDER BY age desc LIMIT 10
	println(sql)
}
