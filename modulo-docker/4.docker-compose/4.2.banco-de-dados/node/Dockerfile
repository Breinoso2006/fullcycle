FROM node:latest

WORKDIR /usr/src/app

EXPOSE 3000

RUN npm install express && \
    npm install mysql2

RUN apt update && \
    apt install -y wget netcat-traditional && \
    wget -q -O /usr/bin/wait-for https://raw.githubusercontent.com/eficode/wait-for/v2.2.3/wait-for && \
    chmod +x /usr/bin/wait-for
