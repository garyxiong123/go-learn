go mod init Gone




go.mod 提供了module, require、replace和exclude 四个命令

module 语句指定包的名字（路径）
require 语句指定的依赖项模块
replace 语句可以替换依赖项模块
exclude 语句可以忽略依赖项模块



go module 安装 package 的原則是先拉最新的 release tag，若无tag则拉最新的commit

go 会自动生成一个 go.sum 文件来记录 dependency tree





可以使用命令 go list -m -u all 来检查可以升级的package，使用go get -u need-upgrade-package 升级后会将新的依赖版本更新到go.mod * 也可以使用 go get -u 升级所有依赖