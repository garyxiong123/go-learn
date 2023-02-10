package client

import (
	"context"
	"testing"
)

//若您没有现有以太坊客户端，您可以连接到infura网关。Infura管理着一批安全，可靠，可扩展的以太坊[geth和parity]节点
func Test_BSC_net(t *testing.T) {
	gasprice, err := BscClient.SuggestGasPrice(context.Background())

	if err != nil {

	}
	println(gasprice)

	networkId, err := BscClient.NetworkID(context.Background())
	chainId, err := BscClient.ChainID(context.Background())
	blockNumber, err := BscClient.BlockNumber(context.Background())
	BscClient.SuggestGasTipCap()
	println(networkId)

	println(chainId)
	println(blockNumber)
	if err != nil {

	}

}
