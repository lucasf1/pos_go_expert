# Desafio 3 - Go Expert

This project is a Go-based application that exposes a set of services for managing orders. It provides three different APIs for interacting with the services: a RESTful API, a gRPC API, and a GraphQL API. The project follows a clean architecture pattern, with a clear separation of concerns between the different layers of the application.

## Main Technologies

*   **Go:** The primary programming language used in the project.
*   **MySQL:** The database used to store the orders.
*   **RabbitMQ:** The message broker used to handle asynchronous events.
*   **GraphQL:** A query language for APIs.
*   **gRPC:** A high-performance, open-source universal RPC framework.
*   **Docker:** Used to run the application and its dependencies in a containerized environment.

## Architecture

The project is divided into the following main components:

*   **`cmd`:** Contains the main entry point of the application.
*   **`configs`:** Contains the configuration files for the application.
*   **`internal`:** Contains the core business logic of the application, including entities, use cases, and repositories.
*   **`pkg`:** Contains reusable packages that can be shared across different projects.

## Building and Running

To build and run the project, you will need to have Docker and Docker Compose installed on your machine. Once you have them installed, you can follow these steps:

1.  **Start the services:**
    ```bash
    docker-compose up -d
    ```
2.  **Run the application:**
    ```bash
    go run cmd/main.go
    ```

The application will be available at the following endpoints:

*   **RESTful API:** `http://localhost:8080`
*   **gRPC API:** `localhost:50051`
*   **GraphQL API:** `http://localhost:8081`

## Interacting with the APIs

### RESTful API

You can use the following `curl` commands to interact with the RESTful API:

**Create an order:**

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "id": "123",
    "price": 10.0,
    "tax": 1.0
}' http://localhost:8080/orders
```

**List all orders:**

```bash
curl http://localhost:8080/orders
```

You can also use the `api/create_order.http` file to interact with the API using a REST client like the VS Code REST Client extension.

### GraphQL API

To interact with the GraphQL API, you can use the GraphQL Playground, which is available at `http://localhost:8081`.

**Create an order:**

```graphql
mutation {
  createOrder(input: {
    id: "123",
    price: 10.0,
    tax: 1.0
  }) {
    id
    price
    tax
    finalPrice
  }
}
```

**List all orders:**

```graphql
query {
  listOrders {
    id
    price
    tax
    finalPrice
  }
}
```

### gRPC API

To interact with the gRPC API, you can use a gRPC client like `grpcurl`.

**Create an order:**

```bash
grpcurl -d '{
    "id": "123",
    "price": 10.0,
    "tax": 1.0
}' -plaintext localhost:50051 pb.OrderService/CreateOrder
```

**List all orders:**

```bash
grpcurl -plaintext localhost:50051 pb.OrderService/ListOrders
```
