package contract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/shopspring/decimal"
	"log"
	"math/big"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	store "github.com/garyxiong123/go-learn/eth/contract/contract_generate" // for demo
)

//gas price ++  BSC 5%   -10%

//gas limit ++ 不可以不加， 什么情况加

//nonce, err := cli.GetPendingNonce(authCli.Address.Hex())
//用 laster 还是 pending

//Returned error: transaction underpriced
//
//Returned error: replacement transaction underpriced

func Test_Increase_10Percent_GasPrice(t *testing.T) {
	gasPriceBefore := getGasPrice()
	nonce := getNonce()
	fmt.Println("set foo default value start--------")
	doTransaction(nonce, gasPriceBefore, "foo", "foo")
	nonce = getNonce()

	gasPriceAfter := decimal.NewFromBigInt(gasPriceBefore, 0).Add(decimal.NewFromBigInt(gasPriceBefore, 0).Mul(decimal.NewFromFloat(0.1))).BigInt()
	fmt.Println("AveragePrice gasPrice:" + gasPriceBefore.String() + " after gasPrice" + gasPriceAfter.String())

	fmt.Println("AveragePrice gasPrice start--------")

	doTransaction(nonce+1, gasPriceBefore, "foo", "gasPriceBeforeValue")

	fmt.Println("after gasPrice start--------")

	doTransaction(nonce+1, gasPriceAfter, "foo", "gasPriceAfterValue")

	fmt.Println("set init default value start--------")

	doTransaction(nonce, gasPriceBefore, "init", "initValue")

	fmt.Println("get foo value start--------")

	getItem("foo")
}

//replacement transaction underpriced
func Test_Increase_9Percent_GasPrice(t *testing.T) {
	gasPriceBefore := getGasPrice()
	nonce := getNonce()
	fmt.Println("set foo default value start--------")
	doTransaction(nonce, gasPriceBefore, "foo", "foo")
	nonce = getNonce()

	gasPriceAfter := decimal.NewFromBigInt(gasPriceBefore, 0).Add(decimal.NewFromBigInt(gasPriceBefore, 0).Mul(decimal.NewFromFloat(0.09))).BigInt()
	fmt.Println("AveragePrice gasPrice:" + gasPriceBefore.String() + " after gasPrice" + gasPriceAfter.String())

	fmt.Println("AveragePrice gasPrice start--------")

	doTransaction(nonce+1, gasPriceBefore, "foo", "gasPriceBeforeValue")

	fmt.Println("after gasPrice start--------")

	doTransaction(nonce+1, gasPriceAfter, "foo", "gasPriceAfterValue")

	fmt.Println("set init default value start--------")

	doTransaction(nonce, gasPriceBefore, "init", "initValue")

	fmt.Println("get foo value start--------")

	getItem("foo")
}

//低于平均价格  transaction underpriced
func TestUnderPrice_Low_Than_AveragePrice(t *testing.T) {
	gasPrice := getGasPrice()
	nonce := getNonce()
	fmt.Println("set foo default value start--------")
	doTransaction(nonce, big.NewInt(gasPrice.Int64()-1), "foo", "foo")
}

//低于上次价格  transaction underpriced
func TestUnderPrice_Low_Than_LastTime(t *testing.T) {
	gasPriceBefore := getGasPrice()
	nonce := getNonce()
	fmt.Println("set foo default value start--------")
	doTransaction(nonce, gasPriceBefore, "foo", "foo")
	nonce = getNonce()

	gasPriceAfter := decimal.NewFromBigInt(gasPriceBefore, 0).Sub(decimal.NewFromBigInt(gasPriceBefore, 0).Mul(decimal.NewFromFloat(0.1))).BigInt()
	fmt.Println("AveragePrice gasPrice:" + gasPriceBefore.String() + " after gasPrice" + gasPriceAfter.String())

	fmt.Println("AveragePrice gasPrice start--------")

	doTransaction(nonce+1, gasPriceBefore, "foo", "gasPriceBeforeValue")

	fmt.Println("after gasPrice start--------")

	doTransaction(nonce+1, gasPriceAfter, "foo", "gasPriceAfterValue")

	fmt.Println("set init default value start--------")

	doTransaction(nonce, gasPriceBefore, "init", "initValue")

	fmt.Println("get foo value start--------")

	getItem("foo")

}

func doTransaction(nonce uint64, gasPrice *big.Int, keyStr string, valuesStr string) {
	client, err := ethclient.Dial(testnet)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei

	auth.GasPrice = decimal.NewFromBigInt(gasPrice, 0).Add(decimal.NewFromBigInt(gasPrice, 0).Mul(decimal.NewFromFloat(0))).BigInt()

	parsed, err := abi.JSON(strings.NewReader(store.StoreABI))
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(contractAddress)
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte(keyStr))
	copy(value[:], []byte(valuesStr))

	encodedData, err := parsed.Pack("setItem", key, value)
	if err != nil {
		log.Fatal(err)
	}

	estimateGasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:     fromAddress,
		To:       &address,
		Data:     encodedData,
		GasPrice: gasPrice,
	})
	if err != nil {
		log.Fatal(err)
	}

	auth.GasLimit = estimateGasLimit

	tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	time.Sleep(30 * time.Second)

	transaction, isPending, err := client.TransactionByHash(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("transaction hash="+" hash="+transaction.Hash().Hex()+" Nonce="+strconv.FormatUint(transaction.Nonce(), 10)+" GasPrice= "+transaction.GasPrice().String()+" isPending=", isPending) // "bar"

	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err == nil {
		fmt.Println("receipt hash=" + " hash=" + receipt.TxHash.Hex() + " Status=" + strconv.FormatUint(receipt.Status, 10)) // "bar"
	} else {
		log.Print(err)
	}

	result, err := instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("key=", keyStr, ",value=", string(result[:])) // "bar"
}

func getNonce() uint64 {
	client, err := ethclient.Dial(testnet)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
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
	return nonce
}

func getGasPrice() *big.Int {
	client, err := ethclient.Dial(testnet)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return gasPrice
}

func getItem(keyStr string) {
	client, err := ethclient.Dial(testnet)
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress(contractAddress)
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}
	key := [32]byte{}
	copy(key[:], []byte(keyStr))
	result, err := instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(result[:])) // "bar"
}
