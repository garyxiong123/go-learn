package _8_zkevm

import (
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/zksync-sdk/zksync2-go"
	"testing"
)

func Test_Init(t *testing.T) {
	// first, init Ethereum Signer, from your mnemonic, and with the chain Id (in zkSync Era Testnet case, 280)
	//ethereumSigner, err := zksync2.NewEthSignerFromMnemonic("<mnemonic words>", 280)
	//
	//println(err)

	pkBytes := []byte("5a5b26d4ab5d2041b100785a3c7484e197c2346f78403faf844a43dd4be8cd34")

	// or from raw PrivateKey bytes
	ethereumSigner, err := zksync2.NewEthSignerFromRawPrivateKey(pkBytes, 280)

	// also, init ZkSync Provider, specify ZkSync2 RPC URL (e.g. testnet)
	zkSyncProvider, err := zksync2.NewDefaultProvider("https://zksync2-testnet.zksync.dev")
	println(err)

	// then init Wallet, passing just created Ethereum Signer and ZkSync Provider
	wallet, err := zksync2.NewWallet(ethereumSigner, zkSyncProvider)

	println(wallet)

	// init default RPC client to Ethereum node (Goerli network in case of ZkSync2 Era Testnet)
	ethRpc, err := rpc.Dial("https://goerli.infura.io/v3/<your_infura_node_id>")
	//
	//// and use it to create Ethereum Provider by Wallet
	ethereumProvider, err := wallet.CreateEthereumProvider(ethRpc)

	println(ethereumProvider)

}
