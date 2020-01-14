package main

import (
	"github.com/andreposman/zipcode-receiver/pkg/amqp"
)

func main() {
	amqp.ReceiveMessage()
}
