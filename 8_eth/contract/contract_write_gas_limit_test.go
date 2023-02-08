package contract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	store "github.com/garyxiong123/go-learn/eth/contract/contract_generate"
	"log"
	"math/big"
	"testing"
)

/**
*
 */
func TestGasLimit(t *testing.T) {
	tx(10000000)
}

func TestGasDifForCalc(t *testing.T) {

}

//exceeds block gas limit
func TestExceedsBlockGasLimit(t *testing.T) {
	tx(100000000)
}

func TestExceedsBlockGas1Limit(t *testing.T) {
	tx(10000000)
}

//ether被最小单位为wei，1ether = 10^{\18}wei
//func checkTxFee(gasPrice *big.Int, gas uint64, cap float64) error {
//	// Short circuit if there is no cap for transaction fee at all.
//	if cap == 0 {
//		return nil
//	}
//	feeEth := new(big.Float).Quo(new(big.Float).SetInt(new(big.Int).Mul(gasPrice, new(big.Int).SetUint64(gas))), new(big.Float).SetInt(big.NewInt(params.Ether)))
//	feeFloat, _ := feeEth.Float64()
//	if feeFloat > cap {
//		return fmt.Errorf("tx fee (%.2f ether) exceeds the configured cap (%.2f ether)", feeFloat, cap)
//	}
//	return nil
//}
//exceeds the configured cap (1.00 ether)
func TestExceedsConfiguredCapGasLimit(t *testing.T) {
	tx(100000001)
}

//intrinsic gas too low
func TestGasTooLow(t *testing.T) {
	tx(21000)
}

func tx(gasLimit uint64) {
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
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = gasLimit   // in units
	auth.GasPrice = gasPrice
	fmt.Println("gas =", gasLimit*gasPrice.Uint64()) // "bar"

	address := common.HexToAddress(contractAddress)
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar"))

	tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

	result, err := instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result[:])) // "bar"
}
