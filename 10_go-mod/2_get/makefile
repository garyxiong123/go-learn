



##2.下载modules到本地cache 目前所有模块版本数据均缓存在 $GOPATH/pkg/mod和 $GOPATH/pkg/sum
#如果没有指明 version 的情况下，则默认先下载打了 tag 的 release 版本，
#比如 v0.4.5 或者 v1.2.3；如果没有 release 版本，则下载最新的 pre release 版本，比如 v0.0.1-pre1。如果还没有则下载最新的 commit
go-get-h:
	go help get

go-get:
	go get github.com/garyxiong123/go-depency-A


go-get-tag:
	go get github.com/garyxiong123/go-depency-A@v1.0.3

##To remove a dependency on a module and downgrade modules that require it:
go-remove:
	go get github.com/garyxiong123/go-depency-A@none
##更新到某个分支最新的代码
go-get-branch:
	go get github.com/garyxiong123/go-depency-A@test

##3.更新到最新的修订版（只改bug的版本）
go-get-fix:
	go get github.com/garyxiong123/go-learn@master
##记得更新缓存


##替代只能FQ下载的库
go-get-replace-repo:
	go mod edit -replace=golang.org/x/crypto@v0.0.0=github.com/golang/crypto@latest





go-list:
	go list -m all



##验证依赖是否正确
go-mod-verify:
	go mod verify

##查找依赖
go-find:
	go mod why all


