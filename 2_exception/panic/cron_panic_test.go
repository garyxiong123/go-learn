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
