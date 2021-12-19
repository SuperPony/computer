# Docker-Compose

Docker-Compose 是一种容器编排技术，用于一次性对多个镜像进行操作；其实现主要通过编写 docker-compose.yml 进行。

# Index

- 安装
  - 删除
- docker-compose.yml
- 命令
- 备注

# 安装

1. `$ sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose` 1.29.2 可以根据需要换为其他版本；

2. `$ sudo chmod +x /usr/local/bin/docker-compose`。

## 删除

`$ sudo rm /usr/local/bin/docker-compose`.

# docker-compose.yml

docker-compose.yml 用于声明编排该如何进行；
