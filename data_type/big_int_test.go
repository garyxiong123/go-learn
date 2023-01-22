package unitTest

import (
	"fmt"
	"math/big"
	"testing"
)

//一种是 uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。
//byte 类型是 uint8 的别名，对于只占用 1 个字节的传统 ASCII 编码的字符来说，完全没有问题，例如 var ch byte = 'A'，字符使用单引号括起来。

func Test_BigInt(t *testing.T) {
	big1 := new(big.Int).SetUint64(uint64(123))
	fmt.Println("big1 is: ", big1)
	big1.Int64()

	big2 := big1.Uint64()
	fmt.Println("big2 is: ", big2)

}

func Test_BigInt_1(t *testing.T) {
	// Creating a new big.Int with the value of 2^64
	i := big.NewInt(1)
	i.Lsh(i, 64)

	// Multiplying the big.Int by 10
	i.Mul(i, big.NewInt(10))

	// Converting the big.Int to a string
	fmt.Println(i.String())

}

func Test_BigInt_2_Unit64(t *testing.T) {
	var smallnum, _ = new(big.Int).SetString("2188824200011112223222", 10)
	smallnum.MarshalJSON()
	var smallnumHex, _ = new(big.Int).SetString("2188824200011112223", 16)
	num := smallnum.Uint64()
	println(num)
	println(smallnumHex)
}

func Test_BigInt_veryBig(t *testing.T) {
	//print(int('af6b80f7c6b8d2e5ce1cfa3a58c7c8530a7f75bc4f73569a8dcffbde3efc0753', 16))
	//print(int('af6b80f7c6b8d2e5ce1cfa3a58c7c8530a7f75bc4f73569a8dcffbde3efc075', 16))
	//print(21888242871839275222246405745257275088548364400416034343698204186575808495617)

	aa, _ := new(big.Int).SetString("af6b80f7c6b8d2e5ce1cfa3a58c7c8530a7f75bc4f73569a8dcffbde3efc0753", 16)
	aa10, _ := new(big.Int).SetString("21888242871839275222246405745257275088548364400416034343698204186575808495617", 10)
	println(aa)
	println(aa10)
	big11 := new(big.Int).SetUint64(uint64(1000000000000000000))

	big11.Add(big11, big11)

	words := big11
	//big11.SetBits([]{})
	fmt.Println("big11 is: ", words)

}
