package contract

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	store "github.com/garyxiong123/go-learn/eth/contract/contract_generate"
)

func Test_write_estimate_gas(t *testing.T) {

	//先查到nonce和燃气价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")

	//我们创建的智能合约有一个名为SetItem的外部方法，它接受solidity“bytes32”格式的两个参数（key，value）。
	//这意味着Go合约包要求我们传递一个长度为32个字节的字节数组。 调用SetItem方法需要我们传递我们之前创建的auth对象（keyed transactor）。
	//在幕后，此方法将使用它的参数对此函数调用进行编码，将其设置为事务的data属性，并使用私钥对其进行签名。 结果将是一个已签名的事务对象

	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar"))

	parsed, err := abi.JSON(strings.NewReader(store.StoreABI))
	if err != nil {
		log.Fatal(err)
	}

	encodedData, err := parsed.Pack("setItem", key, value)
	if err != nil {
		log.Fatal(err)
	}

	estimatedGas, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:     fromAddress,
		To:       &address,
		Data:     encodedData,
		GasPrice: gasPrice,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(estimatedGas)
}
