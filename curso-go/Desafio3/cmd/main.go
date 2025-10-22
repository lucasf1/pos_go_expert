package main

import (
	"database/sql"
	"fmt"

	"github.com/lucasf1/Desafio3/configs"
	"github.com/lucasf1/Desafio3/internal/event/handler"
	"github.com/lucasf1/Desafio3/internal/infra/web/webserver"
	"github.com/lucasf1/Desafio3/pkg/events"
	"github.com/streadway/amqp"

	// driver sqlite
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("sqlite3", "./db_orders.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	// fmt.Println(configs.WebServerPort)

	// createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)

	webserver.AddHandler("/orders", webOrderHandler.Create)
	webserver.AddHandler("/list_orders", webOrderHandler.List)

	fmt.Println("Starting web server on port", configs.WebServerPort)
	webserver.Start()

}

func getRabbitMQChannel() *amqp.Channel {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5674/")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
