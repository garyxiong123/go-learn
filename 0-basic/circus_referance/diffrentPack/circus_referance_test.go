package circus_referance

import "testing"

//go circus reference test

type AA struct {
}

type BB struct {
}

func (A *AA) SayHello(B *BB) {
	B.SayHello(A)
	println("AA")
}

func (B *BB) SayHello(A *AA) {
	A.SayHello(B)
	println("BB")
}

func TestName(t *testing.T) {
	a := &AA{}
	b := &BB{}
	a.SayHello(b)
	b.SayHello(a)
	println("interface type test")
}
