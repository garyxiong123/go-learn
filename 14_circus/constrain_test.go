package circus

import (
	"bytes"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/garyxiong123/go-learn/circus/examples/cubic"
	"testing"
)

func Test_CirCus_constrain_number(t *testing.T) {

	var circuit cubic.Circuit
	circuit = cubic.Circuit{
		33,
		22,
	}

	// compile a circuit
	vr1cs, _ := frontend.Compile(ecc.BN254, r1cs.NewBuilder, &circuit)

	// R1CS implements io.WriterTo and io.ReaderFrom
	var buf bytes.Buffer
	_, _ = vr1cs.WriteTo(&buf)

	// gnark objects (R1CS, ProvingKey, VerifyingKey, Proof) must be instantiated like so:
	newR1CS := groth16.NewCS(ecc.BN254)
	_, _ = newR1CS.ReadFrom(&buf)

	pk, vk, _ := groth16.Setup(vr1cs)

	println(pk)
	println(vk)
}

// when will add constrain
func Test_CirCus_constrain_number_Add(t *testing.T) {

}
