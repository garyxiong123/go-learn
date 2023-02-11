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

func TestBlockGasLimit(t *testing.T) {
	//--------transfer bnb,commit   one block
	//0x380eb2a5c9efce75e0addb4b1a1267355efbd759a71a808c3ebeff1ee76cf420  1 commit block   transfer   1
	//https://dashboard.tenderly.co/tx/bsc-testnet/0x380eb2a5c9efce75e0addb4b1a1267355efbd759a71a808c3ebeff1ee76cf420/gas-usage
	//1+4+4+2+5+2+2+32=52    52+69=121
	//Total Gas - 86,680 Gas
	//Actual Gas Used - 86,680 Gas
	//Initial Gas - 30,560 Gas
	//CALL - 56,120 Gas
	//fmt.Println("transfer bnb,commit block,1 tx per block,1 block", 30560-(21000+56*16+136*4+9*192*4))
	fmt.Println("transfer bnb,commit block,1 tx per block,1 block", 30560-(21000+52*16+69*4+9*121*4))

	//0xb7d1b2b5380290bf6f85e92b5b3d002192ca67b066488ca5126c8e2af7efaf68  1  commit block   transfer  2
	//https://dashboard.tenderly.co/tx/bsc-testnet/0xb7d1b2b5380290bf6f85e92b5b3d002192ca67b066488ca5126c8e2af7efaf68/gas-usage
	//Total Gas - 87,172 Gas
	//Actual Gas Used - 87,172 Gas
	//Initial Gas - 31,052 Gas
	//CALL - 56,120 Gas
	//fmt.Println("transfer bnb,commit block,2 tx per block,1 block", 31052-(21000+2*56*16+2*136*4+8*192*4))
	fmt.Println("transfer bnb,commit block,2 tx per block,1 block", 31052-(21000+2*52*16+2*69*4+8*121*4))

	//0x02e389e0b52f5a3b3ec8bd9fcc78182307e1e8960c4cb12c05be318a498fc662  1  commit block   transfer  3
	//https://dashboard.tenderly.co/tx/bsc-testnet/0x02e389e0b52f5a3b3ec8bd9fcc78182307e1e8960c4cb12c05be318a498fc662/gas-usage
	//Total Gas - 87,676 Gas
	//Actual Gas Used - 87,676 Gas
	//Initial Gas - 31,556 Gas
	//CALL - 56,120 Gas
	fmt.Println("transfer bnb,commit block,3 tx per block,1 block", 31556-(21000+3*52*16+3*69*4+7*121*4))

	//-----transfer bnb,verify  one block
	//https://dashboard.tenderly.co/tx/bsc-testnet/0xd7f05674338178e2a29ea70a87641d1e822462f8870fc53e02d2bf2531b5beba/gas-usage
	//Total Gas - 271,475 Gas
	//Actual Gas Used - 271,475 Gas
	//Initial Gas - 28,260 Gas
	//CALL - 243,215 Gas
	//fmt.Println("transfer bnb,verify block,1 tx per block,1 block", 30560-(21000+52*16+69*4+9*121*4))

	//0x35005d1b80e01f081a37bc402b513699e236078ce4af8e43ddb90259c184e24b  1  commit block   transfer  2
	//https://dashboard.tenderly.co/tx/bsc-testnet/0x35005d1b80e01f081a37bc402b513699e236078ce4af8e43ddb90259c184e24b/gas-usage
	//Total Gas - 271,475 Gas
	//Actual Gas Used - 271,475 Gas
	//Initial Gas - 28,260 Gas
	//CALL - 243,215 Gas
	//fmt.Println("transfer bnb,verify block,2 tx per block,1 block", 31052-(21000+2*52*16+2*69*4+8*121*4))

	//0x398ccb2ff9cd4c7d0435ab22e269137f8ddad6909e52a9d0475e34b009c21da2  1  commit block   transfer  3
	//https://dashboard.tenderly.co/tx/bsc-testnet/0x398ccb2ff9cd4c7d0435ab22e269137f8ddad6909e52a9d0475e34b009c21da2/gas-usage
	//Total Gas - 271,463 Gas
	//Actual Gas Used - 271,463 Gas
	//Initial Gas - 28,248 Gas
	//CALL - 243,215 Gas
	//fmt.Println("transfer bnb,verify block,3 tx per block,1 block", 31556-(21000+3*52*16+3*69*4+7*121*4))

	//-----withdraw bnb  commit  one block
	//0x1c6d99d4cc3c096c522e4fd9a361c1dd2f8187cee61564e925e4ae7d854ed6c3  1 commit block   withdraw   1
	//https://dashboard.tenderly.co/tx/bsc-testnet/0x1c6d99d4cc3c096c522e4fd9a361c1dd2f8187cee61564e925e4ae7d854ed6c3/gas-usage
	//1+4+20+2+16+2+2=47    47+74=121
	//Total Gas - 88,266 Gas
	//Actual Gas Used - 88,266 Gas
	//Initial Gas - 30,520 Gas
	//CALL - 57,746 Gas
	fmt.Println("withdraw bnb,commit block,1 tx per block,1 block", 30520-(21000+47*16+74*4+9*121*4))

	//0x299c6c214ca46aff2ccc8c62af90b2670d5df3c2c2a8dd545d13db3c9aef31e7  1 commit block   withdraw   2
	//https://dashboard.tenderly.co/tx/bsc-testnet/0x299c6c214ca46aff2ccc8c62af90b2670d5df3c2c2a8dd545d13db3c9aef31e7/gas-usage
	//Total Gas - 90,392 Gas
	//Actual Gas Used - 90,392 Gas
	//Initial Gas - 31,020 Gas
	//CALL - 59,372 Gas
	fmt.Println("withdraw bnb,commit block,2 tx per block,1 block", 31020-(21000+2*47*16+2*74*4+8*121*4))

	//0x960dbd920733d5a9c3cacc278872c8d58d55957593393e661fe6a7e0a676759e  1 commit block   withdraw   4
	//https://dashboard.tenderly.co/tx/bsc-testnet/0x960dbd920733d5a9c3cacc278872c8d58d55957593393e661fe6a7e0a676759e/gas-usage
	//Total Gas - 94,634 Gas
	//Actual Gas Used - 94,634 Gas
	//Initial Gas - 32,008 Gas
	//CALL - 62,626 Gas
	fmt.Println("withdraw bnb,commit block,4 tx per block,1 block", 32008-(21000+4*47*16+4*74*4+6*121*4))

	//-----withdraw bnb  commit  1 tx per block,2 block
	//0xcec7153186fe98d50aeebac918c2143900961ba8d18465bf823991982af7863e  1 commit block   withdraw   1
	//https://dashboard.tenderly.co/tx/bsc-testnet/0xcec7153186fe98d50aeebac918c2143900961ba8d18465bf823991982af7863e/gas-usage
	//Total Gas - 130,866 Gas
	//Actual Gas Used - 130,866 Gas
	//Initial Gas - 37,600 Gas
	//CALL - 93,266 Gas
	fmt.Println("withdraw bnb,commit block,1 tx per block,2 block", 37600-(21000+2*47*16+2*74*4+8*121*4))

	//-----withdraw bnb  commit  2 tx per block,2 block
	//0x36f8d963846fc9cd39ee380b630dfe097d1965c6ca71ef871cb87983c98299e7  commit  2 tx per block,2 block
	//https://dashboard.tenderly.co/tx/bsc-testnet/0x36f8d963846fc9cd39ee380b630dfe097d1965c6ca71ef871cb87983c98299e7/gas-usage
	//Total Gas - 135,078 Gas
	//Actual Gas Used - 135,078 Gas
	//Initial Gas - 38,552 Gas
	//CALL - 96,526 Gas
	fmt.Println("withdraw bnb,commit block,2 tx per block,2 block", 38552-(21000+4*47*16+4*74*4+6*121*4))

	//-----withdraw bnb  verify
	//0x010c1b3bd391f943993065a204e22cdcdd61efee7f7b92f1dbf2ba3771b51387  1 commit block   withdraw   1
	//https://dashboard.tenderly.co/tx/bsc-testnet/0x010c1b3bd391f943993065a204e22cdcdd61efee7f7b92f1dbf2ba3771b51387/gas-usage
	//Total Gas - 285,151 Gas
	//Actual Gas Used - 285,151 Gas
	//Initial Gas - 29,412 Gas
	//CALL - 255,739 Gas
	//fmt.Println("withdraw bnb,verify block,1 tx per block,1 block", 30520-(21000+47*16+74*4+9*121*4))

	//0x2a131a0ba41d84bb4a8ad96e5743b9716beb328e1737ea0235d844e7611cffa3  1 commit block   withdraw   2
	//https://dashboard.tenderly.co/tx/bsc-testnet/0x2a131a0ba41d84bb4a8ad96e5743b9716beb328e1737ea0235d844e7611cffa3/gas-usage
	//Total Gas - 298,791 Gas
	//Actual Gas Used - 298,791 Gas
	//Initial Gas - 30,528 Gas
	//CALL - 268,263 Gas
	//fmt.Println("withdraw bnb,verify block,2 tx per block,1 block", 31020-(21000+2*47*16+2*74*4+8*121*4))

	//0x2e293a3593d1b8667768303dc7109e9555d4af308b5e8405605e57182495cefb  1 commit block   withdraw   4
	//https://dashboard.tenderly.co/tx/bsc-testnet/0x2e293a3593d1b8667768303dc7109e9555d4af308b5e8405605e57182495cefb/gas-usage
	//Total Gas - 326,185 Gas
	//Actual Gas Used - 326,185 Gas
	//Initial Gas - 32,868 Gas
	//CALL - 293,317 Gas
	//fmt.Println("withdraw bnb,verify block,4 tx per block,1 block", 32008-(21000+4*47*16+4*74*4+6*121*4))

	//-----withdraw bnb  verify  1 tx per block,2 block
	//0x51f009311a3dcf73ee1699b02b701aeb735ec8c46c033aaec81e074dd3345ab0
	//https://dashboard.tenderly.co/tx/bsc-testnet/0x51f009311a3dcf73ee1699b02b701aeb735ec8c46c033aaec81e074dd3345ab0/gas-usage
	//Total Gas - 400,684 Gas
	//Actual Gas Used - 400,684 Gas
	//Initial Gas - 37,212 Gas
	//CALL - 363,472 Gas
	//fmt.Println("withdraw bnb,verify block,4 tx per block,1 block", 32008-(21000+4*47*16+4*74*4+6*121*4))

	//-----withdraw bnb  verify  2 tx per block,2 block
	//0xd18a10e484d1dfe888662dde75da5a38b1a39ea9a6979aaada3b1ca2546aaeab
	//https://dashboard.tenderly.co/tx/bsc-testnet/0xd18a10e484d1dfe888662dde75da5a38b1a39ea9a6979aaada3b1ca2546aaeab/gas-usage
	//Total Gas - 428,010 Gas
	//Actual Gas Used - 428,010 Gas
	//Initial Gas - 39,480 Gas
	//CALL - 388,530 Gas
	//fmt.Println("withdraw bnb,verify block,4 tx per block,1 block", 32008-(21000+4*47*16+4*74*4+6*121*4))

}
