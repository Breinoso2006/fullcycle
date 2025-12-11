# Build stage
FROM golang:1.23 AS build

WORKDIR /app

# Copia arquivos de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o código fonte
COPY . .

# Ajusta dependências e build da aplicação
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cloudrun ./cmd/main.go

# Runtime stage
FROM scratch

WORKDIR /app

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/cloudrun .

EXPOSE 8080

ENV PORT=8080
ENV weatherKey="{api-key}"

ENTRYPOINT ["./cloudrun"]
