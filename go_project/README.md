
### 安装go环境
brew install go
$ go version

在 ~/.bashrc 中添加 GOPATH 变量

export GOPATH=~/go
export PATH=$PATH:$GOPATH/bin

source ~/.bashrc

go get -u -v github.com/ramya-rao-a/go-outline

git clone https://github.com/golang/tools.git $GOPATH/src/golang.org/x/tools
go get -v github.com/ramya-rao-a/go-outline
go get -v github.com/sqs/goreturns
go get -v github.com/rogpeppe/godef

### 安装gin
gin go get -u -v github.com/gin-gonic/gin
-v：打印出被构建的代码包的名字
-u：已存在相关的代码包，强行更新代码包及其依赖包

### 运行
go run main.go

### api
http://localhost:8080/admin/login
http://localhost:8080/admin/list?name=hcc
http://localhost:8080/list?name=hcc&id=2

http://localhost:8080/admin/morelist
{
    "code": 200,
    "message": "读取成功",
    "data": [
        {
            "id": "1",
            "pass_word": "aaa",
            "user_name": "hechong"
        },
        {
            "id": "2",
            "pass_word": "aaa",
            "user_name": "zhangsan"
        },
        {
            "id": "3",
            "pass_word": "12345",
            "user_name": "hcc1"
        },
        {
            "id": "4",
            "pass_word": "qwerty",
            "user_name": "hechong-hj"
        }
    ]
}


# 构建gin docker镜像

### 尝试运行
go run gindemo.go

### 安装依赖
go get github.com/gin-gonic/gin

### 执行构建镜像
在本地把go程序编译好，再通过docker的方式拷贝到容器镜像中,那么本地环境是存在差异的
在容器镜像中，先是安装go环境，然后配置go可执行文件环境变量，最后编译，强耦合在一起了
docker build . -t morebuild:test
### 检测镜像是否构建成功
# docker images|grep morebuild

### 运行容器
docker run -it -d -p 9091:9091 morebuild:test

### curl访问
curl http://127.0.0.1:9091/ping

### 查看容器日志
# docker logs -f 45f54
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)
[GIN-debug] GET    /ping                     --> main.main.func1 (3 handlers)
[GIN-debug] Listening and serving HTTP on 0.0.0.0:9091
[GIN] 2021/10/29 - 06:23:43 | 200 |      58.471µs |      172.17.0.1 | GET      "/ping"

### 测试访问
http://localhost:9091/ping

docker build -t gin:v1 .

docker run -itd --name gin -p 8080:8080 gin:v1

go run main.go

go build main.go

go run -x main.go

go build -o main.exe main.go

go run -x --work main.go

