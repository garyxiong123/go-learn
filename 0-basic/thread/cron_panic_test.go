package thread

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
	"time"
)

func TestCron_Panic_without_recover(t *testing.T) {

	cronJob := cron.New(cron.WithChain(
		cron.SkipIfStillRunning(cron.DiscardLogger),
	))

	_, err := cronJob.AddFunc("@every 1s", func() {

		panic("error happen")
		println("sss") //不打印
	})
	if err != nil {
		logx.Errorf("set tx pending count failed:%s", err.Error())
	}

	if err != nil {
		panic(err)
	}
	cronJob.Start()
	select {}

	println("job started")
}

func TestCron_Panic_with_recover(t *testing.T) {

	cronJob := cron.New(cron.WithChain(
		cron.SkipIfStillRunning(cron.DiscardLogger),
	))

	_, err := cronJob.AddFunc("@every 1s", func() {
		defer func() {
			if x := recover(); x != nil {
				time.Sleep(5 * time.Second)
				fmt.Printf("[ERROR]: My panic handle error: %s\n", x)
			}
		}()

		panic("error happen")
		println("sss")
	})
	if err != nil {
		logx.Errorf("set tx pending count failed:%s", err.Error())
	}

	if err != nil {
		panic(err)
	}
	cronJob.Start()
	select {}

	println("job started")
}
