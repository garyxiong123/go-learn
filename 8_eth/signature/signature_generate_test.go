package signature

import (
	"fmt"
	"github.com/consensys/gnark-crypto/ecc/bn254/twistededwards/eddsa"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"testing"

	common2 "github.com/bnb-chain/zkbnb/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func Test_ECDSA_Generate(t *testing.T) {

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	data := []byte("hello")
	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hexutil.Encode(signature)) // 0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301
}

func TestPK(t *testing.T) {
	pk := new(eddsa.PublicKey)
	//pk.A.X.SetBytes(common.FromHex("0x0dfc378542f05a22fe8d4e08ddce9b5b63eb55d1d51826f6415a492baa77429e"))
	pk.A.Y.SetBytes(common.FromHex("0x2706f1710710b1f3aa0f5e8dd6af532c1d784ebd0b9aa571c346a70bd4efa7d1"))
	fmt.Println(common.Bytes2Hex(pk.Bytes()))
}

func TestGeneratePK(t *testing.T) {
	pk, err := common2.ParsePubKey("2a991b47d608e74e00f56b7aeb15ad0349c3ec3f0de1c0cc51ba981d416f32ae")
	if err != nil {
		//logx.Errorf("unable to parse pub key: %s", err.Error())
	}
	// because we can get Y from X, so we only need to store X is enough
	fmt.Println(common.Bytes2Hex(pk.Bytes()))
}
