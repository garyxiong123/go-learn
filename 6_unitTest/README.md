
reference the address

https://pkg.go.dev/testing#hdr-Benchmarks




GO 语言里面的单元测试，是使用标准库 testing

有如下简单规则：

导入 test 标准库
单测文件名，后面跟上_test
单测文件中的函数名为 Test开头，且参数必须是 t *testing.T