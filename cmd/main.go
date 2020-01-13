package main

import (
	"github.com/andreposman/zipode-receiver/pkg/amqp"
)

func main() {
	amqp.ReceiveMessage()
}
