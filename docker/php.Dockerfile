FROM php:8.3-fpm-alpine

RUN set -ex \
    && apk update \
    && apk add libcurl curl-dev autoconf gcc g++ linux-headers make sudo \
    && docker-php-ext-install mysqli pdo_mysql curl opcache

COPY ./php/index.php /var/www
