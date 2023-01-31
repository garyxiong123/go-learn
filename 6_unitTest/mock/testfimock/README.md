
go get github.com/stretchr/testify/mock


2:在需要 mock 的接口上定义一个 mock 结构体，并继承 testify/mock.Mock


type MyMock struct {
mock.Mock
}


3:定义接口需要实现的方法


func (m *MyMock) DoSomething(arg1 int, arg2 string) (int, error) {
// method body
}


4:在测试函数中使用 mock 对象

func TestMyFunction(t *testing.T) {
myMock := new(MyMock)

    // set expectations on the mock
    myMock.On("DoSomething", 1, "arg").Return(0, nil)

    // call the function being tested
    result, _ := myFunction(myMock)

    // assert that the function behaved as expected
    assert.Equal(t, 0, result)

    // ensure all expectations were met
    myMock.AssertExpectations(t)
}


5:设置 mock 对象的期望行为

myMock.On("DoSomething", 1, "arg").Return(0, nil)


6:调用 mock 对象的方法并断言结果

result, _ := myFunction(myMock)
assert.Equal(t, 0, result)

7:最后通过 AssertExpectations 方法






