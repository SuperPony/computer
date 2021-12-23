FROM golang:1.16.12-alpine AS builder

ENV GOPROXY=https://goproxy.cn \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /app

# 先复制 go mod 文件
# 利用 docker 缓存机制，判断是否有新的第三方依赖，从而判断是否需要重新 go mod download
COPY go.mod .
COPY go.sum .

RUN go mod download \
    && go mod tidy

# 由于重新构建必然是更新了代码
# 重新复制代码，进行编译
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o blog ./internal

FROM scratch

ENV SERVERPORT=8000
# 从 builder 镜像中二进制文件拷贝到目录中
COPY --from=builder /app/blog /

VOLUME [ "/" ]


EXPOSE $SERVERPORT

# 需要运行的命令
ENTRYPOINT ["/blog"]                             