# Sistema de Consulta de Temperatura por CEP

Sistema desenvolvido em Go que recebe um CEP brasileiro, identifica a cidade através da API ViaCEP e retorna o clima atual (temperatura em graus Celsius, Fahrenheit e Kelvin) usando a API WeatherAPI.

---

## 🎯 Objetivo

Desenvolver um sistema em Go que receba um CEP, identifique a cidade e retorne o clima atual (temperatura em graus celsius, fahrenheit e kelvin). O sistema está publicado no Google Cloud Run.

---

## ✅ Requisitos Implementados

### Funcionalidades
- ✅ Sistema recebe um CEP válido de 8 dígitos
- ✅ Pesquisa do CEP e localização através da API ViaCEP
- ✅ Consulta de temperatura através da API WeatherAPI
- ✅ Conversão e formatação de temperaturas em: Celsius, Fahrenheit e Kelvin
- ✅ Respostas HTTP adequadas para cada cenário

### Cenários de Resposta

#### ✅ Sucesso (200 OK)
```bash
GET /weather/01001000
```
```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.5
}
```

#### ✅ CEP Inválido - Formato Incorreto (422 Unprocessable Entity)
```bash
GET /weather/123
```
```json
{
  "message": "invalid zipcode"
}
```

#### ✅ CEP Não Encontrado (404 Not Found)
```bash
GET /weather/99999999
```
```json
{
  "message": "can not find zipcode"
}
```

---

## 🏗️ Arquitetura

```
cmd/
├── main.go              # Entry point da aplicação

internal/
├── client/              # Clientes HTTP para APIs externas
│   ├── viacep.go       # Cliente ViaCEP
│   └── weatherapi.go   # Cliente WeatherAPI
├── config/              # Configurações com Viper
│   └── config.go
├── handler/             # Handlers HTTP
│   └── cep_handler.go
└── service/             # Lógica de negócio
    ├── viacep.go
    └── wheatherapi.go
```

---

## 🧪 Testes Automatizados

### Testes Implementados
- ✅ `internal/client/weatherapi_test.go` - Testes de conversão de temperatura
- ✅ `internal/handler/cep_handler_test.go` - Testes do handler HTTP
- ✅ `internal/service/viacep_test.go` - Testes de validação de CEP

### Executar Testes
```bash
# Executar todos os testes
make test

# Ou diretamente
go test ./... -v

# Com coverage
make test-coverage
```

### Resultado dos Testes
```
✅ internal/client: PASS - 8 testes
✅ internal/handler: PASS - 11 testes  
✅ internal/service: PASS - 9 testes
```

---

## 🐳 Docker & Docker Compose

### Arquivos Docker
- ✅ `Dockerfile` - Multi-stage build otimizado
- ✅ `docker-compose.yml` - Orquestração de serviços
- ✅ `.dockerignore` - Otimização do build

### Executar com Docker Compose
```bash
# Build e iniciar
make docker-up

# Ou manualmente
docker compose up -d --build

# Ver logs
make docker-logs

# Parar
make docker-down
```

### Testar a API
```bash
# CEP válido
curl http://localhost:8080/weather/01001000

# CEP inválido
curl http://localhost:8080/weather/123

# CEP não encontrado
curl http://localhost:8080/weather/99999999
```

---

## 🌐 APIs Utilizadas

### ViaCEP
- **URL**: https://viacep.com.br/
- **Uso**: Consulta de informações de CEP brasileiro
- **Gratuita**: Sim

### WeatherAPI
- **URL**: https://www.weatherapi.com/
- **Uso**: Consulta de temperatura atual por localização
- **API Key**: Necessária (gratuita)

---

## 🔧 Configuração

### 1. Obter API Key do WeatherAPI
- Acesse: https://www.weatherapi.com/
- Crie uma conta gratuita
- Copie sua API Key

### 2. Configurar Variáveis de Ambiente
```bash
# Copiar exemplo
cp .env.example cmd/.env

# Editar e adicionar sua chave
weatherKey=SUA_CHAVE_AQUI
```

---

## 📐 Conversões de Temperatura

Implementadas conforme especificação:

- **Celsius → Fahrenheit**: `F = C × 1.8 + 32`
- **Celsius → Kelvin**: `K = C + 273`

---

## 🚀 Deploy no Google Cloud Run

### Status do Deploy
- ✅ Deploy realizado no Google Cloud Run (free tier)
- ✅ Endereço ativo e funcional

### URL da Aplicação
```
🌐 https://desafio-cep-567931065104.southamerica-east1.run.app/weather/{cep-aqui}
```

**TODO**: Adicionar URL do deploy aqui após publicação

### Comandos para Deploy
```bash
# Build da imagem
gcloud builds submit --tag gcr.io/SEU_PROJECT_ID/desafio-cep

# Deploy
gcloud run deploy desafio-cep \
  --image gcr.io/SEU_PROJECT_ID/desafio-cep \
  --platform managed \
  --region southamerica-east1 \
  --allow-unauthenticated \
  --set-env-vars weatherKey=SUA_CHAVE_WEATHER_API
```

---

## 📦 Entrega Completa

### ✅ Checklist
- [X] Código-fonte completo da implementação
- [X] Testes automatizados demonstrando o funcionamento
- [X] Docker/Docker Compose para testes da aplicação
- [X] Deploy no Google Cloud Run (free tier)
- [X] Endereço ativo para acesso (adicionar URL acima)

---

## 🛠️ Tecnologias

- **Go 1.23** - Linguagem de programação
- **Viper** - Gerenciamento de configurações
- **net/http** - Servidor HTTP nativo
- **Docker** - Containerização
- **Google Cloud Run** - Plataforma serverless
- **ViaCEP API** - Consulta de CEP
- **WeatherAPI** - Consulta de temperatura

---

## 📄 Comandos Disponíveis (Makefile)

```bash
make help          # Lista todos os comandos
make build         # Compila a aplicação
make run           # Executa localmente
make test          # Executa os testes
make docker-up     # Sobe com Docker Compose
make docker-down   # Para os containers
make docker-logs   # Mostra os logs
make clean         # Remove arquivos gerados
```