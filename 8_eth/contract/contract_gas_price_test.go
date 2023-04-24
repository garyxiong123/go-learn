package contract

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"testing"
)

func Test_gas_price(t *testing.T) {
	client, err := ethclient.Dial("https://bsc-dataseed1.binance.org")
	//先查到nonce和燃气价格
	gasPrice, err := client.SuggestGasPrice(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(gasPrice)
}
