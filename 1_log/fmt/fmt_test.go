package fmt

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T) {

	fmt.Errorf("sss:%s", "33")
	fmt.Errorf("sss:%d", 3)
	//fmt.Errorf("sss:%d", "3") //Errorf format %d has arg "3" of wrong type string

}

type Animal struct {
	name string
	age  int
}

//%v	the value in a default format
//when printing structs, the plus flag (%+v) adds field names
//%#v	a Go-syntax representation of the value
//%T	a Go-syntax representation of the type of the value
//%%	a literal percent sign; consumes no value

func Test_type_general(t *testing.T) {
	cat := Animal{
		name: "gary",
		age:  33,
	}

	fmt.Println("the value in a default format:%T", cat)
}

func Test_type_boolean(t *testing.T) {

}

/**
%b	base 2
%c	the character represented by the corresponding Unicode code point
%d	base 10
%o	base 8
%O	base 8 with 0o prefix
%q	a single-quoted character literal safely escaped with Go syntax.
%x	base 16, with lower-case letters for a-f
%X	base 16, with upper-case letters for A-F
%U	Unicode format: U+1234; same as "U+%04X"
*/
func Test_type_Integer(t *testing.T) {

}
