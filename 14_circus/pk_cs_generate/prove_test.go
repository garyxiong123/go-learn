package pk_cs_generate

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

func Test_Prove(t *testing.T) {
	var circuit cubic.Circuit

	// compile a circuit
	vr1cs, _ := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)

	// R1CS implements io.WriterTo and io.ReaderFrom
	var buf bytes.Buffer
	_, _ = vr1cs.WriteTo(&buf)

	// gnark objects (R1CS, ProvingKey, VerifyingKey, Proof) must be instantiated like so:
	newR1CS := groth16.NewCS(ecc.BN254)
	_, _ = newR1CS.ReadFrom(&buf)

	pk, vk, _ := groth16.Setup(vr1cs)

	println(pk)
	println(vk)

	// ensure prove / verify works well with valid witnesses
	valid_circuit := cubic.Circuit{
		1,
		7,
	}

	validWitness, err := frontend.NewWitness(&valid_circuit, ecc.BN254.ScalarField())

	proof, err := groth16.Prove(newR1CS, pk, validWitness)

	proof_file, err := os.Create(fmt.Sprintf("%s.txt", "proof"))
	if err != nil {
	}

	proof.WriteTo(proof_file)

	proof_file.Close()

}
