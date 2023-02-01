package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_2_json(t *testing.T) {

	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}

	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)

	}

	println(b)

	fmt.Println(string(b))
	var col = &ColorGroup{}
	json.Unmarshal(b, col)
	fmt.Println(*col)

}
