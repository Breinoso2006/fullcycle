FROM mysql:latest

CMD [ "--innodb-use-native-aio=0" ]

COPY create-database.sql /docker-entrypoint-initdb.d/