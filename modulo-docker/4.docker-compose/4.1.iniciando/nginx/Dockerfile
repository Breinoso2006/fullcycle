FROM nginx:1.15.0-alpine

# alterando configuração padrão do nginx
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/conf.d

# criando pasta e arquivo php para ter algo para o nginx ter algo para chamar
RUN mkdir /var/www/laravel/public -p && touch /var/www/laravel/public/index.php
