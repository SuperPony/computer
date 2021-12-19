# Dockerfile

Dockerfile 文件用于构建镜像, 每一条命令都将作为一层记录。

# 常用命令

- `FORM IMAGE[:tag]` 该镜像所基于的镜像;
- `LABEL key="val"` 自定义标签，通常存放一些关于镜像的信息；
- `RUN command` 表示镜像构建时执行的命令;
- `CMD [ "executable","param",...]` 容器运行时，执行的命令；

  - param 表示选项，param 可选；
  - Dockerfile 中，如果存在多个 CMD，则最后一个生效；
  - CMD 会被被运行时设置的命令覆盖；

  ```
  Dockerfile

  CMD [ "echo","hello" ] # 不生效
  CMD [ "ls", "-al" ]

  // echo thanks 命令覆盖掉 CMD 设置的 ls -al 命令，最终运行容器时，输出 thanks
  [root@VM-0-5-centos ~]# docker run IMAGE echo thanks
  ```

- `ENTRYPOINT [ "executable","param",...]`: 与 CMD 类似，区别在于，运行时设置的命令追加为 ENTRYPOINT 设置的命令的选项；

  - 多个 ENTRYPOINT 仅最后一个生效；
  - CMD 可以配合 ENTRYPOINT 使用， CMD 作为 ENTRYPOINT 的默认选项存在。

  ```
  Dockerfile

  ENTRYPOINT [ "nginx" , "-c"]
  CMD [ "/etc/nginx/nginx.conf"]

  // 执行效果 nginx -c /etc/nginx/nginx.conf
  [root@VM-0-5-centos ~]# docker run IMAGE

   // 执行效果 nginx -c /etc/nginx/demo.conf
  [root@VM-0-5-centos ~]# docker run IMAGE /etc/nginx/demo.conf
  ```

- `EXPOSE port`: 声明容器对外暴露的端口，主要用于使用时做端口映射；

  - `EXPOSE 7000-8000`: 暴露范围端口;
  - `EXPOSE 80 443`: 暴露指定端口.

- `WORKDIR /the/workdir/path`: 设置进入容器后的工作目录；

- `ENV key="val" key2="val2"...`: 设置容器环境变量

- `COPY [--chown=<user>:<group>] <源路径1>... <目标路径>` 在构建时将指定宿主机文件拷贝至容器内目标路径；

  - `--chown=<user>:<group>`: 表示文件所属用户:用户组；
  - 目标路径可以不存在，构建时，自动创建。

- `ADD [--chown=<user>:<group>] <源路径1>... <目标路径>` 与 COPY 命令类似，区别在于如果源文件是压缩包，则会自动解压至目标路径；

  - 最佳实践：如果源文件不是解压包，则使用 COPY。

- `VOLUME [ "/data",... ]` 定义匿名数据卷。

- `ONBUILD INSTRUCTION` 延迟构建的命令，声明该命令的 Dockerfile 不会执行该命令，当有其他 Dockerfile 继承该镜像时，则在构建时执行 ONBUILD 中的命令；

```
# Dockerfile1 假定镜像声明为 test

FROM centos

ONBUILD CMD [ "ls", '-al' ]
ONBUILD VOLUME [ "/data2" ]

# Dockerfile2 该镜像在构建时，执行 test 镜像中 ONBUILD 命令

FORM test

```

# Guide

```
FROM  centos:7

LABEL author="superpony"
LABEL email="superponyyy@gmail.com"
LABEL version=0.1.0

# RUN yum -y install vim
# RUN yum -y install lsof

# RUN yum -y install vim\
#     yum -y install lsof

RUN yum -y install vim\
    &&yum -y install lsof

ENTRYPOINT [ "nginx" , "-c"]

CMD [ "/etc/nginx/nginx.conf"]

EXPOSE 80 443

# ENV WORKPATH="/var/lib"\
#    LOGPATH="/var/log"

VOLUME [ "/var/lib" ]
172.18.0.3 172.18.0.2
WORKDIR /home
```
