# command

记录 docker 常用命令

# Index

- 帮助命令
  - version
  - info
  - help
- 镜像命令
  - images
  - search
  - pull
  - rmi
- 容器命令

# 帮助命令

- docker version: 显示当前 docker 版本；
- docker info|-i: 显示 docker 的详细信息;

```
Client:
 Context:    default
 Debug Mode: false
 Plugins:
  app: Docker App (Docker Inc., v0.9.1-beta3)
  buildx: Docker Buildx (Docker Inc., v0.7.1-docker)
  scan: Docker Scan (Docker Inc., v0.12.0)

Server:
  Containers: 2 // 容器数量
  Running: 0
  Paused: 0
  Stopped: 2
Images: 1 // 镜像数量
Server Version: 20.10.12  // docker 版本
......
```

- docker help|-h docker 帮助文档
  - docker [options] COMMAND -h: 指定命令的帮助文档

# 镜像命令

- docker images [options]: 查看本地镜像列表；
  - -a: 列出所有镜像;
  - -q: 仅展示镜像 ID。

```
// 镜像名称 镜像标签（latest 表示最新版） 镜像 ID 创建时间 镜像大小
REPOSITORY()    TAG       IMAGE ID       CREATED        SIZE
hello-world   latest    feb5d9fea6a5   2 months ago   13.3kB
```

- docker search name: 搜索指定名称的镜像；

```
[root@VM-0-5-centos docker]# docker search mysql
NAME                              DESCRIPTION                                     STARS     OFFICIAL   AUTOMATED
mysql                             MySQL is a widely used, open-source relation…   11803     [OK]
mariadb                           MariaDB Server is a high performing open sou…   4492      [OK]
mysql/mysql-server                Optimized MySQL Server Docker images. Create…   885                  [OK]
phpmyadmin                        phpMyAdmin - A web interface for MySQL and M…   391       [OK]
centos/mysql-57-centos7           MySQL 5.7 SQL database server                   92
mysql/mysql-cluster               Experimental MySQL Cluster Docker images. Cr…   89
centurylink/mysql                 Image containing mysql. Optimized to be link…   59                   [OK]
databack/mysql-backup             Back up mysql databases to... anywhere!         54
```

- docker pull name[:tag]: 下载指定镜像;

  - tag 表示指定的版本，默认是 latest（最新）版， 指定的版本必须是仓库中存在的。

- docker rmi name[:tag] | IMAGE ID: 删除指定镜像，批量删除以空格分割。
