package eth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/garyxiong123/go-learn/eth_learn_eth/erc721"
	"log"
	"math/big"
	"testing"
)

func Test_Deploy_Erc721(t *testing.T) {
	transactOpts, err := BuildDefaultTransactOpts()
	if err != nil {
	}
	client, err := ethclient.Dial("https://bsc-testnet.nodereal.io/v1/a1cee760ac744f449416a711f20d99dd")

	addr, tx, erc721, err := erc721.DeployErc721(transactOpts, client, "nft-gary", "nft-gary")

	//transfer from to
	transactOpts1, err := BuildDefaultTransactOpts()
	tx, err = erc721.Mint(transactOpts1, common.HexToAddress("0x4909d4D440E8ffF61738E8Cb7b2b0a4aaFF7b896"), big.NewInt(2))
	if err != nil {

	}

	//erc721.Approve()
	log.Print(tx)
	log.Print(addr)
	log.Print(tx)
	log.Print(erc721)
	log.Print(err)

}
