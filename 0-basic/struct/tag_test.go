package _struct

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type struct1 struct {
	id   int    "an id"
	name string "name of struct1"
}

//3. 带标签的结构体
//　　定义结构体字段时，可以给字段添加标签，并通过 reflect 获取字段对应的标签（tag）
func init() {
	st1 := struct1{12, "zhangsan"}

	str1 := reflect.TypeOf(st1)
	field1 := str1.Field(0)
	// 输出：struct1 field1 name:id, tag:an id
	fmt.Printf("struct1 field1 name:%v, tag:%v \n", field1.Name, field1.Tag)
	field2 := str1.Field(1)
	// 输出：struct1 field2 name:name, tag:name of struct1
	fmt.Printf("struct1 field2 name:%v, tag:%v \n", field2.Name, field2.Tag)
}

type Person struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name,omitempty"`
}

//json 包可以查看字段的标签并被告知如何映射 json <=> struct 字段，以及额外的选项，例如在序列化回 json 时是否应该忽略空字段。

func Test_Json(t *testing.T) {
	json_string := `
    {
        "first_name": "John",
        "last_name": "Smith"
    }`

	person := new(Person)
	json.Unmarshal([]byte(json_string), person)
	fmt.Println(person)

	new_json, _ := json.Marshal(person)
	fmt.Printf("%s\n", new_json)
}
