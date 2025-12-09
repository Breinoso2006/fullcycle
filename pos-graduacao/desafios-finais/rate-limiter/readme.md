# Rate Limiter API

Sistema de rate limiting em Go para controlar o tráfego de requisições HTTP com base em endereço IP ou token de acesso (API Key), utilizando **Redis** como mecanismo de persistência.

## 📋 Descrição

Este projeto implementa um rate limiter configurável que pode ser utilizado como middleware em servidores web Go. O sistema limita o número de requisições por segundo com base em dois critérios:

- **Por IP**: Limita requisições de um mesmo endereço IP
- **Por Token**: Limita requisições de um token de acesso específico (header `API_KEY`)

**Prioridade**: Quando uma requisição possui token, o limite do token sobrepõe o limite do IP.

**Persistência**: Utiliza **Redis** para armazenar contadores e bloqueios, permitindo que o rate limiting funcione de forma distribuída entre múltiplas instâncias da aplicação.

## 🏗️ Arquitetura

O projeto segue princípios de Clean Architecture e utiliza o **Strategy Pattern** para permitir troca fácil do mecanismo de persistência (caso necessário no futuro):

```
├── cmd/
│   └── server.go          # Entry point da aplicação
├── internal/
│   ├── config/            # Gerenciamento de configurações
│   ├── storage/           # Strategy Pattern (Redis)
│   ├── limiter/           # Lógica de rate limiting
│   └── middleware/        # Middleware HTTP
├── docker-compose.yml     # Orquestração de containers
├── Dockerfile             # Build da aplicação
└── .env                   # Configurações
```

### Componentes Principais

#### 1. Storage (Strategy Pattern)
Interface que define o contrato para persistência, atualmente implementada com **Redis**:
- **RedisStorage**: Implementação principal usando Redis

> **Nota**: Caso necessário, é possível criar outras implementações (PostgreSQL, MongoDB, etc.) sem alterar a lógica de negócio, apenas implementando a interface `Storage`.

#### 2. Rate Limiter
Lógica de negócio que:
- Verifica se IP/Token está bloqueado
- Controla contadores de requisições via Redis
- Bloqueia quando limite é excedido
- Prioriza token sobre IP

#### 3. Middleware
Intercepta requisições HTTP e aplica rate limiting antes de chegar nos handlers.

## ⚙️ Configuração

### Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto (ou use variáveis de ambiente do sistema):

```env
# Limites de requisições por segundo
RATE_LIMIT_IP=5           # Limite por IP
RATE_LIMIT_TOKEN=100      # Limite por Token

# Tempo de bloqueio em segundos
BLOCK_DURATION=300        # 5 minutos

# Redis
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0

# Servidor
SERVER_PORT=8080
```

### Hierarquia de Configuração

1. **Variáveis de ambiente do sistema** (maior prioridade)
2. **Arquivo `.env`**
3. **Valores padrão** (menor prioridade)

## 🚀 Como Executar

### Usando Docker Compose (Recomendado)

```bash
# Inicia Redis + Aplicação
docker compose up -d

# Verifica logs
docker compose logs -f app

# Para os containers
docker compose down
```

A aplicação estará disponível em `http://localhost:8080`

### Executando Localmente

```bash
# Instala dependências
go mod download

# Inicia Redis (em outro terminal)
docker run -p 6379:6379 redis:7-alpine

# Executa a aplicação
go run cmd/server.go
```

## 📡 Endpoint

### `GET /`
Endpoint único de teste que retorna status do rate limiter.

**Resposta quando permitido:**
```json
{
  "message": "Rate limiter is working!",
  "status": "ok"
}
```

**Exemplos:**
```bash
# Requisição sem token (limite: 5 req/s)
curl http://localhost:8080/

# Requisição com token (limite: 100 req/s)
curl -H "API_KEY: meu-token-123" http://localhost:8080/
```

## 🔑 Usando Tokens (API Key)

Envie o token no header `API_KEY`:

```bash
# Requisição com token (limite: 100 req/s)
curl -H "API_KEY: abc123" http://localhost:8080/

# Requisição sem token (limite: 5 req/s)
curl http://localhost:8080/
```

## 🎯 Como Funciona

### Fluxo de Requisição

```
Cliente faz requisição
    ↓
Middleware extrai IP e Token
    ↓
Rate Limiter verifica:
  - Token existe? → Usa limite do token
  - Não → Usa limite do IP
    ↓
Redis verifica:
  1. Está bloqueado?
  2. Contador de requisições na janela
  3. Excedeu o limite?
    ↓
    ├─ SIM → Retorna HTTP 429
    └─ NÃO → Permite requisição (200)
```

### Bloqueio

Quando o limite é excedido:
- O IP ou Token é bloqueado por `BLOCK_DURATION` segundos
- Todas as requisições retornam **HTTP 429** durante o bloqueio
- Após expirar, o contador é resetado

### Exemplo Prático

**Configuração:** `RATE_LIMIT_IP=5`, `BLOCK_DURATION=300`

```bash
# Requisições 1-5: ✅ Permitidas
curl http://localhost:8080/  # 200 OK
curl http://localhost:8080/  # 200 OK
curl http://localhost:8080/  # 200 OK
curl http://localhost:8080/  # 200 OK
curl http://localhost:8080/  # 200 OK

# Requisição 6: ❌ Bloqueada (ativa bloqueio de 5 minutos)
curl http://localhost:8080/  # 429 Too Many Requests

# Próximos 5 minutos: todas as requisições retornam 429
curl http://localhost:8080/  # 429 Too Many Requests

# Após 5 minutos: volta ao normal
curl http://localhost:8080/  # 200 OK
```

## 🧪 Testes

### Executar Testes Unitários

```bash
# Todos os testes
go test ./... -v

# Apenas limiter
go test ./internal/limiter -v

# Apenas middleware
go test ./internal/middleware -v

# Com cobertura
go test ./... -cover
```

> **Nota**: Os testes usam mocks do storage, não requerem Redis rodando.

### Testes de Carga

Use ferramentas como Apache Bench, wrk, ou stress-test personalizado:

```bash
# Exemplo com Apache Bench
ab -n 100 -c 10 http://localhost:8080/

# 100 requisições, 10 concorrentes
# Deve bloquear após exceder o limite
```

## �� Extensibilidade (Strategy Pattern)

Embora o projeto use **Redis** como storage principal, a arquitetura permite trocar facilmente o mecanismo de persistência:

### Como adicionar outro Storage

1. **Implemente a interface Storage:**

```go
type Storage interface {
    Allow(key string, limit int, window time.Duration) (bool, error)
    Block(key string, duration time.Duration) error
    IsBlocked(key string) (bool, error)
}
```

2. **Crie sua implementação:**

```go
// Exemplo: PostgreSQL
type PostgresStorage struct {
    db *sql.DB
}

func (p *PostgresStorage) Allow(...) (bool, error) {
    // Implementação com Postgres
}
```

3. **Use na aplicação:**

```go
// Em cmd/server.go
postgresStorage := storage.NewPostgres(...)
rateLimiter := limiter.New(postgresStorage, cfg.RateLimitIP, ...)
```

## 📊 Estrutura Redis

O Redis armazena as seguintes chaves:

```
counter:ip:192.168.1.1        → Contador de requisições (TTL: 1s)
counter:token:abc123          → Contador de requisições (TTL: 1s)
blocked:ip:192.168.1.1        → Marca bloqueio (TTL: blockDuration)
blocked:token:abc123          → Marca bloqueio (TTL: blockDuration)
```

### Comandos Redis Utilizados

- **INCR**: Incrementa contador de requisições
- **EXPIRE**: Define TTL das chaves
- **SET com TTL**: Cria chaves de bloqueio com expiração
- **EXISTS**: Verifica se chave de bloqueio existe

## 🛠️ Tecnologias Utilizadas

- **Go 1.23**: Linguagem de programação
- **Fiber v3**: Framework web
- **Redis 7**: Armazenamento de contadores e bloqueios
- **Viper**: Gerenciamento de configurações
- **Docker**: Containerização

## 📝 Resposta HTTP 429

Quando o limite é excedido, a API retorna:

```json
{
  "error": "you have reached the maximum number of requests or actions allowed within a certain time frame"
}
```

## 🐛 Troubleshooting

### Redis não conecta
```bash
# Verifica se Redis está rodando
docker ps | grep redis

# Testa conexão
redis-cli -h localhost -p 6379 ping

# Verifica logs do Redis
docker-compose logs redis
```

### Porta 8080 em uso
```bash
# Muda a porta no .env
SERVER_PORT=8081
```

### Imports não encontrados
```bash
# Atualiza dependências
go mod tidy
go mod download
```

### Rate limiter não está bloqueando
```bash
# Verifica configurações no .env
cat .env

# Verifica se Redis está funcionando
redis-cli -h localhost -p 6379 KEYS "*"
```

## 📄 Licença

Este projeto foi desenvolvido como desafio do curso Full Cycle.
