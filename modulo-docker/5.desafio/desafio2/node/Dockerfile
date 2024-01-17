FROM node:latest

WORKDIR /usr/src/app

COPY index.js .

RUN npm install express mysql2 @ngneat/falso

RUN apt update && \
    apt install -y wget netcat-traditional && \
    wget -q -O /usr/bin/wait-for https://raw.githubusercontent.com/eficode/wait-for/v2.2.3/wait-for && \
    chmod +x /usr/bin/wait-for
