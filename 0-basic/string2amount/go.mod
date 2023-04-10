module github.com/garyxiong123/go-learn/basic1/amount

go 1.17

require github.com/bnb-chain/zkbnb-crypto v0.0.8-0.20230217071307-211b4b7e8923

require (
	github.com/consensys/gnark-crypto v0.7.0 // indirect
	github.com/ethereum/go-ethereum v1.10.26 // indirect
)

require (
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/sys v0.2.0 // indirect
)

replace github.com/consensys/gnark-crypto => github.com/bnb-chain/gnark-crypto v0.7.1-0.20230203031630-7c643ad11891

replace github.com/bnb-chain/zkbnb-crypto => github.com/15000785133/zkbnb-crypto v0.0.8-0.20230330200044-c6f6b08fc397

replace github.com/bnb-chain/zkbnb-eth-rpc => github.com/15000785133/zkbnb-eth-rpc v0.0.3-0.20230323082501-ae497b894451
