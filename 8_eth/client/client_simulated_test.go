package client

//
//import (
//	"context"
//	"fmt"
//	"log"
//	"math/big"
//	"testing"
//
//	"github.com/ethereum/go-ethereum/accounts/abi/bind"
//	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/core"
//	"github.com/ethereum/go-ethereum/core/types"
//	"github.com/ethereum/go-ethereum/crypto"
//)
//
////使用模拟客户端
//
////快速轻松地在本地测试您的交易，非常适合单元测试。为了开始，我们需要一个带有初始ETH的账户。为此，首先生成一个账户私钥。
//func TestSimulated(t *testing.T) {
//
//	//有初始ETH的账户。为此，首先生成一个账户私钥
//	privateKey, err := crypto.GenerateKey()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	//接着从accounts/abi/bind包创建一个NewKeyedTransactor，并为其传递私钥。
//	auth := bind.NewKeyedTransactor(privateKey)
//
//	//下一步是创建一个创世账户并为其分配初始余额。我们将使用core包的GenesisAccount类型。
//	balance := new(big.Int)
//	balance.SetString("10000000000000000000", 10) // 10 eth in wei
//
//	address := auth.From
//	genesisAlloc := map[common.Address]core.GenesisAccount{
//		address: {
//			Balance: balance,
//		},
//	}
//
//	blockGasLimit := uint64(4712388)
//	//现在我们将创世分配结构体和配置好的汽油上限传给account/abi/bind/backends包的NewSimulatedBackend方法，
//	//该方法将返回一个新的模拟以太坊客户端。
//	client := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)
//
//	fromAddress := auth.From
//
//	//新的交易并进行广播
//	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	value := big.NewInt(1000000000000000000) // in wei (1 eth)
//	gasLimit := uint64(21000)                // in units
//	gasPrice, err := client.SuggestGasPrice(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
//	var data []byte
//	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
//	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = client.SendTransaction(context.Background(), signedTx)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex()) // tx sent: 0xec3ceb05642c61d33fa6c951b54080d1953ac8227be81e7b5e4e2cfed69eeb51
//
//	//模拟客户端的Commit方法手动开采区块。
//	client.Commit()
//
//	//现在您可以获取交易收据并看见其已被处理。
//	receipt, err := client.TransactionReceipt(context.Background(), signedTx.Hash())
//	if err != nil {
//		log.Fatal(err)
//	}
//	if receipt == nil {
//		log.Fatal("receipt is nil. Forgot to commit?")
//	}
//
//	fmt.Printf("status: %v\n", receipt.Status) // status: 1
//}
