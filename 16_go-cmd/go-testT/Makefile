


##会自动读取源码目录下面名为 *_test.go 的文件，生成并运行测试用的可执行文件
#测试用例文件不会参与正常源码编译，不会被包含到可执行文件中。
#测试用例文件使用go test指令来执行，没有也不需要 main() 作为函数入口。所有在以_test结尾的源码内以Test开头的函数会自动被执行。
#测试用例可以不传入 *testing.T 参数
go-test:
	 go test

#-v 显示测试的详细命令。
go-test-v:
	 go test -v

#
go-test-v-timeout:
	 go test -v -run TestL2Suite -timeout 30m


#-cover 开启测试覆盖率；
go-test-cover:
	 go test -cover

#-run regexp 只运行 regexp 匹配的函数，例如 -run=Array 那么就执行包含有 Array 开头的函数；
go-test-regexp:
	 go test -run=Array

#名称支持正则表达式，使用-run TestA$
go-test-regexp1:
	 go test -run=ArrayA

#-bench regexp 执行相应的 benchmarks，例如 -bench=.；

go-test-benchmarks:
	 go test -run=Array


