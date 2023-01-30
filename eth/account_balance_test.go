package eth

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Test_Account_Balance(t *testing.T) {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")

	//区块号设置为nil将返回最新的余额。
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance) // 25893860171173005034

	blockNumber := big.NewInt(5532993)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balanceAt) // 25729324269165216042

	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041

	//待处理的余额
	//有时您想知道待处理的账户余额是多少，例如，在提交或等待交易确认后。客户端提供了类似BalanceAt的方法，名为PendingBalanceAt，它接收账户地址作为参数。
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance) // 25729324269165216042
}
