package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_defer(t *testing.T) {

	defer fmt.Println("1")

	defer fmt.Println("2")

	defer fmt.Println("3")

	fmt.Println("print normal")

}

/**
**处理业务或逻辑中涉及成对的操作是一件比较烦琐的事情，比如打开和关闭文件、接收请求和回复请求、加锁和解锁等。在这些操作中，最容易忽略的就是在每个函数退出处正确地释放和关闭资源。
*
 */
func Test_usage(t *testing.T) {
	fileSize("11")
}

func fileSize(filename string) int64 {
	f, err := os.Open(filename)
	if err != nil {
		return 0
	}
	// 延迟调用Close, 此时Close不会被调用
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		// defer机制触发, 调用Close关闭文件
		return 0
	}
	size := info.Size()
	// defer机制触发, 调用Close关闭文件
	return size
}
