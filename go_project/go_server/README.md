
# 初始化模块
go mod init ProjectName
go mod init go_server
# 下载依赖包
go mod tidy

go env //查看配置信息
set GO111MODULE=on / off
go env -w GO111MODULE=off
go env -w GO111MODULE=on



### 安装
go get -u -v github.com/jinzhu/gorm

go get -u -v github.com/gin-contrib/sessions


### 下载:

go get   github.com/jinzhu/gorm

go get   github.com/go-sql-driver/mysql

### model 定义表结构体时，加入

gorm.Model
新增数据时会自动维护表的 created_at 和 updated_at 字段

init函数在你导入该package时程序会自动调用init函数

### 下载
go get XXX 下载的包，默认会安装在GOPATH 的第一个路径里。



