# Observabilidade com OpenTelemetry - Go

Sistema de consulta de temperatura por CEP com rastreamento distribuído usando OpenTelemetry e Zipkin.

## 🚀 Como Executar

### Pré-requisitos
- Docker
- Docker Compose

### ⚙️ Configuração da API Key

**IMPORTANTE:** Antes de executar o projeto, você precisa configurar sua chave da WeatherAPI.

1. **Obtenha uma chave gratuita** em: https://www.weatherapi.com/
   - Crie uma conta gratuita
   - Copie sua API Key

2. **Configure a variável de ambiente:**

   **Opção 1 - Variável de ambiente do sistema (Recomendado):**
   ```bash
   export WEATHER_API_KEY="sua-chave-aqui"
   ```

   **Opção 2 - Arquivo .env (Desenvolvimento local):**
   
   Edite o arquivo `orchestrator/cmd/.env` e adicione sua chave:
   ```env
   WEATHER_API_KEY=sua-chave-aqui
   ```

   **Opção 3 - Docker Compose (Desenvolvimento):**
   
   Edite o arquivo `docker-compose.yml` na linha da variável `WEATHER_API_KEY`:
   ```yaml
   environment:
     - WEATHER_API_KEY=sua-chave-aqui
   ```

### Iniciar todos os serviços

```bash
docker-compose up --build
```

Isso irá:
1. Subir o Zipkin na porta 9411
2. Subir o Orchestrator (Service B) na porta 8081
3. Subir o Server (Service A) na porta 8080

### Testar a aplicação

```bash
# Fazer uma requisição POST com JSON
curl -X POST http://localhost:8080/weather \
  -H "Content-Type: application/json" \
  -d '{"cep": "29902555"}'

# Resposta esperada (200 OK)
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.5
}

# Teste com CEP inválido (422)
curl -X POST http://localhost:8080/weather \
  -H "Content-Type: application/json" \
  -d '{"cep": "123"}'

# Resposta: {"message": "invalid zipcode"}

# Teste com CEP não encontrado (404)
curl -X POST http://localhost:8080/weather \
  -H "Content-Type: application/json" \
  -d '{"cep": "99999999"}'

# Resposta: {"message": "can not find zipcode"}
```

### Visualizar os traces no Zipkin

Acesse: http://localhost:9411

1. Clique em "Run Query" para ver os traces
2. Clique em um trace para ver os detalhes
3. Visualize a timeline com todos os spans:
   - Server: `GET /weather/:cep`
   - Orchestrator: `GET /weather/:cep`
   - ViaCEP: `viacep-api-call`
   - WeatherAPI: `weatherapi-api-call`

## 📊 Arquitetura

```
User → Server (8080) → Orchestrator (8081) → APIs Externas
                ↓              ↓                    ↓
              Zipkin (9411) ←──────────────────────┘
```

## 🛑 Parar os serviços

```bash
docker-compose down
```

## 🧪 Testar diferentes cenários

```bash
# CEP válido
curl -X POST http://localhost:8080/weather \
  -H "Content-Type: application/json" \
  -d '{"cep": "01310100"}'

# CEP inválido (formato)
curl -X POST http://localhost:8080/weather \
  -H "Content-Type: application/json" \
  -d '{"cep": "123"}'

# CEP não encontrado
curl -X POST http://localhost:8080/weather \
  -H "Content-Type: application/json" \
  -d '{"cep": "99999999"}'
```

## 📝 Variáveis de Ambiente

As variáveis podem ser ajustadas no `docker-compose.yml`:

**Server:**
- `PORT`: Porta do servidor (padrão: 8080)
- `ORCHESTRATOR_URL`: URL do orchestrator
- `SERVICE_NAME`: Nome do serviço para OTEL
- `ZIPKIN_ENDPOINT`: Endpoint do Zipkin

**Orchestrator:**
- `PORT`: Porta do servidor (padrão: 8081)
- `WEATHER_API_KEY`: Chave da WeatherAPI
- `SERVICE_NAME`: Nome do serviço para OTEL
- `ZIPKIN_ENDPOINT`: Endpoint do Zipkin
