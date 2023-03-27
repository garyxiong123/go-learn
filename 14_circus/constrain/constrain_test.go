package pk_cs_generate

import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/garyxiong123/go-learn/circus/examples/cubic"

	"testing"
)

type CircuitAddFor struct {
	// struct tags on a variable is optional
	// default uses variable name and secret visibility.
	X frontend.Variable `gnark:"x"`
	Y frontend.Variable `gnark:",public"`
}

// Define declares the circuit constraints

// x**3 + x + 5 == y
//1 x2 x3

func (circuit *CircuitAddFor) Define(api frontend.API) error {
	for i := 0; i < 5; i++ {
		x3 := api.Mul(circuit.X, circuit.X, circuit.X)
		// x**3 + x + 5 == y
		api.AssertIsEqual(circuit.Y, api.Add(x3, circuit.X, 5))
	}

	return nil
}

func Test_Constrain_Add_Factor(t *testing.T) {
	var circuit cubic.Circuit

	// compile a circuit
	//frontend.ToBinary(
	vr1cs, _ := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	println(vr1cs.GetNbConstraints())
	println(vr1cs.GetNbVariables())
	var circuitAddFor CircuitAddFor

	// compile a circuit
	vr1cs_Add_For, _ := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuitAddFor)
	println(vr1cs_Add_For.GetNbConstraints())
	println(vr1cs_Add_For.GetNbVariables()) //内部变量增加

}
