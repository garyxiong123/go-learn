package circus

import (
	"bytes"
	"fmt"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/garyxiong123/go-learn/circus/examples/cubic"
	"os"
	"testing"
)

func Test_Pk_Vk_Generate(t *testing.T) {

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

	//var pk groth16_bn254.ProvingKey
	//var vk groth16_bn254.VerifyingKey
	// setup
	pk, vk, _ := groth16.Setup(vr1cs)

	println(pk)
	println(vk)

	ccsFile, err := os.Create(fmt.Sprintf("%s.ccs.save", "session"))
	if err != nil {
	}

	vr1cs.WriteTo(ccsFile)

	_ = ccsFile.Close()

}
