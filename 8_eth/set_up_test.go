package __eth

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"testing"
)

//Go testing package also has a way to run function before and after
//each test function, Setup and Teardown function respectively and it can
//be used to do tasks like creating a temporary database, opening a file,
//starting a server and clean up the resources after each test.

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	println(code)
	tearDown()
}

var (
	client      *ethclient.Client
	privateKey  *ecdsa.PrivateKey
	publicKey   *ecdsa.PublicKey
	fromAddress common.Address
)

func setUp() {
	println("gary ---- setUp-----")

	client, err := ethclient.Dial("http://rinkeby.infura.io")

	println(client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Println(fromAddress)

}

func tearDown() {
	println("gary ---- tearDown-----")
}
