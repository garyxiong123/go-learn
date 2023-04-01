package json

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"
)

type Person struct {
	Name    string
	Age     *big.Int `json:",string"`
	Address string
}

func TestConvertObj2String(t *testing.T) {

	//

	//mm := big.NewInt(1234567890123456789)
	//println(mm)
	mm, _ := new(big.Int).SetString("123456789012345678900123456789001234567890", 10)
	p := Person{Name: "John Doe", Age: mm, Address: "123 Main St."}

	// Convert Person struct to JSON string
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	jsonString := string(jsonBytes)
	fmt.Println(jsonString)

	obj := convertString2Obj(jsonString)
	println(obj)
}

func convertString2Obj(objstring string) interface{} {
	//jsonString := `{"Name":"John Doe","Age":30,"Address":"123 Main St."}`

	// Convert JSON string to Person struct
	var p Person
	err := json.Unmarshal([]byte(objstring), &p)
	if err != nil {
		fmt.Println("Error:", err)
		return p
	}
	fmt.Printf("Name: %s, Age: %d, Address: %s\n", p.Name, p.Age, p.Address)
	return p
}
