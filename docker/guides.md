# 最佳实践

# Image 最佳实践

- 使用分段构建的方法，减轻镜像的实际大小；
- 基础镜像尽可能选择轻的版本，例如 alpine 版本的镜像；

# Dockerfile 最佳实践

- 命令优先使用追加的方式编写，在 Dockerfile 中插入命令，会导致后面的镜像层缓存全部失效；
- 命令优先使用联合的方式声明，以减少镜像层；

- `COPY` 命令具有检测目标文件内容变化的机制，当内容没有发生变化时，镜像层直接走缓存；目标文件发生变化，则`COPY`镜像层缓存失效，引发后面镜像层失效；对于一些依赖第三方工具的应用来说，该机制能确保当项目依赖变化时，构建的镜像会正常的下载依赖。最佳的实践是，将项目的第三方依赖管理文件，通过`COPY`命令优先处理，然后再处理代码。

```
FROM golang:1.16.12-alpine AS builder

# go.mod、 go.sum 优先复制，当内容发生变动时，触发下载；反之不触发。
COPY go.mod .
COPY go.sum .

# 联合形式声明指令，减少镜像层
RUN go mod download \
    && go mod tidy

COPY . .

RUN go build -o app-api ./internal

# 最终镜像
FROM scratch

# 从 builder 镜像中将编译后的应用直接拷贝到目录中
COPY --from=builder /app/app-api /

ENTRYPOINT ["/app-api"]
```

# Container 最佳实践

- 1 个容器 1 个进程原则；
- 容器是轻量级、无状态的，故而通过数据卷的方式存储容器状态。
