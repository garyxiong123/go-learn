// Copyright 2020 ConsenSys AG
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package equals

import (
	"encoding/json"
	"github.com/consensys/gnark/frontend"
	"testing"

	"github.com/consensys/gnark/test"
)

type Circuit_Byte struct {
	// struct tags on a variable is optional
	// default uses variable name and secret visibility.
	X frontend.Variable `gnark:"x"`
	Y frontend.Variable `gnark:",public"`
}

// Define declares the circuit constraints

// x**3 + x + 5 == y
//1 x2 x3

func (circuit *Circuit_Byte) Define(api frontend.API) error {
	x3 := api.Mul(circuit.X, circuit.X, circuit.X)
	// x**3 + x + 5 == y

	api.AssertIsEqual(circuit.Y, api.Add(x3, circuit.X, 5))

	by, _ := json.Marshal("sss")

	api.AssertIsEqual(x3, by)
	return nil
}

func Test_IS_Equals_Byte(t *testing.T) {

	assert := test.NewAssert(t)

	var cubicCircuit Circuit_Byte

	assert.ProverSucceeded(&cubicCircuit, &Circuit{
		X: 3,
		Y: 35,
	})

}
