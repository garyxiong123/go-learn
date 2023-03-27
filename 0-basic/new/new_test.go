package new

import "testing"

type Person struct {
	age uint
}

func (p *Person) Person() {
	println("create new")
}

func TestNew(t *testing.T) {
	p := new(Person)
	println(p)
}
