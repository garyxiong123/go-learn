package suit

import (
	"testing"
)

// share global variable
var age uint

// TestSuite is a test suite that groups together related tests.
func TestSuite(t *testing.T) {
	t.Run("TestFoo", TestFoo)
	t.Run("TestBar", TestBar)
}

// TestFoo is a test function that tests the Foo function.
func TestFoo(t *testing.T) {
	result := Foo()
	age++
	println(age)
	if result != "foo" {
		t.Errorf("Foo() = %s; expected 'foo'", result)
	}
}

func Foo() interface{} {
	return "foo"
}

// TestBar is a test function that tests the Bar function.
func TestBar(t *testing.T) {
	result := Bar()
	age++
	println(age)

	if result != "bar" {
		t.Errorf("Bar() = %s; expected 'bar'", result)
	}
}

func Bar() interface{} {
	return "bar"
}
