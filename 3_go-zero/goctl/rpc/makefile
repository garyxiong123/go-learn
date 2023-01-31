#1: goctl rpc new greet
new:
	goctl rpc new greet

#2:生成proto模板

generate-template:
	goctl rpc template -o=user.proto

#3:Generate rpc service code

generate-rpc-code:
	goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=.