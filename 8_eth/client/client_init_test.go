package client

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
)

//若您没有现有以太坊客户端，您可以连接到infura网关。Infura管理着一批安全，可靠，可扩展的以太坊[geth和parity]节点
func Test_Init_Client_Infura(t *testing.T) {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client
	gasprice, err := client.SuggestGasPrice(context.Background())

	if err != nil {

	}
	println(gasprice)

	networkId, err := client.NetworkID(context.Background())
	println(networkId)
	if err != nil {

	}

}

func Test_Init_Client_Local_IPC(t *testing.T) {
	client, err := ethclient.Dial("/home/user/.ethereum/geth.ipc")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client
}

//makefile 中启动ganache
func Test_Init_Client_Ganache(t *testing.T) {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client
}
