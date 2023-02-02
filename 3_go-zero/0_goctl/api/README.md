$ goctl api -h
NAME:
goctl api - generate api related files

USAGE:
goctl api command [command options] [arguments...]

COMMANDS:
new       fast create api service
format    format api files
validate  validate api file
doc       generate doc files
go        generate go files for provided api in yaml file
java      generate java files for provided api in api file
ts        generate ts files for provided api in api file
dart      generate dart files for provided api in api file
kt        generate kotlin code for provided api file
plugin    custom file generator

OPTIONS:
-o value        output a sample api file
--home value    the goctl home path of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
--remote value  the remote git repo of the template, --home and --remote cannot be set at the same time, if they are, --remote has higher priority
The git repo directory must be consistent with the https://github.com/zeromicro/go-zero-template directory structure
--branch value  the branch of the remote repo, it does work with --remote
--help, -h      show help