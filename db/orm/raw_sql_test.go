package orm

import (
	"database/sql"
	"gorm.io/gorm"
	"testing"
	"time"
)

type Result struct {
	ID   int
	Name string
	Age  int
}

func Test_query_raw(t *testing.T) {

	var result Result
	Db.Raw("SELECT id, name, age FROM users WHERE name = ?", 3).Scan(&result)

	Db.Raw("SELECT id, name, age FROM users WHERE name = ?", 3).Scan(&result)

	var age int
	Db.Raw("SELECT SUM(age) FROM users WHERE role = ?", "admin").Scan(&age)

	var users []User
	Db.Raw("UPDATE users SET name = ? WHERE age = ? RETURNING id, name", "jinzhu", 20).Scan(&users)

}

func Test_exec_raw(t *testing.T) {

	Db.Exec("DROP TABLE users")
	Db.Exec("UPDATE orders SET shipped_at = ? WHERE id IN ?", time.Now(), []int64{1, 2, 3})

	// Exec with SQL Expression
	Db.Exec("UPDATE users SET money = ? WHERE name = ?", gorm.Expr("money * ? + ?", 10000, 1), "jinzhu")

}

func Test_Named_Argument_raw(t *testing.T) {
	var user User
	var result3 Result
	Db.Where("name1 = @name OR name2 = @name", sql.Named("name", "jinzhu")).Find(&user)
	// SELECT * FROM `users` WHERE name1 = "jinzhu" OR name2 = "jinzhu"

	Db.Where("name1 = @name OR name2 = @name", map[string]interface{}{"name": "jinzhu2"}).First(&result3)
	// SELECT * FROM `users` WHERE name1 = "jinzhu2" OR name2 = "jinzhu2" ORDER BY `users`.`id` LIMIT 1

	// Named Argument with Raw SQL
	Db.Raw("SELECT * FROM users WHERE name1 = @name OR name2 = @name2 OR name3 = @name",
		sql.Named("name", "jinzhu1"), sql.Named("name2", "jinzhu2")).Find(&user)
	// SELECT * FROM users WHERE name1 = "jinzhu1" OR name2 = "jinzhu2" OR name3 = "jinzhu1"

	Db.Exec("UPDATE users SET name1 = @name, name2 = @name2, name3 = @name",
		sql.Named("name", "jinzhunew"), sql.Named("name2", "jinzhunew2"))
	// UPDATE users SET name1 = "jinzhunew", name2 = "jinzhunew2", name3 = "jinzhunew"

	Db.Raw("SELECT * FROM users WHERE (name1 = @name AND name3 = @name) AND name2 = @name2",
		map[string]interface{}{"name": "jinzhu", "name2": "jinzhu2"}).Find(&user)
	// SELECT * FROM users WHERE (name1 = "jinzhu" AND name3 = "jinzhu") AND name2 = "jinzhu2"

	type NamedArgument struct {
		Name  string
		Name2 string
	}

	Db.Raw("SELECT * FROM users WHERE (name1 = @Name AND name3 = @Name) AND name2 = @Name2",
		NamedArgument{Name: "jinzhu", Name2: "jinzhu2"}).Find(&user)
	// SELECT * FROM users WHERE (name1 = "jinzhu" AND name3 = "jinzhu") AND name2 = "jinzhu2"
}

var name = "gary"
var age = 33

func Test_sql_row(t *testing.T) {
	// Use GORM API build SQL
	row := Db.Table("users").Where("name = ?", "gary").Select("name", "age").Row()

	row.Scan(&name, &age)

	// Use Raw SQL
	//row1 := Db.Raw("select name, age, email from users where name = ?", "jinzhu").Row()
	//row1.Scan(&name, &age, &email)
}

func Test_get_result_as_sql_row(t *testing.T) {

	// Use GORM API build SQL
	rows, err := Db.Model(&User{}).Where("name = ?", "gary").Select("name, age").Rows()
	println(err)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&name, &age)

		// do something
	}

	// Raw SQL
	rows1, err := Db.Raw("select name, age from users where name = ?", "gary").Rows()
	println(err)
	defer rows1.Close()
	for rows1.Next() {
		rows1.Scan(&name, &age)
		// do something
	}
}

func Test_scan_sql_into_struct(t *testing.T) {
	rows, err := Db.Model(&User{}).Where("name = ?", "gary").Select("name, age").Rows() // (*sql.Rows, error)

	println(err)

	defer rows.Close()

	var user User
	for rows.Next() {
		// ScanRows scan a row into user
		Db.ScanRows(rows, &user)

		//println(user)
		// do something
	}

}
