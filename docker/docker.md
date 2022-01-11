# docker

- 帮助命令
  - version
  - info
  - help

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
