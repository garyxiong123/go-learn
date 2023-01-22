# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.


api-server:
	cd $(API_SERVER) && goctl api go -api server.api --dir .;
	@echo "Done generate server api";


.PHONY: api-server
