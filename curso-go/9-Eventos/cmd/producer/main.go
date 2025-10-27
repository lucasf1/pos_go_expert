package main

import "github.com/lucasf1/goexpert/9-Eventos/fcutils/pkg/rabbitmq"

func main() {
	
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello World!", "amq.direct")
}
