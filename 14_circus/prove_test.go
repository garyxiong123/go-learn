package circus

import (
	"bytes"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/test"
	"github.com/garyxiong123/go-learn/circus/examples/cubic"
	"testing"
)

func Test_Prove(t *testing.T) {
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
	assert := test.NewAssert(t)

	// ensure prove / verify works well with valid witnesses
	validWitness, err := frontend.NewWitness(&circuit, ecc.BN254)

	proof, err := groth16.Prove(newR1CS, pk, validWitness)
	checkError(err)

	validPublicWitness, err := frontend.NewWitness(&circuit, ecc.BN254, frontend.PublicOnly())
	err = groth16.Verify(proof, vk, validPublicWitness)
	checkError(err)
	assert.ProverSucceeded(&circuit, &cubic.Circuit{
		X: 42,
		Y: 42,
	})

}

func checkError(err error) {

}

func Test_GenerateProve(t *testing.T) {

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
