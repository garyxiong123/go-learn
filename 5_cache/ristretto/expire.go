package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dgraph-io/ristretto"
)

func main() {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters:        100,
		MaxCost:            10,
		BufferItems:        64,
		IgnoreInternalCost: true,
		// Called when setting cost to 0 in `Set/SetWithTTL`
		Cost: func(value interface{}) int64 {
			return 1
		},
		OnEvict: func(item *ristretto.Item) {
			fmt.Printf("OnEvict key=%d \n", item.Key)
		},
	})
	if err != nil {
		panic(err)
	}

	// set item with 1s ttl
	cache.SetWithTTL("foo", "bar", 1, time.Second)

	// wait for value to pass through buffers
	cache.Wait()

	if val, ok := cache.Get("foo"); !ok {
		log.Printf("cache missing")
	} else {
		log.Printf("got foo: %v", val)
	}

	// sleep longer and try again
	time.Sleep(2 * time.Second)
	if val, ok := cache.Get("foo"); !ok {
		log.Printf("cache missing")
	} else {
		log.Printf("got foo: %v", val)
	}
}
