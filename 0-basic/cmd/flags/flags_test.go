package flags

import (
	"fmt"
	"github.com/urfave/cli"
	"runtime"
	"testing"
)

var (
	gitCommit = "unknown"
	gitDate   = "unknown"
	version   = "unknown"
)

func Test_Flags(t *testing.T) {

	cli.VersionPrinter = func(ctx *cli.Context) {
		fmt.Println("Version:", ctx.App.Version)
		fmt.Println("Git Commit:", gitCommit)
		fmt.Println("Git Commit Date:", gitDate)
		fmt.Println("Architecture:", runtime.GOARCH)
		fmt.Println("Go Version:", runtime.Version())
		fmt.Println("Operating System:", runtime.GOOS)
	}

}
