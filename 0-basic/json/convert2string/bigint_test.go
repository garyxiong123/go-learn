package convert2string

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"
)

type PersonBigInit struct {
	Name    string
	Age     *big.Int `json:",string"`
	Address string
}

func TestEncode_bigint(t *testing.T) {

	//mm := big.NewInt(1234567890123456789)
	//println(mm)
	mm, _ := new(big.Int).SetString("1234567890123456", 10)
	p := PersonBigInit{Name: "John Doe", Age: mm, Address: "123 Main St."}

	// Convert PersonBigInit struct to JSON string
	jsonBytes, err := json.Marshal(p)
	json.NewDecoder()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	jsonString := string(jsonBytes)
	fmt.Println(jsonString)

	obj := convertString2Obj(jsonString)
	println(obj)
}

func TestConvertString2Obj(t *testing.T) {
	jsonString := `{"Name":"John Doe","my_key":"12345678901234567890","Address":"123 Main St."}`

	// Convert JSON string to PersonBigInit struct
	var p PersonBigInit
	err := json.Unmarshal([]byte(jsonString), &p)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Name: %s, Age: %d, Address: %s\n", p.Name, p.Age, p.Address)
}

type MyStruct struct {
	MyField big.Int `json:"my_key"`
}

func TestName(t *testing.T) {

	jsonString := `{"my_key": "12345678901234567890"}`
	myStruct := MyStruct{}
	err := json.Unmarshal([]byte(jsonString), &myStruct)
	if err != nil {
		// handle error
	}
	fmt.Println(myStruct.MyField.String()) // output: 12345678901234567890
}
