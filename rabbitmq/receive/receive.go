package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "faild to connect to RabbitMq")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Faild to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", false, false, false, false, nil,
	)
	failOnError(err, "Faild to declare a queue")

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)

	failOnError(err, "Faild to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	log.Printf("[*] waiting for messages. To exit press CTRL + C ")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
