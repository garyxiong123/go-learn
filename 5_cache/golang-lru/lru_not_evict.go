package main

import (
	lru "github.com/hashicorp/golang-lru"
	"github.com/zeromicro/go-zero/core/logx"
)

// test that Add returns true/false if an eviction occurred

func main() {
	onEvicted := func(k interface{}, v interface{}) {
		logx.Infof("onEvicted key=%s,value=%s", string(k.(int)), string(v.(int)))
	}

	l, err := lru.NewWithEvict(10, onEvicted)
	if err != nil {
		logx.Severef("err: %v", err)
	}
	l.Add(1, 1)
	l.Add(2, 2)
	for _, key := range l.Keys() {
		value, _ := l.Get(key)
		logx.Infof("key=%s,value=%s", string(key.(int)), string(value.(int)))
	}
}
