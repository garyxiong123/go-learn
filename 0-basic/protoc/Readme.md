

tidy:
  go mod tidy

api-server:
   goctl api go -api server.api --dir .

.PHONY: api-server 




help              Help about any command
kube              Generate kubernetes files
migrate           Migrate from tal-tech to zeromicro
model             Generate model code
quickstart        quickly start a project
rpc               Generate rpc code
template          Template operation
upgrade           Upgrade goctl to latest version
