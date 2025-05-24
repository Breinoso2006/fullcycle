package main

import (
	"github.com/breinoso2006/fullcycle/pos-graduacao-go-expert/11-manipulando-eventos/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello World2", "amq.direct")
}