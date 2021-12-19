# Network

默认情况下（bridge），每运行一个容器，docker 会自动为其分配一个内网 ip，从而使容器可以网络访问宿主机或其他容器。

# Index

- 分类
- network 命令
- 容器互联
- 自定义网络
- 补充

# 分类

Docker 的网络共分为 5 种，分别如下：

- bridge：默认的驱动模式，即“网桥”，通常用于单机（更准确地说，是单个 Docker 守护进程）
- overlay：Overlay 网络能够连接多个 Docker 守护进程，通常用于集群；
- host：直接使用主机（也就是运行 Docker 的机器）网络，仅适用于 Docker 17.06+ 的集群服务
- macvlan：Macvlan 网络通过为每个容器分配一个 MAC 地址，使其能够被显示为一台物理设备，适用于希望直连到物理网络的应用程序（例如嵌入式系统、物联网等等）
- none：禁用此容器的所有网络。

# network 命令

network 命令用于管理 Docker 的网络

- network ls: 查看所有网络；

```
NETWORK ID     NAME      DRIVER    SCOPE
0eb374dad3d2   bridge    bridge    local # docker bridge
52d626038d12   host      host      local # docker host
3e5a169c013d   none      null      local # docker none
[root@VM-0-5-centos git]#
```

- network inspect NETWORK: 查看指定网络的元数据，重点关注 IPAM、Containers 项，分别表示 IP 地址信息以及使用该网络的容器信息

```
[
    {
        "Name": "bridge",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": null,
            "Config": [
                {
                    "Subnet": "172.18.0.0/16",
                    "Gateway": "172.18.0.1"
                }
            ]
        },
        "Containers": {
            "70bc53e51b6250fab945f4d54befc819a07be707144b53ed41bde74f0e65b9e0": {
                "Name": "redis1",
                "EndpointID": "3f64370ed7dbdae79cf4244fe3f63de31857d6695a4213642ae135f0f3f79b0f",
                "MacAddress": "02:42:ac:12:00:03",
                "IPv4Address": "172.18.0.3/16",
                "IPv6Address": ""
            },
            "cf6fae0911e5a3224328fd2076c1ba718b32ed2ebcf75056a22de15671293563": {
                "Name": "mysql1", # container name
                "EndpointID": "e28f5f86f02568007e9783c4f9207961ca9951429ff77cb807881189d1f97dfe",
                "MacAddress": "02:42:ac:12:00:02",
                "IPv4Address": "172.18.0.2/16",
                "IPv6Address": ""
            }
        }
    }
]
```

- network prune: 删除不使用的网络；
- network rm NETWORK: 删除指定网络；

# 容器互联

- run --link CONTAINER2 CONTAINER1: 通过 --link 选项，可以使容器 1 直接通过访问容器 2 名称来进行网络访问 ，其本质是在容器 1 内 hosts 文件内加入容器 2 内网 ip 与 容器 2 名的映射；
  - 通过 --link CONTAINER 配置，可以在容器内直接以访问域名的方式访问另一个容器；
  - 目前已经不建议使用 --link 来配置容器互联了，该种方式太不灵活，不能实现双方之间通过“域名”形式的直接网络访问。

```
[root@VM-0-5-centos git]# docker run --link os2 --name os1 centos


[root@VM-0-5-centos git]# docker exec -it os1 cat /etc/hosts
127.0.0.1       localhost
::1     localhost ip6-localhost ip6-loopback
fe00::0 ip6-localnet
ff00::0 ip6-mcastprefix
ff02::1 ip6-allnodes
ff02::2 ip6-allrouters
172.18.0.2      os2 cf6fae0911e5
172.18.0.3      89af7a89c279

# 容器1，正常 ping 通
[root@VM-0-5-centos git]# docker exec -it os1 ping os2

# 容器2 host，无法 ping 通
[root@VM-0-5-centos git]# docker exec -it os2 ping os1
```

# 自定义网络

通过自定义配置网络，可以实现更高程度的定制化，以及实现统一网络内容器互联，通过“域名”（容器名）进行网络访问。

1. network create [OPTIONS] NETWORK: 创建自定义网络；

   - -d 设置网络驱动，默认 bridge；
   - --subnet: 分配子网络的网段，例如 192.168.0.0/16, 实际场景中，通常不同的集群，分配各自的子网段，从而实现网段隔离。（不同网段之间无法互联）

2. run --network NETWORK: 指定连接的网络，不添加该选项时，其实默认连接 docker 自带的 bridge 网络。

3. network connect NETWORK CONTAINER: 该命令用于将一些容器加入指定网络，通常用于将一些其他网段的容器加入指定网段，或是一些运行时忘记添加网络的容器追加网络。
