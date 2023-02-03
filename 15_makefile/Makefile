# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.


api-server:
	@echo "--> api-server";




test: api-server
	@echo "--> Running go test"




integration-test:
	sudo bash -x ./local-test.sh

tools:
	go install -u github.com/zeromicro/go-zero/tools/goctl@v1.4.0

build: api-server build-only

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.0
	golangci-lint run ./...

build-only:
	go build -o build/bin/zkbnb -ldflags="-X main.version=${VERSION} -X main.gitCommit=${GIT_COMMIT} -X main.gitDate=${GIT_COMMIT_DATE}" ./cmd/zkbnb

docker-image:
	go mod vendor # temporary, should be removed after open source
	docker build . -t ${IMAGE_NAME}

.PHONY: api-server deploy integration-test test tools build lint build-only docker-image
