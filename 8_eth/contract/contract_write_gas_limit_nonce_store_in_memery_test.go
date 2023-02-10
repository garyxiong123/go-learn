package contract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	store "github.com/garyxiong123/go-learn/eth/contract/contract_generate" // for demo
	"github.com/shopspring/decimal"
	"log"
	"math/big"
	"strings"
	"sync"
	"testing"
)

//gas price ++  BSC 5%   -10%

//gas limit ++ 不可以不加， 什么情况加

//nonce, err := cli.GetPendingNonce(authCli.Address.Hex())
//用 laster 还是 pending

//Returned error: transaction underpriced
//
//Returned error: replacement transaction underpriced

var testnet = "https://data-seed-prebsc-1-s3.binance.org:8545"
var privateKeyStr = "a848c1798bbb423e2ebfaa1a6f6ae9eae905393ee54dc6900a0ce73251843d4e"
var contractAddress = "0xA0cD2B032FeF21c3FEfD113bF00eEa490D429115"
var waitGroup = sync.WaitGroup{}
var mutex = sync.Mutex{}

var globalNonceUsed = int64(-1)

func Test_Deploy(t *testing.T) {
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

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	address, _, _, err := store.DeployStore(auth, client, "1.0.0.0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract address :" + (address.String())) // "bar"
}

//https://testnet.bscscan.com/tx/0xcf18cac21fd2cdc7c78de3e4e4b3466f9af19056db5429bb7d362a12b96f03b7
func Test_Pending_Nonce(t *testing.T) {
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func() {
			mutex.Lock()
			defer mutex.Unlock()
			asyncFunc(false)
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	fmt.Println("don")
}

func Test_Latest_Nonce(t *testing.T) {
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func() {
			mutex.Lock()
			defer mutex.Unlock()
			asyncFunc(true)
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	fmt.Println("don")
}

func asyncFunc(latest bool) {
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

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("before gasPrice:" + gasPrice.String())
	nonce := uint64(0)
	if latest {
		c, err := rpc.DialContext(context.Background(), testnet)
		if err != nil {
			log.Fatal(err)
		}
		var result hexutil.Uint64
		err = c.CallContext(context.Background(), &result, "eth_getTransactionCount", fromAddress, "latest")
		if err != nil {
			log.Fatal(err)
		}
		nonce = uint64(result)
	} else {
		nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal(err)
		}
	}

	//nonce, err := client.PendingNonceAt(context.Background(), fromAddress)

	//fmt.Println("latestNonce=", strconv.FormatInt(int64(latestNonce), 10), ",nonce="+strconv.FormatInt(int64(nonce), 10)) // "bar"
	//nonce = latestNonce
	if globalNonceUsed >= int64(nonce) {
		fmt.Println("globalNonceUsed=nonce") // "bar"
		nonce = uint64(globalNonceUsed + 1)
	}
	fmt.Println("nonce=", nonce) // "bar"

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce + 0))
	auth.Value = big.NewInt(0) // in wei
	//auth.GasLimit = uint64(1000000000) // in units
	//auth.GasPrice = big.NewInt(10000000000)
	//gasPrice = big.NewInt(11000000000)

	auth.GasPrice = decimal.NewFromBigInt(gasPrice, 0).Add(decimal.NewFromBigInt(gasPrice, 0).Mul(decimal.NewFromFloat(0))).BigInt()

	//fmt.Println("after gasPrice:" + auth.GasPrice.String())

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
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar000"))

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

	//fmt.Println("estimateGasLimit:" + strconv.FormatUint(estimateGasLimit, 10))

	auth.GasLimit = estimateGasLimit
	//auth.GasLimit = estimateGasLimit - estimateGasLimit/10 // in units

	tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		//if errors.Is(err, core.ErrGasLimit) {
		//	log.Println(err)
		//}
		//if errors.Is(err, core.ErrIntrinsicGas) {
		//	log.Println(err)
		//}
		//
		//if errors.Is(err, core.ErrGasUintOverflow) {
		//	log.Println(err)
		//}
		log.Println(err)
		return
	}
	globalNonceUsed = int64(nonce)

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

	//result, err := instance.Items(nil, key)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(string(result[:])) // "bar"
}
