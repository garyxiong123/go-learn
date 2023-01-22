$ mkdir go-zero-demo
$ cd go-zero-demo
$ go mod init go-zero-demo
$ goctl api new greet
$ go mod tidy
Done.



启动服务

$ cd greet
$ go run greet.go -f etc/greet-api.yaml
输出如下，服务启动并侦听在8888端口：

Starting server at 0.0.0.0:8888...
访问服务

$ curl -i -X GET http://localhost:8888/from/you

测试
