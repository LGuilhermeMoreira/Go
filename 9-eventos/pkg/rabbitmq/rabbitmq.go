package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")

	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()

	if err != nil {
		panic(err)
	}
	return ch
}

func Consume(ch *amqp.Channel, out chan<- amqp.Delivery) error {
	msgs, err := ch.Consume(
		"minha-fila",
		"consumer",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	for msg := range msgs {
		out <- msg
	}

	return nil
}

func Produce(ch *amqp.Channel, body string) error {
	err := ch.Publish(
		"amq.direct", // exchange
		"",           // fila (mas como não vou botar direto na fila deixa em branco)
		false,
		false,
		amqp.Publishing{ //
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	return err
}
