package A

import "github.com/garyxiong123/go-learn/basic/circus_referance/diffrentPack/B"

//go circus reference test

type AA struct {
}

func (A *AA) SayHello(B *B.BB) {
	B.SayHello(A)
	println("AA")
}
