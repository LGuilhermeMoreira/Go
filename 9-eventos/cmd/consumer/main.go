package main

import (
	"eventos/pkg/rabbitmq"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	ch := rabbitmq.OpenChannel()
	defer ch.Close()

	msgs := make(chan amqp091.Delivery)

	go rabbitmq.Consume(ch, msgs)

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false) // indicando que a mensagem ja foi lida e que não precisa ser posta na fila
	}
}
