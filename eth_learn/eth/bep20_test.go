package eth

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/garyxiong123/go-learn/eth_learn_eth/bep20"
	"github.com/garyxiong123/go-learn/eth_learn_eth/util"
	"log"
	"math/big"
	"testing"
)

func Test_Deploy_Bep20(t *testing.T) {
	transactOpts, err := BuildDefaultTransactOpts()
	if err != nil {
	}
	client, err := ethclient.Dial("https://bsc-testnet.nodereal.io/v1/a1cee760ac744f449416a711f20d99dd")
	var InitialSupply = util.ToIntByPrecise("5", 22)
	addr, tx, bep20, err := bep20.DeployBep20(transactOpts, client, InitialSupply, "name-gary", "symbol-gary")

	//transfer from to
	transactOpts1, err := BuildDefaultTransactOpts()
	tx, err = bep20.Transfer(transactOpts1, common.HexToAddress("0x4909d4D440E8ffF61738E8Cb7b2b0a4aaFF7b896"), util.ToIntByPrecise("5", 18))
	if err != nil {

	}

	//bep20.Approve()
	log.Print(tx)
	log.Print(addr)
	log.Print(tx)
	log.Print(bep20)
	log.Print(err)

}

func BuildDefaultTransactOpts() (ops *bind.TransactOpts, err error) {

	client, err := ethclient.Dial("https://bsc-testnet.nodereal.io/v1/a1cee760ac744f449416a711f20d99dd")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("3721e672963d8fb16f6146f091e12fca4ee09a3b47235d45370b4671943d3fac")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice
	return auth, nil

}
