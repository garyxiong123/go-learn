package extend

import "testing"

type TxExecutor interface {
	Prepare() error
	VerifyInputs() error
	ApplyTransaction() error
	GeneratePubData() error
}

type Animal interface {
	Prepare() error
}

/**
没有显示的引用 多个同名方法  确定 很难区分，容易乱
*/
type BaseExecutor struct {
	name string
}

func (b BaseExecutor) Prepare() error {
	println(b.name)
	return nil
}

func (b BaseExecutor) VerifyInputs() error {
	panic("implement me")
}

func (b BaseExecutor) ApplyTransaction() error {
	panic("implement me")
}

func (b BaseExecutor) GeneratePubData() error {
	panic("implement me")
}

type mintNftExecutor struct {
	BaseExecutor
}

type rmintNftExecutor struct {
	*BaseExecutor
}

func Test_extend(t *testing.T) {
	txExecutor := &mintNftExecutor{}
	txExecutor.Prepare()
	println(txExecutor)
}

func Test_extend_referance(t *testing.T) {
	println()
}
