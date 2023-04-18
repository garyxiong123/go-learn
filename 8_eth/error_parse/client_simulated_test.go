package main

import (
	"context"
	"fmt"
	zkbnb "github.com/bnb-chain/zkbnb-eth-rpc/core"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"strings"
	"testing"
)

// github.com
// /bnb-chain/zkbnb-eth-rpc => github.com/15000785133/zkbnb-eth-rpc v0.0.3-0.20230403192823-23498c95979c
func TestName(t *testing.T) {

	client, err := ethclient.Dial("https://bsc-testnet.nodereal.io/v1/a1cee760ac744f449416a711f20d99dd")
	if err != nil {
		panic(err)
	}

	//https://testnet.bscscan.com/tx/0xcb71039c70553fa99ab0e45082b115a38292f36d9c5c16ade6e2ced5e59a42c6
	txHash := common.HexToHash("0xcb71039c70553fa99ab0e45082b115a38292f36d9c5c16ade6e2ced5e59a42c6")
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		panic(err)
	}

	if receipt.Status == 0 {
		fmt.Println("Transaction was reverted")
		ZkBNBContractAbi, _ := abi.JSON(strings.NewReader(zkbnb.ZkBNBMetaData.ABI))

		//// Get the contract ABI
		//contractAbi, err := abi.JSON(strings.NewReader(ZkBNBContractAbi))
		//if err != nil {
		//	panic(err)
		//}

		// Get the function signature of the failed function call
		signature := receipt.Logs[0].Topics[0].Hex()

		// Get the function inputs
		inputData := receipt.Logs[0].Data

		// Parse the function signature and input data
		parsed, err := ZkBNBContractAbi.Unpack(signature, inputData)
		if err != nil {
			panic(err)
		}

		// Get the revert reason string
		reason := parsed[0].(string)
		fmt.Println("Revert reason:", reason)
	} else {
		fmt.Println("Transaction was successful")
	}
}
