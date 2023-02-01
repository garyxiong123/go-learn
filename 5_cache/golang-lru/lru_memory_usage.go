package main

import (
	"encoding/json"
	"fmt"
	lru "github.com/hashicorp/golang-lru"
	"github.com/zeromicro/go-zero/core/logx"
	"math/big"
	"runtime"
	"strconv"
)

func main() {

	memConsumed := func() uint64 {
		//runtime.GC() //GC，排除对象影响
		var memStat runtime.MemStats
		runtime.ReadMemStats(&memStat)
		return memStat.Sys
	}

	before := memConsumed() //获取创建goroutine前内存

	// 1000000     7666.950 M
	// 100000       798.119 M
	// 10000         82.471 M
	// 1000           9.875 M
	reloadAccounts(1000000)

	after := memConsumed() //获取创建goroutine后内存

	fmt.Printf("%.3f M\n", float64(after-before)/1024/1024)
}

type AccountAsset struct {
	AssetId                  int64
	Balance                  *big.Int
	OfferCanceledOrFinalized *big.Int
}

type AccountInfo struct {
	AccountId       uint
	AccountIndex    int64
	AccountName     string
	PublicKey       string
	AccountNameHash string
	L1Address       string
	Nonce           int64
	CollectionNonce int64
	AssetInfo       map[int64]*AccountAsset // key: index, value: balance
	AssetRoot       string
	Status          int
}

func toFormatAccountInfo() (formatAccountInfo *AccountInfo) {
	var assetInfo map[int64]*AccountAsset
	str := "{\"0\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"1\":{\"AssetId\":1,\"Balance\":100000000000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"2\":{\"AssetId\":2,\"Balance\":100000000010000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"3\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"4\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"5\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"6\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"7\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"8\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"9\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"10\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"11\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"12\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"13\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"14\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"15\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"16\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"17\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"18\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"19\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"20\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"21\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"22\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"23\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"24\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"25\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"26\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"27\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"28\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"29\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"30\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"31\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"32\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"33\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"34\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"35\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"36\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"37\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"38\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}," +
		"\"39\":{\"AssetId\":0,\"Balance\":99988000000000000,\"OfferCanceledOrFinalized\":0}}"
	_ = json.Unmarshal([]byte(str), &assetInfo)

	formatAccountInfo = &AccountInfo{
		AccountId:       100000000,
		AccountIndex:    100000000,
		AccountName:     "zkbnbqa02.zkbnb",
		PublicKey:       "7d344522973c6790f37a5cf5d31c13e2bc105b1903698eb138efac13c990cc25",
		AccountNameHash: "00bcfda5afb9bcbee48bf52b8911d0ad52089d5fd4aa4ffe838c18157c4d3ad1",
		L1Address:       "0x94F037c94731e82EDD79DFB03CaB14fBF798100f",
		Nonce:           100000,
		CollectionNonce: 100000,
		AssetInfo:       assetInfo,
		AssetRoot:       "301f69702abaa96ff347924f6d7fa49e1ef860211b371950565090a539de975b",
		Status:          1,
	}
	return formatAccountInfo
}

func reloadAccounts(count int) {
	onEvicted := func(k interface{}, v interface{}) {
		logx.Infof("onEvicted key=%s,value=%s", strconv.FormatInt(int64(k.(int)), 10), strconv.FormatInt(int64(v.(int)), 10))
	}

	l, err := lru.NewWithEvict(count, onEvicted)
	if err != nil {
		logx.Severef("err: %v", err)
	}
	for i := 0; i < count; i++ {
		accountInfo := toFormatAccountInfo()
		l.Add(i, accountInfo)
	}

	logx.Infof("len(l.Keys())=%s", strconv.FormatInt(int64(len(l.Keys())), 10))

}
