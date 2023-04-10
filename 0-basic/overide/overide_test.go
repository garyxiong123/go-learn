package overide

import "testing"

//type IBlockchain interface {
//	VerifyExpiredAt(expiredAt int64) error
//	VerifyNonce(accountIndex int64, nonce int64) error
//	VerifyGas(gasAccountIndex, gasFeeAssetId int64, txType int, gasFeeAmount *big.Int, skipGasAmtChk bool) error
//	StateDB() *sdb.StateDB
//	DB() *sdb.ChainDB
//	//CurrentBlock() *block.Block
//}

type Iname interface {
	VerifyExpiredAt(expiredAt int64) error
	VerifyExpiredAtB(expiredAt int64) error
}

type name struct {
}

func (A *name) VerifyExpiredAt(expiredAt int64) error {
	return nil
}
func (A *name) VerifyExpiredAtB(expiredAt int64) error {
	return nil
}

type nameA struct {
	*name
}

type nameB struct {
	*name
}
type nameC struct {
	*name
}

func (A *nameC) VerifyExpiredAt(expiredAt int64) error {
	println("nameC")
	return nil
}
func TestName(t *testing.T) {

	a := &nameA{}
	b := &nameB{}
	c := &nameC{}
	a.VerifyExpiredAt(1)
	b.VerifyExpiredAt(1)
	c.VerifyExpiredAt(1)
	c.VerifyExpiredAtB(1)
	sayHello(a)
	sayHello(b)
	sayHello(c)
}

// Output:
// interface type test
func sayHello(A Iname) {
	A.VerifyExpiredAt(1)
}
