package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

//func Test_multy_core(t *testing.T) {
//	runtime.GOMAXPROCS(1)
//
//}

func sum(seq int, ch chan int) {
	defer close(ch)
	sum := 0
	for i := 1; i <= 100000000000; i++ {
		sum += i
	}
	fmt.Printf("子协程%d运算结果:%dn", seq, sum)
	fmt.Println()
	ch <- sum
}

func Test_multi_core(t *testing.T) {

	// 最大 CPU 核心数
	cpus := runtime.NumCPU()
	run_calculation(cpus)
}

func Test_multi_core_half(t *testing.T) {

	// 最大 CPU 核心数
	run_calculation(6)
}

func Test_multi_core_1(t *testing.T) {

	// 最大 CPU 核心数
	run_calculation(1)
}

func run_calculation(cpus int) {
	runtime.GOMAXPROCS(cpus)

	// 启动时间
	start := time.Now()
	chs := make([]chan int, 12)
	for i := 0; i < len(chs); i++ {
		chs[i] = make(chan int, 1)
		go sum(i, chs[i])
	}
	sum := 0
	for _, ch := range chs {
		res := <-ch
		sum += res
	} // 结束时间
	end := time.Now() // 打印耗时
	fmt.Printf("最终运算结果: %d, 执行耗时(s): %fn", sum, end.Sub(start).Seconds())

	fmt.Println("")
}
