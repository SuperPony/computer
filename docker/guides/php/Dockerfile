FROM php:7.4-fpm-alpine

LABEL Author ="SupeProny" \
    Email="superponyyy@gmail.com"

ENV TZ=Asia/Shanghai

# 设置时区
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone

# 阿里云镜像加速
RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g" /etc/apk/repositories \
    && apk update


# PHPIZE_DEPS 包含 gcc g++ 等编译辅助类库，完成编译后删除
# docker-php-ext-enable 指令用于开启 pelc 添加的拓展
RUN apk add --no-cache $PHPIZE_DEPS \
    && apk add --no-cache libstdc++ libzip-dev vim\
    && apk update \
    && pecl install redis-5.3.4 \
    && pecl install zip \
    && pecl install swoole \
    && docker-php-ext-enable redis zip swoole \
    && apk del $PHPIZE_DEPS

# docker-php-ext-install 指令是 docker php 镜像提供的用于安装拓展的指令，并且已经包含编译辅助类库的删除逻辑
RUN apk add --no-cache freetype libpng libjpeg-turbo freetype-dev libpng-dev libjpeg-turbo-dev \
    && apk update \
    && docker-php-ext-configure gd --with-freetype --with-jpeg \
    && docker-php-ext-install -j$(nproc) gd \
    && docker-php-ext-install -j$(nproc) pdo_mysql \
    && docker-php-ext-install -j$(nproc) opcache \
    && docker-php-ext-install -j$(nproc) bcmath \
    && docker-php-ext-install -j$(nproc) mysqli 

# 镜像中，自带 php.ini-production，php.ini-development 两份为不同环境配置的配置文件
# php.ini-production 开启了一系列参数优化
# PHP_INI_DIR 环境变量是 php 镜像中指向配置文件的目录的环境变量
RUN mv "$PHP_INI_DIR/php.ini-production" "$PHP_INI_DIR/php.ini"

EXPOSE 9000