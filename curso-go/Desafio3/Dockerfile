FROM golang:1.24-alpine AS builder
WORKDIR /src

# cache go modules
COPY go.mod go.sum ./
RUN go mod download

# copy project
COPY . .

# build the server from the cmd folder (includes main.go and wire_gen.go)
WORKDIR /src/cmd
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -ldflags="-w -s" -o /server main.go wire_gen.go

FROM alpine:latest
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
COPY --from=builder /server /server
USER appuser
ENTRYPOINT ["/server"]