package main

import (
	"fmt"
	"github.com/dgraph-io/ristretto"
	"unsafe"
)

func main() {
	cache, err := ristretto.NewCache(&ristretto.Config{
		// num of keys to track frequency, usually 10*MaxCost
		NumCounters: 10 * 16 * 10,
		// cache size(max num of items)
		MaxCost: 10 * 16,
		// number of keys per Get buffer
		BufferItems: 64,
		// !important: always set true if not limiting memory
		IgnoreInternalCost: true,
		// Called when setting cost to 0 in `Set/SetWithTTL`
		Cost: func(value interface{}) int64 {
			fmt.Printf("Sizeof key=%d \n", unsafe.Sizeof(value))
			return int64(unsafe.Sizeof(value))
		},
		OnEvict: func(item *ristretto.Item) {
			fmt.Printf("OnEvict key=%d \n", item.Key)
		},
	})
	if err != nil {
		panic(err)
	}

	// put 20(>10) items to cache
	for i := 0; i < 20; i++ {
		cache.Set(i, i, 0)
	}

	// wait for value to pass through buffers
	cache.Wait()

	cntCacheMiss := 0
	for i := 0; i < 20; i++ {
		if _, ok := cache.Get(i); !ok {
			cntCacheMiss++
		} else {
			fmt.Printf("key=%d \n", i)
		}
	}
	fmt.Printf("%d of 20 items missed\n", cntCacheMiss)
}
