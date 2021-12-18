# container

容器即镜像的运行时, 命令后的 container 表示容器的 id 或名字。

# Index

- 常用命令
  - run
  - start
  - restart
  - stop
  - rm
  - kill
  - exec
  - attach
  - ps
  - cp
  - logs
  - top
  - inspect
  - commit

# 常用命令

- run [options] image: 运行指定镜像，`docker run -it centos /bin/bash`；

  - -p port:containerport: 以指定端口运行，80:8080 表示宿主机的 80 端口与容器内 8080 端口映射起来。
  - -P: 内部端口映射到宿主机的随机端口；
  - -it: i 表示交互模式运行，t 表示终端模式运行，通常配合使用，以交互方式运行，进入容器内查看内容；
  - -d: 后台模式运行,注意，后台运行时必须要有一个前台进程，否则容器自动停止。
  - --name: `--name CONTAINERName` 指定容器的名字;
  - -m num: 设置最大的运行内存，`run -m 3G|3MB...`;

- start CONTAINER: 启动一个已经停止的容器；
- restart CONTAINER: 重启一个容器；
- stop CONTAINER: 停止一个启动的容器；
- rm [options] CONTAINER: 删除一个容器；
  - -f: 默认情况下无法删除正在运行的容器，-f 表示强制删除。
- kill CONTAINER: 干掉一个容器；
- attach CONTAINER: 进入一个运行中的容器，并使用容器中正在运行的终端，由于 attach 使用的是运行中的终端，因此当退出时，会导致容器的停止。
- exec [options] CONTAINER: 新启动一个终端进入正在运行中的容器，优先使用 exec， 以 exec 进入后，退出容器时不会导致容器的停止。
  - -it: i 表示交互模式运行，t 表示终端模式运行，通常配合使用，以交互方式运行，进入容器内查看内容；
- ps [options]: 查看容器列表。

  - -a: 查看所有容器，包括已经停止的；
  - -q: 仅查看容器 ID。

- cp: 将服务器上指定文件拷贝进指定容器中，也可以将容器中指定文件拷贝到服务器中；

  - `cp path CONTAINERId:path`
  - `cp CONTAINERId:path path`

- logs [options] CONTAINER: 显示容器日志、输出；

  - t: 显示时间戳；
  - f: 跟踪实时日志；
  - n: 显示条数。

- top CONTAINER: 显示容器内的进程；

- inspect CONTAINER: 容器的元数据，以 json 格式返回。

- commit [OPTIONS] CONTAINER [REPOSITORY[:TAG]]: 提交一个容器，保存为镜像。
  - -a: 作者；
  - -m: 提交信息。
