all: build

#ABI (Application Binary Interface) 是指合约与外部程序之间的接口。生成 ABI 可以使外部程序能够调用合约中的方法。
#
#在 Solidity 中，可以使用编译器来生成 ABI。
#
#安装 Solidity 编译器:

install-solc:
	npm install -g solc

#安装abigen 生成有问题，所以采用了copy的方式
install-abigen:
	go get github.com/ethereum/go-ethereum
	cd $GOPATH/src/github.com/ethereum/go-ethereum/
	make
	make devtools	make devtools


#2:使用编译器编译合约，并使用 --abi 选项输出 ABI:
generate-abi:
	solc --abi Store.sol -o ./


#3: 也可以使用 abigen 工具，它可以将 solidity 源代码编译成 go 语言的 ABI 文件

#--abi: Mandatory path to the contract ABI to bind to
#--pkg: Mandatory Go package name to place the Go code into
#--type: Optional Go type name to assign to the binding struct
#--out: Optional output path for the generated Go source file (not set = stdout)
generate-go-abi:
	abigen --abi Store.abi --pkg contract_generate --type Store --out Store.go


#4: Deploying contracts to Ethereum =>Bin
generate-bin:
	solc --bin Store.sol --overwrite -o ./



#5: DeployStorage function has been injected:
generate-bin-inject:
	abigen --abi Store.abi --pkg contract_generate --type Store --out Store.go --bin Store.bin











