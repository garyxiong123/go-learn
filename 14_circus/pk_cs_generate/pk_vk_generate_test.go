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

	pk, vk, _ := groth16.Setup(vr1cs)

	validWitness, err := frontend.NewWitness(&circuit, ecc.BN254)
	proof, err := groth16.Prove(vr1cs, pk, validWitness)

	println(pk)
	println(vk)

	ccsFile, err := os.Create(fmt.Sprintf("%s.txt", "constraint"))
	if err != nil {
	}

	pk_file, err := os.Create(fmt.Sprintf("%s.txt", "pk"))
	if err != nil {
	}

	vk_file, err := os.Create(fmt.Sprintf("%s.txt", "vk"))
	if err != nil {
	}
	proof_file, err := os.Create(fmt.Sprintf("%s.txt", "proof"))
	if err != nil {
	}

	vr1cs.WriteTo(ccsFile)

	pk.WriteTo(pk_file)

	vk.WriteTo(vk_file)

	proof.WriteRawTo(proof_file)

	_ = ccsFile.Close()
	_ = pk_file.Close()
	_ = vk_file.Close()
	_ = proof_file.Close()

}
