FROM php:7.4-fpm-alpine

LABEL Author ="SupeProny" \
      Email="superponyyy@gmail.com"

ENV TZ=Asia/Shanghai

# 设置时区
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone

# 更换国内镜像
RUN mv /etc/apt/sources.list /etc/apt/sources.list.bak \
    && echo 'deb http://mirrors.163.com/debian/ stretch main non-free contrib' > /etc/apt/sources.list \
    && echo 'deb http://mirrors.163.com/debian/ stretch-updates main non-free contrib' >> /etc/apt/sources.list \
    && echo 'deb http://mirrors.163.com/debian-security/ stretch/updates main non-free contrib' >> /etc/apt/sources.list \
    && apt-get update

# bcmath, calendar, exif, gettext, sockets, dba,
# mysqli, pcntl, pdo_mysql, shmop, sysvmsg, sysvsem, sysvshm 扩展
Run docker-php-ext-install -j$(nproc) bcmath calendar exif \
    gettext sockets dba mysqli pcntl pdo_mysql shmop sysvmsg \
    sysvsem sysvshm iconv

# GD 扩展
RUN apt-get install -y --no-install-recommends libfreetype6-dev libjpeg62-turbo-dev libpng-dev \
    && rm -r /var/lib/apt/lists/* \
    && docker-php-ext-configure gd --with-freetype-dir=/usr/include/ --with-jpeg-dir=/usr/include/ \
    && docker-php-ext-install -j$(nproc) gd

# imagick 扩展
RUN export CFLAGS="$PHP_CFLAGS" CPPFLAGS="$PHP_CPPFLAGS" LDFLAGS="$PHP_LDFLAGS" \
    && apt-get install -y --no-install-recommends libmagickwand-dev \
    && rm -r /var/lib/apt/lists/* \
    && pecl install imagick-3.4.4 \
    && docker-php-ext-enable imagick

# mcrypt 拓展
Run apt-get install -y --no-install-recommends libmcrypt-dev \
        && rm -r /var/lib/apt/lists/* \
        && pecl install mcrypt-1.0.2 \
        && docker-php-ext-enable mcrypt

# redis、opcache、xdebug、 swoole 扩展
RUN pecl install redis-5.0.0 && docker-php-ext-enable redis \
    docker-php-ext-configure opcache --enable-opcache \
    && docker-php-ext-install opcache \
    && install xdebug-2.7.2 && docker-php-ext-enable xdebug \
    && pecl install swoole-4.4.0 && docker-php-ext-enable swoole