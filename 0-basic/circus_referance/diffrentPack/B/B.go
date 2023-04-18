package B

import "github.com/garyxiong123/go-learn/basic/circus_referance/diffrentPack/A"

//go circus reference test

type BB struct {
}

func (B *BB) SayHello(A *A.AA) {
	//A.SayHello(B)
	println("BB")
}
