# 镜像分层

# Index

- 说明
  - 容器和镜像的区别
  - 基础镜像
- CopyOnWrite
- 最佳实践

# 说明

在 Docker 中，镜像就像一栋楼房，是由一系列镜像以层为单位并以堆栈形式构成的；在 Dockerfile 中，每一条指令即是一层。构建镜像时，当 Docker 发现本地拥有该层时，则会复用该层，这种机制一方面使得构建更加高效，另一方面也极大程度的复用了现有资源。

Docker 会为每一层镜像生成唯一的 ID

![](/docker/image/setp1.jpg)

层分为两类：

- 镜像层：构建时生成的一系列层级的统称，镜像层只读；
- 容器层：在容器启动时基于镜像层动态添加的层，该层可读可写，容器内部的数据、文件操作，便是在该层进行的；当容器删除时，该层随之删除。
  - 相同镜像的多个容器间，容器层相互独立，如果需要实现数据共享，则需要使用数据卷。

```
// Dockerfile

FROM ubuntu:18.04
COPY . /app
RUN make /app
CMD python /app/app.py
```

![](/docker/image/setp5.jpeg)

## 容器和镜像的区别

二者之间的主要区别便在于读写层，镜像是不具备读写层的，而容器则在镜像的基础层之上，拥有了读写层，容器内的数据操作，便是在该读写层进行。

容器 = 镜像 + 读写层

![](/docker/image/setp6.jpeg)

## 基础镜像

基础镜像指当前镜像所依赖的基本镜像，即 `FROM image:[tag]` 所声明的依赖镜像。

在构建时，当一个 Dockerfile 中，有多个 `FROM` 时，只会将最后一个 `FROM` 开始声明的指令作为镜像，但后面的指令可以使用前面的资源。通过这种机制，可以有效的减轻镜像的实际体积，这种方式被称为多段构建。

```
FROM golang:1.16.12-alpine AS builder

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o app-api ./internal

# 最终镜像
FROM scratch

# 从 builder 镜像中将编译后的应用直接拷贝到目录中
COPY --from=builder /app/app-api /

ENTRYPOINT ["/app-api"]
```

# CopyOnWrite

写时复制机制是 Docker 的核心机制之一，正如上文所讲，同一镜像产生的多个容器之间拥有各自独立的读写层，同时共用镜像层；但当其中一个容器对文件进行编辑时（例如 /etc 下的某个文件），为了保证镜像层只读以及不影响其他容器的原则，故而 Docker 是将该文件从镜像层复制给对应容器的读写层，从而保证了镜像的完整以及容器之间的隔离。
