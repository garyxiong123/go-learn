package flags

import (
	"github.com/urfave/cli"
)

var (
	ConfigFlag = &cli.StringFlag{
		Name:  "config",
		Usage: "the config file",
	}
	ContractAddrFlag = &cli.StringFlag{
		Name:  "contractAddr",
		Usage: "the contract addresses file",
	}
)
