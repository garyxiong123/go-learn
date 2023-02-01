package main

import (
	lru "github.com/hashicorp/golang-lru"
	"github.com/zeromicro/go-zero/core/logx"
)

// this is thread safe for get and update and add and remove

func main() {
	onEvicted := func(k interface{}, v interface{}) {
		logx.Infof("onEvicted key=%s,value=%s", string(k.(int)), string(v.(int)))
	}

	l, err := lru.NewWithEvict(5, onEvicted)
	if err != nil {
		logx.Severef("err: %v", err)
	}
	l.Add(1, 1)
	l.Add(2, 2)
	l.Add(3, 3)
	l.Add(4, 4)
	l.Get(1)
	l.Add(5, 5)
	l.Add(6, 6)

	for _, key := range l.Keys() {
		value, _ := l.Get(key)
		logx.Infof("key=%s,value=%s", string(key.(int)), string(value.(int)))
	}
}
