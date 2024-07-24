FROM nginx:stable

COPY ./docker/nginx.conf /etc/nginx/conf.d/default.conf
COPY ./php/index.php /var/www
