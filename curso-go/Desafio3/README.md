# Desafio 3 - Go Expert

Este projeto é uma aplicação Go que expõe um conjunto de serviços para gerenciar pedidos. Ele fornece três APIs diferentes para interagir com os serviços: uma API RESTful, uma API gRPC e uma API GraphQL. O projeto segue um padrão de arquitetura limpa, com uma clara separação de responsabilidades entre as diferentes camadas da aplicação.

## Principais Tecnologias

*   **Go:** A linguagem de programação utilizada no projeto.
*   **MySQL:** O banco de dados utilizado para armazenar os pedidos.
*   **RabbitMQ:** O message broker utilizado para lidar com eventos assíncronos.
*   **GraphQL:** Uma linguagem de consulta para APIs.
*   **gRPC:** Um framework RPC universal de alto desempenho e código aberto.
*   **Docker:** Utilizado para executar a aplicação e suas dependências em um ambiente conteinerizado.

## Arquitetura

O projeto está dividido nos seguintes componentes principais:

*   **`cmd`:** Contém o ponto de entrada principal da aplicação.
*   **`configs`:** Contém os arquivos de configuração para a aplicação.
*   **`internal`:** Contém a lógica de negócios principal da aplicação, incluindo entidades, casos de uso e repositórios.
*   **`pkg`:** Contém pacotes reutilizáveis que podem ser compartilhados entre diferentes projetos.

## Construindo e Executando

Para construir e executar o projeto, você precisará ter o Docker e o Docker Compose instalados em sua máquina. Depois de instalados, você pode seguir estas etapas:

1.  **Iniciar os serviços:**
    ```bash
    docker-compose up -d
    ```
2.  **Executar a aplicação:**
    ```bash
    DB_DRIVER=mysql \
    DB_HOST=127.0.0.1 \
    DB_PORT=3306 \
    DB_USER=root \
    DB_PASSWORD=root \
    DB_NAME=orders \
    WEB_SERVER_PORT=8000 \
    GRPC_SERVER_PORT=50051 \
    GRAPHQL_SERVER_PORT=8081 \
    RABBITMQ_URL=amqp://guest:guest@localhost:5672/ \
    go run ./cmd/main.go ./cmd/wire_gen.go
    ```

A aplicação estará disponível nos seguintes endpoints:

*   **API RESTful:** `http://localhost:8000`
*   **API gRPC:** `localhost:50051`
*   **API GraphQL:** `http://localhost:8081`

## Interagindo com as APIs

### API RESTful

Você pode usar os seguintes comandos `curl` para interagir com a API RESTful:

**Criar um pedido:**

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "id": "aaaaa",
    "price": 10.0,
    "tax": 1.0
}' http://localhost:8000/order
```

**Listar todos os pedidos:**

```bash
curl http://localhost:8000/order
```

Você também pode usar o arquivo `api/create_order.http` para interagir com a API usando um cliente REST como a extensão VS Code REST Client.

### API GraphQL

Para interagir com a API GraphQL, você pode usar o GraphQL Playground, que está disponível em `http://localhost:8081`.

**Criar um pedido:**

```graphql
mutation createOrder {
  createOrder(input: {
    id: "ccccc",
    Price: 10.0,
    Tax: 1.0
  }) {
    id
    Price
    Tax
    FinalPrice
  }
}
```

**Listar todos os pedidos:**

```graphql
query queryOrders{
  orders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

### API gRPC

Para interagir com a API gRPC, você pode usar um cliente gRPC como `grpcurl`.

**Criar um pedido:**

```bash
grpcurl -d '{
    "id": "eeeee",
    "price": 10.0,
    "tax": 1.0
}' -plaintext localhost:50051 pb.OrderService/CreateOrder
```

**Listar todos os pedidos:**

```bash
grpcurl -plaintext localhost:50051 pb.OrderService/ListOrders
```
