package json

import (
	"github.com/bnb-chain/zkbnb-crypto/util"
	"math/big"
	"testing"
)

func TestBig2PackedAmount(t *testing.T) {
	k, _ := new(big.Int).SetString("1000000000000000000000", 10)

	amout, _ := util.ToPackedAmount(k)

	println(amout)

}
