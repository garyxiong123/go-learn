# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

GOBIN?=${GOPATH}/bin

api:
	cd /Users/user/sdk/go_path/src/go-learn/go_ctrl && goctl api go -api server.api --dir .;
	@echo "Done generate server api";

ts:
	cd /Users/user/sdk/go_path/src/go-learn/go_ctrl && goctl api ts -api server.api --dir .;
	@echo "Done generate server api";
java:
	cd /Users/user/sdk/go_path/src/go-learn/go_ctrl && goctl api java -api server.api --dir .;
	@echo "Done generate server api";

doc:
	cd /Users/user/sdk/go_path/src/go-learn/go_ctrl && goctl api doc --dir .;
	@echo "Done generate server api";

.PHONY: api ts doc
