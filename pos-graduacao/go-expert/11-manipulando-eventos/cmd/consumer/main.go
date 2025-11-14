package main

import (
	"fmt"

	"github.com/breinoso2006/fullcycle/pos-graduacao-go-expert/11-manipulando-eventos/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs, "orders")
	for msg := range msgs {
		fmt.Println(string(msg.Body))
		// Como padrão, é melhor chamar o ack manualmente para evitar que a mensagem seja perdida
		msg.Ack(false)
	}
}
