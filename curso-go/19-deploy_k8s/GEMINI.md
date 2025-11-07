# Project Overview

This is a simple Go web server that listens on port 8080 and responds with "Hello World". The project is configured to run within a Docker container and is orchestrated using Docker Compose.

## Building and Running

### Local Development

To run the application directly on your local machine:

```bash
go run main.go
```

The server will be accessible at `http://localhost:8080`.

### Docker

To build and run the application using Docker and Docker Compose:

```bash
docker-compose up --build
```

The server will be accessible at `http://localhost:8080`.

## Development Conventions

*   The main application entrypoint is `main.go`.
*   The application is containerized using the `golang:latest` base image as defined in the `Dockerfile`.
*   `docker-compose.yaml` is configured to mount the project directory into the container at `/app`, allowing for live code reloading.
