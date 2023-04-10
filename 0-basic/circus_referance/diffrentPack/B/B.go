package circus_referance

//go circus reference test

type BB struct {
}

func (B *BB) SayHello(A *A.AA) {
	A.SayHello(B)
	println("BB")
}
