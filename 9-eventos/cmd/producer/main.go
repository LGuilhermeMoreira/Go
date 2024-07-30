package main

import "eventos/pkg/rabbitmq"

func main() {
	ch := rabbitmq.OpenChannel()

	defer ch.Close()

	if err := rabbitmq.Produce(ch, "Ola mundo!"); err != nil {
		panic(err)
	}
}
