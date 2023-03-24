package pk

import (
	"github.com/consensys/gnark-crypto/ecc/bn254/twistededwards/eddsa"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestParsePubKey(t *testing.T) {
	//pk := new(eddsa.PublicKey)
	//pk.A.X.SetInt64(0)
	//pk.A.Y.SetInt64(0)

	//pk := &eddsa.PublicKey{
	//	A: curve.Point{
	//		X: fr.NewElement(0),
	//		Y: fr.NewElement(0),
	//	},
	//}
	//ddd := common.Bytes2Hex(pk.Bytes())
	pkBytes := common.FromHex("0000000000000000000000000000000000000000000000000000000000000000")
	pk1 := new(eddsa.PublicKey)
	_, err := pk1.A.SetBytes(pkBytes)
	if err != nil {

	}
}
