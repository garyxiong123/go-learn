package thread

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
)

func TestCron(t *testing.T) {
	cronJob := cron.New(cron.WithChain(
		cron.SkipIfStillRunning(cron.DiscardLogger),
	))
	_, err := cronJob.AddFunc("@every 1s", func() {
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

func TestCron_Single(t *testing.T) {
	i := 0
	c := cron.New()
	spec := "5/* * * * * ?"
	c.AddFunc(spec, func() {
		i++
		fmt.Println("cron running:", i)
	})
	c.Start()

	select {}
}

type TestJob struct {
}

func (this TestJob) Run() {
	panic("error happen")
	fmt.Println("testJob1...")
}

type Test2Job struct {
}

func (this Test2Job) Run() {
	fmt.Println("testJob2...")
}

//多个定时cron任务
func TestCron_Multi(t *testing.T) {
	i := 0
	c := cron.New()

	//AddFunc
	spec := "0/5 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		fmt.Println("cron running:", i)
	})

	//AddJob方法
	c.AddJob("@every 1s", TestJob{})
	c.AddJob(spec, Test2Job{})

	//启动计划任务
	c.Start()

	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()

	select {}
}
