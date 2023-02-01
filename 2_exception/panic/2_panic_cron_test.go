package panic

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"testing"
	"time"
)

func newCronJob() *cron.Cron {
	cronJob := cron.New(cron.WithChain(
		cron.SkipIfStillRunning(cron.DiscardLogger),
	))
	return cronJob
}

//协程会终止 因此该过程一直在调用栈中重复发生：函数停止执行，
//调用延迟执行函数等。如果一路在延迟函数中没有recover函数的调用，则会到达该携程的起点，该携程结束，
//然后终止其他所有携程，包括主携程（类似于C语言中的主线程，该携程ID为1）

func TestCron_Panic_without_recover(t *testing.T) {

	cronJob := newCronJob()

	cronJob.AddFunc("@every 1s", func() {

		panic("error happen")

	})

	cronJob.Start()

	select {}

}

func TestCron_Panic_with_recover(t *testing.T) {

	cronJob := newCronJob()

	cronJob.AddFunc("@every 1s", func() {
		defer func() {
			if x := recover(); x != nil {
				time.Sleep(5 * time.Second)
				fmt.Printf("[ERROR]: My panic handle error: %s\n", x)
			}
		}()

		panic("error happen")

	})

	cronJob.Start()
	select {}

}
