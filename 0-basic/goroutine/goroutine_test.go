package goroutine

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func Test_Add(t *testing.T) {
	go say("world")
	say("hello")
}

func Test_multi_thread(t *testing.T) {

	for i := 0; i < 5; i++ {
		go AsyncFunc(i)
	}

}

func Test_Get_CPU_NUM(t *testing.T) {

	cpuNum := runtime.NumCPU() //获得当前设备的cpu核心数
	fmt.Println("cpu核心数:", cpuNum)
	runtime.GOMAXPROCS(cpuNum) //设置需要用到的cpu数量

}

func AsyncFunc(index int) {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += 1
	}
	fmt.Printf("线程%d, sum为:%d\n", index, sum)
}

/**
在执行一些昂贵的计算任务时，我们希望能够尽量利用现代服务器普遍具备的多核特性来尽量将任务并行化，从而达到降低总计算时间的目的。此时我们需要了解 CPU 核心的数量，并针对性地分解计算任务到多个 goroutine 中去并行运行。

下面我们来模拟一个完全可以并行的计算任务：计算 N 个整型数的总和。我们可以将所有整型数分成 M 份，M 即 CPU 的个数。让每个 CPU 开始计算分给它的那份计算任务，最后将每个 CPU 的计算结果再做一次累加，这样就可以得到所有 N 个整型数的总和：
*/

type Vector []float64

// 分配给每个CPU的计算任务
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {

		//v[i] += u.(v[i])
	}
	c <- 1 // 发信号告诉任务管理者我已经计算完成了
}

const NCPU = 16 // 假设总共有16核
func (v Vector) DoAll(u Vector) {
	c := make(chan int, NCPU) // 用于接收每个CPU的任务完成信号
	for i := 0; i < NCPU; i++ {
		go v.DoSome(i*len(v)/NCPU, (i+1)*len(v)/NCPU, u, c)
	}
	// 等待所有CPU的任务完成
	for i := 0; i < NCPU; i++ {
		<-c // 获取到一个数据，表示一个CPU计算完成了
	}
	// 到这里表示所有计算已经结束
}
