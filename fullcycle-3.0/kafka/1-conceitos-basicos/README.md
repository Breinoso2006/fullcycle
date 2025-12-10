# Kafka - Comandos Básicos

## Subindo o ambiente

```bash
docker compose up -d
```

## Acessando o container

Para facilitar o trabalho, você pode entrar no container do Kafka:

```bash
docker exec -it 1-conceitos-basicos-kafka-1 bash
```

Uma vez dentro do container, todos os comandos Kafka estarão disponíveis diretamente.

## Comandos principais

> **Nota:** Os comandos abaixo assumem que você está **dentro do container**. Se preferir executar de fora, adicione `docker exec -it 1-conceitos-basicos-kafka-1` antes de cada comando.

### Criar um tópico

```bash
kafka-topics --create --topic=test --bootstrap-server=localhost:9092 --partitions=3
```

### Listar tópicos

```bash
kafka-topics --list --bootstrap-server=localhost:9092
```

### Descrever um tópico

```bash
kafka-topics --describe --topic=test --bootstrap-server=localhost:9092
```

### Produzir mensagens

```bash
kafka-console-producer --topic=test --bootstrap-server=localhost:9092
```

Após executar o comando, digite as mensagens e pressione Enter. Use Ctrl+C para sair.

### Consumir mensagens

Consumir do último offset:
```bash
kafka-console-consumer --topic=test --bootstrap-server=localhost:9092
```

Consumir desde o início:
```bash
kafka-console-consumer --topic=test --bootstrap-server=localhost:9092 --from-beginning
```

### Consumir com Consumer Group

```bash
kafka-console-consumer --topic=test --bootstrap-server=localhost:9092 --group=my-group
```

### Listar Consumer Groups

```bash
kafka-consumer-groups --list --bootstrap-server=localhost:9092
```

### Descrever Consumer Group

```bash
kafka-consumer-groups --describe --group=my-group --bootstrap-server=localhost:9092
```

### Deletar um tópico

```bash
kafka-topics --delete --topic=test --bootstrap-server=localhost:9092
```

## Control Center

Interface web disponível em: http://localhost:9021

## Parando o ambiente

```bash
docker compose down
```
