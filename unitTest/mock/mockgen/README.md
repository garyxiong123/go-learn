
使用 mockgen 来生成 mock 的过程：

1:安装 mockgen:


go get github.com/golang/mock/mockgen



2:使用 mockgen 来生成 mock 代码:


mockgen -source=myinterface.go -destination=myinterface_mock.go



3:在测试文件中引用 mock 代码


import "myinterface_mock"

myMock := new(MyInterfaceMock)
