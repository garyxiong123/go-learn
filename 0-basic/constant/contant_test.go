package constant

import "testing"

type AbiId uint64

const (
	DefaultAbi AbiId = iota
	TransferAbi
	WithdrawAbi
	CreateCollectionAbi
	WithdrawNftAbi
	TransferNftAbi
	MintNftAbi
	AtomicMatchAbi
	CancelOfferAbi
)

func Test_Constant(t *testing.T) {
	println(DefaultAbi)
	println(TransferAbi)
}
