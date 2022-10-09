
### 初始化模块
```bush
go mod init ProjectName
go mod init go_server

```

### 下载依赖包
```bush
go mod tidy
go env //查看配置信息
set GO111MODULE=on / off
go env -w GO111MODULE=off
go env -w GO111MODULE=on

```

### 安装
```bush
go get -u -v github.com/jinzhu/gorm
go get -u -v github.com/gin-contrib/sessions

```

### 下载:
```bush
go get   github.com/jinzhu/gorm
go get   github.com/go-sql-driver/mysql
```

### model 定义表结构体时，加入
```bush
gorm.Model
新增数据时会自动维护表的 created_at 和 updated_at 字段
init函数在你导入该package时程序会自动调用init函数

```
### 下载
```bush
go get XXX 下载的包，默认会安装在GOPATH 的第一个路径里。

```

### 参考文章
```bush

gorm中文文档
https://learnku.com/docs/gorm/v2/models/9729

Grom之数据查询
https://blog.csdn.net/weixin_52690231/article/details/124542785?spm=1001.2101.3001.6650.1&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-124542785-blog-126673506.pc_relevant_3mothn_strategy_and_data_recovery&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-124542785-blog-126673506.pc_relevant_3mothn_strategy_and_data_recovery&utm_relevant_index=2

Go语言框架Gin之5 数据库（原生数据库、gorm和xorm）
https://blog.csdn.net/weixin_42117918/article/details/111561023?utm_medium=distribute.pc_relevant.none-task-blog-2~default~baidujs_baidulandingword~default-0-111561023-blog-117017396.pc_relevant_layerdownloadsortv1&spm=1001.2101.3001.4242.1&utm_relevant_index=3

gorm连接mysql数据库以及建表和自动迁移
https://blog.csdn.net/qq_42956653/article/details/124783019

```
### 接口api
```bush

接口：提交数据
http://localhost:8080/v1/testinsert
参数：
{
    "title": "title",
    "testcol":"testcol",
    "Status": 1,
    "description":"description"
}

response:
{
    "code": 1,
    "data": 3,
    "message": "success"
}


接口：获取数据列表
http://localhost:8080/v1/testlist

response:
{
    "code": 1,
    "data": [
        {
            "Id": 1,
            "Testcol": "testcol",
            "Title": "title",
            "Description": "description",
            "Status": 1,
            "CreatedAt": "2022-10-09T23:03:56+08:00",
            "UpdatedAt": "2022-10-09T23:03:56+08:00",
            "DeletedAt": null
        },
        {
            "Id": 2,
            "Testcol": "testcol",
            "Title": "title",
            "Description": "description",
            "Status": 1,
            "CreatedAt": "2022-10-09T23:04:16+08:00",
            "UpdatedAt": "2022-10-09T23:04:16+08:00",
            "DeletedAt": null
        },
        {
            "Id": 3,
            "Testcol": "testcol",
            "Title": "title",
            "Description": "description",
            "Status": 1,
            "CreatedAt": "2022-10-09T23:04:17+08:00",
            "UpdatedAt": "2022-10-09T23:04:17+08:00",
            "DeletedAt": null
        }
    ],
    "message": "success"
}


```

### 构建镜像docker
```bush

http://localhost:8080/ping

go run main.go

go build main.go

docker build -t gin:v1 .

docker run -itd --name gin -p 8080:8080 gin:v1

```