package pk_cs_generate

import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"testing"
)

type CircuitVariable struct {
	// struct tags on a variable is optional
	// default uses variable name and secret visibility.
	X frontend.Variable   `gnark:"x"`
	Y frontend.Variable   `gnark:",public"`
	M []frontend.Variable `gnark:",public"`
}

// Define declares the circuit constraints

// x**3 + x + 5 == y
//1 x2 x3

func (circuit *CircuitVariable) Define(api frontend.API) error {

	for i := 0; i < cap(circuit.M); i++ {
		x3 := api.Mul(circuit.X, circuit.X, circuit.X)
		// x**3 + x + 5 == y
		api.IsZero(circuit.M[i])
		api.AssertIsEqual(circuit.Y, api.Add(x3, circuit.X, 5))
	}

	return nil
}

/**
* 加上 For 循环 参数变量， 但是
 */
func Test_Constrain_Add_Variable_withError(t *testing.T) {
	var circuitAddVariable CircuitVariable

	// compile a circuit
	vr1cs_Add_For, err := frontend.Compile(ecc.BN254, r1cs.NewBuilder, &circuitAddVariable)
	println(err)
	println(vr1cs_Add_For.GetNbConstraints())
	println(vr1cs_Add_For.GetNbVariables()) //内部变量增加

}

// 必须确定  数组的个数
func Test_Constrain_Add_Variable_Fix_Error(t *testing.T) {
	var circuitAddVariable CircuitVariable

	circuitAddVariable.M = make([]frontend.Variable, 2)
	// compile a circuit
	vr1cs_Add_Variable, err := frontend.Compile(ecc.BN254, r1cs.NewBuilder, &circuitAddVariable)
	println(err)
	println(vr1cs_Add_Variable.GetNbConstraints())
	println(vr1cs_Add_Variable.GetNbVariables()) //内部变量增加

}
