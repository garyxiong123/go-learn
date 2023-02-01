package main

import (
	lru "github.com/hashicorp/golang-lru"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"time"
)

// this is thread safe for get and update and add and remove

func main() {
	onEvicted := func(k interface{}, v interface{}) {
		logx.Infof("onEvicted key=%s,value=%s", strconv.FormatInt(int64(k.(int)), 10), strconv.FormatInt(int64(v.(int)), 10))
	}

	l, err := lru.NewWithEvict(10, onEvicted)
	if err != nil {
		logx.Severef("err: %v", err)
	}

	for i := 1; i <= 10; i++ {
		go AsyncFunc(l, i, i)
	}
	time.Sleep(10 * time.Second)
	for _, key := range l.Keys() {
		value, _ := l.Get(key)
		logx.Infof("key=%s,value=%s", strconv.FormatInt(int64(key.(int)), 10), strconv.FormatInt(int64(value.(int)), 10))
	}
}

func AsyncFunc(c *lru.Cache, key int, value int) {
	c.Add(key, value)
}
