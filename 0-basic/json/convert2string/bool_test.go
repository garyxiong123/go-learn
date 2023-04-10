package convert2string

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

type parsedBool bool

func (basi *parsedBool) UnmarshalJSON(data []byte) error {
	v, err := strconv.ParseBool(string(data))
	if err != nil {
		return err
	}
	*basi = parsedBool(v)
	return nil
}

func (basi parsedBool) MarshalJSON() ([]byte, error) {
	if basi {
		return []byte(`"1"`), nil // here ADDED BY TOBY
	} else {
		return []byte(`"0"`), nil // and here ADDED BY TOBY
	}
}
func TestBool(t *testing.T) {

	wrapped := struct {
		P parsedBool `json:",string"`
		B bool       `json:",string"`
	}{
		P: true,
		B: true,
	}
	data, _ := json.Marshal(wrapped)
	fmt.Printf("JSON marshalled: %+v\n", string(data))
	_ = json.Unmarshal([]byte(`{"P":"1","B":"true"}`), &wrapped)
	fmt.Printf("Unmarshalled value: %+v\n", wrapped)
}
