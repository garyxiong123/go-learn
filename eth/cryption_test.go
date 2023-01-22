package eth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"testing"
)

func Test_eth_sigToPub(t *testing.T) {

	client, _ := rpc.Dial("http://localhost:8545")

	var pubkey string
	err := client.Call(&pubkey, "eth_sigToPub", "0x68656c6c6f20776f726c64", "0x1b20f1cfc3f8af8b1c59d4e2f9c4a8e5e4f5b5f5c53d5c8e8f0c3f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f9f")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Public key:", pubkey)
	}
}
