## FROM ... AS builder : 表示依赖的镜像只是使用在编译阶段
#FROM golang:1.18.1 AS builder
#
## 编译阶段的工作目录，也可以作为全局工作目录
#WORKDIR /app
#
## 把当前目录的所有内容copy到 WORKDIR指定的目录中
#COPY . .
#
## 定义go build的工作环境，
## 例如GOOS=linux、GOARCH=amd64，
## 这样编译出来的 'main可执行文件' 就只能在linux的amd64架构中使用
#ARG TARGETOS
#ARG TARGETARCH
#
## 执行go build； --mount：在执行build时，会把/go 和 /root/.cache/go-build 临时挂在到容器中
#RUN --mount=type=cache,target=/go --mount=type=cache,target=/root/.cache/go-build \
#    GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o main
#
#FROM alpine:3.14.0
#
## 把执行builder阶段的结果 /app/main拷贝到/app中
#COPY --from=builder /app/main /app
## 把配置文件copy到/app/tsvbin中
#COPY ./tsvbin /app/tsvbin
#
## 运行main命令，启动项目
## /app/main 指向RUN命令的 go build -o main的结果
#ENTRYPOINT ["/app/main"]


FROM golang:1.18-alpine3.16 AS builder

WORKDIR /build
RUN adduser -u 10001 -D app-runner

ENV GOPROXY https://goproxy.cn
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o httpserver .

FROM alpine:3.16 AS final

WORKDIR /app
COPY --from=builder /build/httpserver /app/
#COPY --from=builder /build/config /app/config
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER app-runner

ENTRYPOINT ["/app/httpserver"]