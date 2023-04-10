package convert2string

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Person_Unit64 struct {
	Name    string
	Age     int64 `json:",string"`
	Address string
}

func TestConvertObj2String_Unit64(t *testing.T) {

	//mm := big.NewInt(1234567890123456789)
	//println(mm)
	p := Person_Unit64{Name: "John Doe", Age: int64(123), Address: "123 Main St."}

	// Convert PersonBigInit struct to JSON string
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	jsonString := string(jsonBytes)
	fmt.Println(jsonString)

	obj := convertString2Obj(jsonString)
	println(obj)
	//创建随机数
	//rand.Seed(time.Now().UnixNano())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//取最大值
	//max := big.NewInt(0)
	//写单元测试
	//max.Exp(big.NewInt(2), big.NewInt(256), nil).Sub(max, big.NewInt(1))
	//println(max)
	//rand.Intn(max)
	//rand.Int63n(max)

}

func convertString2Obj(objstring string) interface{} {
	jsonString := `{"Name":"John Doe","Age":"30","Address":"123 Main St."}`

	// Convert JSON string to PersonBigInit struct
	var p Person_Unit64
	err := json.Unmarshal([]byte(jsonString), &p)
	if err != nil {
		fmt.Println("Error:", err)
		return p
	}
	fmt.Printf("Name: %s, Age: %d, Address: %s\n", p.Name, p.Age, p.Address)
	return p
}
func TestConvertString2ObjUnit64(t *testing.T) {
	jsonString := `{"Name":"John Doe","Age":"30","Address":"123 Main St."}`

	// Convert JSON string to PersonBigInit struct
	var p Person_Unit64
	err := json.Unmarshal([]byte(jsonString), &p)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Name: %s, Age: %d, Address: %s\n", p.Name, p.Age, p.Address)
}
