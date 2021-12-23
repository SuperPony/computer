# Docker-Compose

Docker-Compose 是一种容器编排技术，用于一次性对多个镜像进行操作；其实现主要通过编写 docker-compose.yml 进行。

# Index

- 安装
  - 删除
- 备注
- docker-compose.yml

# 安装

1. `$ sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose` 1.29.2 可以根据需要换为其他版本；

2. `$ sudo chmod +x /usr/local/bin/docker-compose`。

## 删除

`$ sudo rm /usr/local/bin/docker-compose`.

# 备注

- docker-compose 挂载目录数据卷流程：
  1. 启动镜像阶段会执行一次文件挂载
  2. 如果宿主机不存在该目录，会新建空的文件夹
  3. 然后将宿主机目录的内容覆盖容器内的内容
  - 该流程导致第一次运行时容器内对应的挂载目录全部清空

# docker-compose.yml

docker-compose.yml 用于声明编排该如何进行；此处列举常用的配置项，具体配置项参考官方 https://docs.docker.com/compose/compose-file/compose-file-v3。

docker-compose.yml 主要分为三大块

1. version: 指 docker-compose 构建所依赖的版本;
2. services: services 代表需要构建的服务镜像；
3. other: other 通常代表一些其他的配置，通常是辅助 services，例如创建 services 需要的自定义网络、数据卷等。

```
# 构建基于的 docker-compose 版本
version: "3.7"

# 构建指令
services:
  # 服务名称
  nginx:
    # 该服务所基于的镜像 image: IMAGE[:tag]
    image: nginx
    # 对外暴露的端口
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - "nginx-config: /etc/nginx"

    # 声明运行时的容器名，默认命名格式为 docker-compose.yml所在目录名_nginx_num, num 代表编号，在集群中有意义。
    container_name: "my-nginx"

    # 声明构建顺序，该服务构建晚于 db；此配置用于确保对于构建顺序有严格要求的应用能够正确构建
    depends_on:
      - db

    # 该命令声明服务运行时加入指定网络
    # 默认情况下在启动编排时，会自动创建一个网络，将服务的所有容器全部连接该网络
    networks:
      - my-network
      - web-network

  db:
    build:
      # 设置构建时所处的上下文
      context: ./
      # Dockerfile 所在目录
      dockerfile: ./Docker-mysql
      # 设置构建的镜像名称，默认命名格式为 docker-compose.yml所在目录名_db
      target: app-db

    ports:
      - "33066:3306"
    # 设置环境变量
    environment:
      - MYSQL_ROOT_PASSWORD=12345678
    # 数据卷挂载，挂载模式与容器挂载数据卷一致
    # 此处具名挂载不会自动创建数据卷，需要在全局的 volumes 声明创建
    volumes:
      - "mysql-config:/etc/mysql"
      - "mysql-data:/var/lib/mysql"
    # 声明容器暴露的端口
    # expose:
    #   - "80"

    # 服务运行时执行的命令
    command: "--character-set-server=utf8"


# 其他指令

# 创建具名数据卷
volumes:
  nginx-conf:
    # 声明数据卷名称，默认命名格式为 docker-compose.yml所在目录名_nginx-conf
    name: "nginx-conf"
  mysql-conf: null
  mysql-data:
    # 声明使用已经存在的 mysql-data 数据卷
    external: true

# 创建自定义网络
networks:
  my-network:
    external: true #  表明使用已经存在的 my-network 网络
  web-network:
    # 声明网络驱动类型,默认驱动即 bridge
    driver: bridge
    # 声明自定义网络的名称，不加该选项，则实际生成的自定义网络名称为 docker-compose.yml所在目录名_web-network
    name: web-network
```
