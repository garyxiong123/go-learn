package client

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var cloudClient ethclient.Client
var localClient ethclient.Client

func init() {
	initCloudClient()
	initLocalClient()

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
