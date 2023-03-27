package _8_zkevm

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/zksync-sdk/zksync2-go"
	"math/big"
	"testing"
)

func Test_Deposit(t *testing.T) {

	err, wallet, ep := InitWallet()

	println(wallet)

	tx, err := ep.Deposit(
		zksync2.CreateETH(),
		big.NewInt(1000000000000000),
		common.HexToAddress("0x4909d4D440E8ffF61738E8Cb7b2b0a4aaFF7b896"),
		nil,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tx hash", tx.Hash())

}

func InitWallet() (error, *zksync2.Wallet, *zksync2.DefaultEthProvider) {
	pkBytes := []byte("5a5b26d4ab5d2041b100785a3c7484e197c2346f78403faf844a43dd4be8cd34")

	ethereumSigner, err := zksync2.NewEthSignerFromRawPrivateKey(pkBytes, 280)

	// also, init ZkSync Provider, specify ZkSync2 RPC URL (e.g. testnet)
	zkSyncProvider, err := zksync2.NewDefaultProvider("https://zksync2-testnet.zksync.dev")
	println(err)

	// then init Wallet, passing just created Ethereum Signer and ZkSync Provider
	wallet, err := zksync2.NewWallet(ethereumSigner, zkSyncProvider)

	// init default RPC client to Ethereum node (Goerli network in case of ZkSync2 Era Testnet)
	ethRpc, err := rpc.Dial("https://goerli.infura.io/v3/")
	//
	//// and use it to create Ethereum Provider by Wallet
	ethereumProvider, err := wallet.CreateEthereumProvider(ethRpc)
	return err, wallet, ethereumProvider
}

func Test_Transfer(t *testing.T) {
	err, wallet, _ := InitWallet()
	hash, err := wallet.Transfer(
		common.HexToAddress("<target_L2_address>"),
		big.NewInt(1000000000000),
		nil,
		nil,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tx hash", hash)

}
