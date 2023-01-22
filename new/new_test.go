package new

import (
	"go-learn/db/orm"
	"testing"
)

// 不用输入key value 也能new
func Test_new(t *testing.T) {

	person := orm.User{
		"gary",
		33}

	println(person)

	person1 := &orm.User{
		"gary",
		33}

	println(person1)
}
