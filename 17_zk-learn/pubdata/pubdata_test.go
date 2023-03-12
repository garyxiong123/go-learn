package pubdata

import (
	"bytes"
	"github.com/bnb-chain/zkbnb-crypto/ffmath"
	common2 "github.com/bnb-chain/zkbnb/common"
	"github.com/bnb-chain/zkbnb/types"
	"testing"
)

func Test_Generate_Deposit(t *testing.T) {
	ffmath.Add(
		nil,
		nil,
	)
	txInfo, err := types.ParseDepositTxInfo("{\"TxType\":2,\"AccountNameHash\":\"JjLJ5kxx/cVYqxUUFbbu9ZO+nLO1LMVHxwYiRc6adLI=\",\"AssetId\":65533,\"AssetAmount\":10203000000000000,\"AccountIndex\":0}")
	println(err)
	println(txInfo)
	//executor := &executor.DepositExecutor{
	//	BaseExecutor: executor.BaseExecutor{},
	//}
	//executor.SetTxInfo(txInfo)
	//
	//executor.GeneratePubData()
	var buf bytes.Buffer
	buf.WriteByte(uint8(types.TxTypeDeposit))
	buf.Write(common2.Uint32ToBytes(uint32(txInfo.AccountIndex)))
	buf.Write(common2.Uint16ToBytes(uint16(txInfo.AssetId)))
	buf.Write(common2.Uint128ToBytes(txInfo.AssetAmount))
	chunk1 := common2.SuffixPaddingBufToChunkSize(buf.Bytes())
	buf.Reset()
	buf.Write(chunk1)
	buf.Write(common2.PrefixPaddingBufToChunkSize(txInfo.AccountNameHash))
	buf.Write(common2.PrefixPaddingBufToChunkSize([]byte{}))
	buf.Write(common2.PrefixPaddingBufToChunkSize([]byte{}))
	buf.Write(common2.PrefixPaddingBufToChunkSize([]byte{}))
	buf.Write(common2.PrefixPaddingBufToChunkSize([]byte{}))
	pubData := buf.Bytes()
	println(pubData)

}
