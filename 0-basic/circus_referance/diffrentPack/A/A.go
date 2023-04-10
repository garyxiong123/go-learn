package circus_referance

//go circus reference test

type AA struct {
}

func (A *AA) SayHello(B *BB) {
	B.SayHello(A)
	println("AA")
}
