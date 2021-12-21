#!/bin/bash

# 删除旧版 docker
sudo yum remove docker \
  docker-client \
  docker-client-latest \
  docker-common \
  docker-latest \
  docker-latest-logrotate \
  docker-logrotate \
  docker-engine

# 删除原有 docker-compose
sudo rm /usr/local/bin/docker-compose

sudo yum install -y yum-utils

# 添加 yum 源
sudo yum-config-manager \
  --add-repo \
  https://download.docker.com/linux/centos/docker-ce.repo

# 安装 docker
sudo yum install -y docker-ce docker-ce-cli containerd.io

# 配置阿里云镜像加速
sudo mkdir -p /etc/docker
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://zjwcef2k.mirror.aliyuncs.com"]
}
EOF

# 安装 docker-compose
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# 重载 docker
sudo systemctl daemon-reload
sudo systemctl restart docker