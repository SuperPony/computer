FROM php:7.4-fpm

LABEL author.name="SupeProny" \
    author.email="superponyyy@gmail.com"
ENV TZ=Asia/Shanghai

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone \
    && echo $TZ > /etc/timezone \
    && pecl install redis \
    && pecl install xdebug-2.8.1 \
    && pecl install swoole\
    && docker-php-ext-enable redis xdebug

