FROM php:7.4-cli as builder

WORKDIR /var/www

RUN apt update && apt upgrade -y && \
    apt install libzip-dev -y && \
    docker-php-ext-install zip &&\
    php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');" && \
    php -r "if (hash_file('sha384', 'composer-setup.php') === 'e21205b207c3ff031906575712edab6f13eb0b361f2085f1f1237b7126d785e826a450292b6cfd1d64d92e6563bbde02') { echo 'Installer verified'; } else { echo 'Installer corrupt'; unlink('composer-setup.php'); } echo PHP_EOL;" && \
    php composer-setup.php && \
    php -r "unlink('composer-setup.php');" && \
    php composer.phar create-project --prefer-dist laravel/laravel laravel

ENTRYPOINT [ "php","laravel/artisan","serve" ]

CMD [ "--host=0.0.0.0" ]

# outra imagem, prod por exemplo
FROM php:7.4-fpm-alpine

WORKDIR /var/www

# específico para este caso do php e laravel.
RUN rm -rf /var/www/html

COPY --from=builder /var/www .

# específico para este caso do php e laravel.
RUN chown -R www-data:www-data /var/www

EXPOSE 9000

# comando que vai executar o php
CMD ["php-fpm"]