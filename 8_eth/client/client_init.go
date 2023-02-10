package client

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var cloudClient ethclient.Client
var localClient ethclient.Client
var BscClient *ethclient.Client

func init() {
	initCloudClient()
	initLocalClient()
	initBscClient()

}

func initLocalClient() {
	localClient, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = localClient
}

func initCloudClient() {
	cloud_client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	println(cloud_client)
}

func initBscClient() {
	BscClient, _ = ethclient.Dial("https://bsc-testnet.nodereal.io/v1/a1cee760ac744f449416a711f20d99dd")
	println(BscClient)

}

func initWsClient() {

}
