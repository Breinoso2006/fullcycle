# Sistema de Leilões - Fechamento Automático

API REST para gerenciamento de leilões com **fechamento automático** utilizando goroutines.

## �� Funcionalidade Implementada

**Fechamento automático de leilões** baseado em tempo configurável via variável de ambiente, utilizando goroutines e controle de concorrência.

## 🚀 Como Executar

### Pré-requisitos
- Docker e Docker Compose instalados

### Executar o projeto

```bash
# 1. Subir os containers
docker compose up -d

# 2. Aplicação rodando em http://localhost:8080

# 3. Para parar
docker compose down

# 4. Para limpar dados do banco
docker compose down -v
```

## ⚙️ Configuração

Arquivo `cmd/auction/.env`:

```env
AUCTION_INTERVAL=20s          # Tempo de duração do leilão (fechamento automático)
BATCH_INSERT_INTERVAL=20s     # Intervalo de processamento de lances
MAX_BATCH_SIZE=4              # Tamanho do lote de lances
```

## 🧪 Executar Testes

```bash
# Subir o MongoDB
docker compose up -d mongodb

# Executar testes
MONGODB_URL="mongodb://admin:admin@localhost:27017/auctions?authSource=admin" \
  go test ./internal/infra/database/auction/... -v

# Resultado esperado: 4 testes passando (incluindo teste de fechamento automático)
```

## 📡 Endpoints Principais

### Criar leilão
```bash
curl -X POST http://localhost:8080/auction \
  -H "Content-Type: application/json" \
  -d '{
    "product_name": "Notebook Dell",
    "category": "Electronics",
    "description": "Notebook em excelente estado",
    "condition": 1
  }'
```

### Listar leilões
```bash
# Ativos (status=0)
curl "http://localhost:8080/auction?status=0"

# Fechados (status=1)
curl "http://localhost:8080/auction?status=1"
```

### Criar usuários (necessário para lances)
```bash
docker exec -i mongodb mongosh -u admin -p admin --authenticationDatabase admin --eval '
use auctions
db.users.insertMany([
  {"_id": "6ba7b810-9dad-11d1-80b4-00c04fd430c8", "name": "João Silva"},
  {"_id": "7ba7b810-9dad-11d1-80b4-00c04fd430c9", "name": "Maria Santos"}
])
'
```

### Criar lance
```bash
curl -X POST http://localhost:8080/bid \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
    "auction_id": "SEU_AUCTION_ID_AQUI",
    "amount": 1500.00
  }'
```

**Nota:** Lances são processados em batch. Aguarde ~20s para visualizar.

### Listar lances
```bash
curl "http://localhost:8080/bid/SEU_AUCTION_ID_AQUI"
```

## 🔄 Como Funciona o Fechamento Automático

1. **Goroutine em background** monitora leilões continuamente
2. A cada `AUCTION_INTERVAL`, verifica leilões expirados
3. Atualiza status para `Completed` automaticamente
4. **Thread-safe** com uso de `sync.Mutex`

**Implementação:** `internal/infra/database/auction/create_auction.go`

**Testes:** `internal/infra/database/auction/create_auction_test.go`

## 🎬 Exemplo Completo

```bash
# 1. Subir ambiente
docker compose up -d

# 2. Criar usuários
docker exec -i mongodb mongosh -u admin -p admin --authenticationDatabase admin --eval '
use auctions
db.users.insertMany([{"_id": "6ba7b810-9dad-11d1-80b4-00c04fd430c8", "name": "João"}])
'

# 3. Criar leilão
curl -X POST http://localhost:8080/auction \
  -H "Content-Type: application/json" \
  -d '{"product_name": "Notebook", "category": "Tech", "description": "Novo", "condition": 1}'

# 4. Listar e copiar o ID
curl "http://localhost:8080/auction?status=0"

# 5. Criar lance
curl -X POST http://localhost:8080/bid \
  -H "Content-Type: application/json" \
  -d '{"user_id": "6ba7b810-9dad-11d1-80b4-00c04fd430c8", "auction_id": "ID_AQUI", "amount": 1000}'

# 6. Aguardar AUCTION_INTERVAL (20s) e verificar fechamento
sleep 25
curl "http://localhost:8080/auction?status=1"
```

## 📂 Estrutura

```
├── cmd/auction/                    # Entry point
├── internal/
│   ├── entity/                    # Entidades
│   ├── infra/database/auction/    # ⭐ Implementação do fechamento automático
│   │   ├── create_auction.go      # Goroutine + controle de concorrência
│   │   └── create_auction_test.go # Testes automatizados
│   ├── infra/api/                 # Controllers HTTP
│   └── usecase/                   # Casos de uso
├── docker-compose.yml
└── Dockerfile
```

## 📝 Tecnologias

- Go 1.20+
- MongoDB
- Gin Framework
- Docker & Docker Compose

---

**Desenvolvido como parte do curso Full Cycle - Go Expert**
