package proc

import (
	"github.com/zeromicro/go-zero/core/proc"
	"testing"
)

func TestProc(t *testing.T) {

	proc.AddWrapUpListener(func() {
		println("rapUpListener")
	})
	proc.SetTimeToForceQuit(5)
	proc.AddShutdownListener(func() {
		println("gary shut down")
	})

	proc.Env("test")

	proc.StartProfile()
}
