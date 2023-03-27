package equals

import (
	"encoding/json"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/test"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

type Circuit_Binary struct {
	// struct tags on a variable is optional
	// default uses variable name and secret visibility.
	X frontend.Variable `gnark:"x"`
}

// Define declares the circuit constraints

// x**3 + x + 5 == y
//1 x2 x3

func (circuit *Circuit_Binary) Define(api frontend.API) error {
	txFiledBits := api.ToBinary(circuit.X, 160)
	//copyLittleEndianSliceAndShiftOffset(api, txInfo.ToAddress, AddressBitsSize, &currentOffset, pubData[:])
	println(txFiledBits)

	by, _ := json.Marshal("0x4909d4D440E8ffF61738E8Cb7b2b0a4aaFF7b896")

	txFiledByte := api.ToBinary(by, 20)

	println(txFiledByte)

	return nil
}

func Test_To_Binary(t *testing.T) {

	new(big.Int).SetBytes(common.FromHex("0x4909d4D440E8ffF61738E8Cb7b2b0a4aaFF7b896")).FillBytes(make([]byte, 20))
	var circuit Circuit_Binary

	assert := test.NewAssert(t)

	assert.ProverSucceeded(&circuit, &Circuit_Binary{
		X: "0x4909d4D440E8ffF61738E8Cb7b2b0a4aaFF7b896",
	})

}
