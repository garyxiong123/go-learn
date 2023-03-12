package equals

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

		api.AssertIsEqual(5, "33")
		api.IsZero(x3)
		m := api.Cmp(circuit.X, circuit.X)
		println(m)

	}

	return nil
}

func Test_IS_Equals(t *testing.T) {
	var circuit cubic.Circuit

	// compile a circuit
	//frontend.ToBinary(
	vr1cs, _ := frontend.Compile(ecc.BN254, r1cs.NewBuilder, &circuit)
	println(vr1cs.GetNbConstraints())
	println(vr1cs.GetNbVariables())

}
