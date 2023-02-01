package unitTest

import (
	"crypto/sha256"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
	"unsafe"
)

//一种是 uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。
//byte 类型是 uint8 的别名，对于只占用 1 个字节的传统 ASCII 编码的字符来说，完全没有问题，例如 var ch byte = 'A'，字符使用单引号括起来。

func Test_Byte(t *testing.T) {
	//byte = uint8
	var ch byte = 65
	println(ch)

	var a string = "hello  world"

	var b []byte = []byte(a) // string转[]byte

	a = string(b) // []byte转string

	//Uint32ToBytes := common2.Uint32ToBytes(uint32(1))
	//println(Uint32ToBytes)
}

//int类型的大小为 8 字节
//int8类型大小为 1 字节
//int16类型大小为 2 字节
//int32类型大小为 4 字节
//int64类型大小为 8 字节
//有符号整数类型
func Test_int(t *testing.T) {
	var a int = 1
	var b int8 = 2
	var c int16 = 3
	var d int32 = 4
	var e int64 = 5
	//unsafe.Sizeof() 只返回数据类型的大小，不管引用数据的大小，单位为Byte
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(c))
	fmt.Println(unsafe.Sizeof(d))
	fmt.Println(unsafe.Sizeof(e))
}

func Test_uint_int(t *testing.T) {
	var ss int8

	ss = 2
	ss = 25
	ss = 127
	//ss = 128 error
	println(ss)

	var uint_v uint8
	uint_v = 255

	var ssss uint16

	var ssss8 uint8

	println(unsafe.Sizeof(ssss8))
	ssss = 256*256 - 1
	println(ssss)
	println(uint_v)

}

func Test_base_x(t *testing.T) {

	// 标准Base64编码
	src := "hello world"
	res16 := hex.EncodeToString([]byte(src))
	println(res16)

	//大写
	res32 := base32.StdEncoding.EncodeToString([]byte(src))
	println(res32)

	//大小写
	res64 := base64.StdEncoding.EncodeToString([]byte(src))

	println(res64)

}

//sha256 256b bit = 32 byte
func Test_sha256(t *testing.T) {
	src := "hello world"
	srcByte := []byte(src)
	sha256New := sha256.New()
	sha256Bytes := sha256New.Sum(srcByte)
	sha256String := hex.EncodeToString(sha256Bytes)
	println(len(sha256String))
	println(sha256String)

}

func Test_sha256_1(t *testing.T) {
	src := "hello world"
	m := sha256.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))

	println(res)
	//b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9   =》 64 【16】=》256/4 =64
	println(len(res))

}

func Test_base58(t *testing.T) {

	println(len("Qmf27YAnamhdqnUNtEPAARe3j6bwzTYWfFx7yAWp43Xh5V"))
	println(len("QmZTR5bcpQD7cFgTorqxZDYaew1Wqgfbd2ud9QqGPAkK2V"))
	println(len("QmRkBEucWA35gJN4tvDdWQTjkCGSwK8R3T2QNKSuX4DBm9"))
	println(len("f01701220c3c4733ec8affd06cf9e9ff50ffc6bcd2ec85a6170004bb709669c31de94391a"))

	//bafybeiccfclkdtucu6y4yc5cpr6y3yuinr67svmii46v5cfcrkp47ihehy
	//
	//bWqxBEKC3P8tqsKc98xmWNzrzDtRLMiMPL8wBuTGsMnR

	//cid.Cid{}

	//https://ipfs.io/ipfs/f01701220c3c4733ec8affd06cf9e9ff50ffc6bcd2ec85a6170004bb709669c31de94391a
	//                 32/3O1E0921GU4ECVCHBVT0R7PT7VL1VU6NJ9EP1D62S009ERGIPKS67F98E8Q

}
