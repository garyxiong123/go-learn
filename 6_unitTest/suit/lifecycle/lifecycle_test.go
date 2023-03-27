package lifecycle

// ssuite_test.go

import (
	"crypto/rand"
	"fmt"
	"github.com/stretchr/testify/suite"
	"math/big"
	"testing"
)

type _Suite struct {
	suite.Suite
	id *big.Int
}

func (s *_Suite) AfterTest(suiteName, testName string) {
	fmt.Printf("AfterTest: suiteName=%s, testName=%s\n", suiteName, testName)
}

func (s *_Suite) BeforeTest(suiteName, testName string) {
	s.id, _ = rand.Int(rand.Reader, big.NewInt(100))
	fmt.Printf("gary test=%d\n", s.id)

	fmt.Printf("BeforeTest: suiteName=%s, testName=%s\n", suiteName, testName)
}

func (s *_Suite) SetupSuite() {
	fmt.Printf("SetupSuite()...\n")
}

func (s *_Suite) TearDownSuite() {
	fmt.Printf("TearDownSuite()...\n")
}

func (s *_Suite) SetupTest() {
	fmt.Printf("SetupTest()...\n")
}

func (s *_Suite) TearDownTest() {
	fmt.Printf("TearDownTest()...\n")
}

func (s *_Suite) TestFoo() {
	foo()
}

func foo() {
	println("foo")
}

func (s *_Suite) TestGoo() {
	goo()
}

func goo() {
	println("goo")
}

// 让 go test 执行测试
func TestGooFoo(t *testing.T) {
	suite.Run(t, new(_Suite))
}
