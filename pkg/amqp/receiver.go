package amqp

import (
	"log"

	"github.com/andreposman/zipcode-receiver/data"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// ReceiveMessage from a rabbitmq queue
func ReceiveMessage() {

	connString := "amqp://guest:guest@localhost:5672/"

	conn, err := amqp.Dial(connString)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	channel, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"zipcodes", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to decalre queue")

	messages, err := channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, "Failed to consume the messages")

	forever := make(chan bool)
	var ZipCode []string

	for msg := range messages {
		log.Printf("\nReceived a message: %s", msg.Body)
		ZipCode = append(ZipCode, string(msg.Body))
		data.SaveToFile(ZipCode)
	}

	log.Printf("\n [*] Waiting for messages. To exit press CTRL + C")
	<-forever
}
