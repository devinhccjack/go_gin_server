# FROM golang:1.17 as build

# WORKDIR /go/src/goproject

# COPY . .
# # RUN ls

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOPROXY=https://goproxy.cn,direct 
# RUN go build -ldflags="-w -s" -o goproject .


# FROM ubuntu:18.04
# WORKDIR /usr/local/demo
# COPY --from=build /go/src/project/goproject .
# EXPOSE 9091
# CMD ["./main.go"]


# #编译和运行都是用这个镜像
# FROM golang:alpine

# #设置工作路径
# WORKDIR /build
# RUN pwd

# #把项目所有文件复制到镜像的根目录文件夹中
# ADD . ./

# # 设置Go语言的环境变量，打开Go Module模式。设置包下载源，有利于快速下载包
# ENV GO111MODULE=on \
# GOPROXY=https://goproxy.cn

# #下载go.mod里面的包
# RUN go mod download

# # RUN go mod tidy
# # RUN go build .

# #编译
# RUN go build -o  .

# #打开端口
# EXPOSE 8080

# #运行项目
# ENTRYPOINT  ["./build"]


FROM golang:1.18

WORKDIR /src/gin_blog

COPY . /src/gin_blog

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn"
RUN go mod download
RUN go build main.go
RUN go run -x main.go


EXPOSE 8080

# ENTRYPOINT ["./gin_log"]
CMD ["./gin_log/main.go"]