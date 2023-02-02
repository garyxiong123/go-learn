package main

import (
	"fmt"
	lru "github.com/hashicorp/golang-lru"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"runtime"
	"strconv"
	"time"
)

func main() {

	memConsumed := func() uint64 {
		//runtime.GC() //GC，排除对象影响
		var memStat runtime.MemStats
		runtime.ReadMemStats(&memStat)
		return memStat.Sys
	}

	before := memConsumed() //获取创建goroutine前内存

	// 10000000   3387.616 M
	// 1000000     362.268 M
	// 100000       38.187 M
	// 10000         9.312 M
	// 1000           0 M
	reloadNfts(10000000)

	after := memConsumed() //获取创建goroutine后内存

	fmt.Printf("%.3f M\n", float64(after-before)/1024/1024)
}

type L2Nft struct {
	gorm.Model
	NftIndex            int64 `gorm:"uniqueIndex"`
	CreatorAccountIndex int64
	OwnerAccountIndex   int64  `gorm:"index:idx_owner_account_index"`
	NftContentHash      string `gorm:"index:idx_owner_account_index"`
	CreatorTreasuryRate int64
	CollectionId        int64
	L2BlockHeight       int64 `gorm:"index:idx_nft_index"`
}

func reloadNfts(count int) {
	onEvicted := func(k interface{}, v interface{}) {
		logx.Infof("onEvicted key=%s,value=%s", strconv.FormatInt(int64(k.(int)), 10), strconv.FormatInt(int64(v.(int)), 10))
	}

	l, err := lru.NewWithEvict(count, onEvicted)
	if err != nil {
		logx.Severef("err: %v", err)
	}
	for i := 0; i < count; i++ {
		l2Nft := &L2Nft{
			Model:               gorm.Model{ID: 1000, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			NftIndex:            100000000,
			CreatorAccountIndex: 100000000,
			OwnerAccountIndex:   100000000,
			NftContentHash:      "08f930cb5011aba0ce6a1ffc58e389a5269739c4c5946ede312c9b0ac9155aa7",
			CreatorTreasuryRate: 1,
			CollectionId:        100000,
			L2BlockHeight:       100000,
		}
		l.Add(i, l2Nft)
	}

	logx.Infof("len(l.Keys())=%s", strconv.FormatInt(int64(len(l.Keys())), 10))

}
