package unitTest

import (
	"errors"
	"fmt"
	"testing"
)

func Test_No_Error(t *testing.T) {

	key, error := getNoError()
	fmt.Print(key)
	fmt.Print(error)
}

func getNoError() (int, error) {
	return 3, nil
}

func Test_Error(t *testing.T) {

	key, error := getError()
	fmt.Print(key)
	fmt.Print(error)
}

func getError() (int, error) {
	error := errors.New("publicKey is not *ecdsa.PublicKey")
	return 3, error
}

//go

func Test_Error_By_fmt(t *testing.T) {

	age, err := getErrorByFmt()

	fmt.Errorf("prometheus.Register commitTreeMetics error: %v", err)

	fmt.Errorf("prometheus.Register commitTreeMetics error: %s", err)

	fmt.Print(age, err)

}

func getErrorByFmt() (int, error) {
	return 3, fmt.Errorf("ss")
}
