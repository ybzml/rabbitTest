package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {
	connction, _ := amqp.Dial("amqp://admin:123@39.107.239.138:5672/")
	channel, _ := connction.Channel()
	msg := "hello world"
	channel.ExchangeDeclare(
		"derectEx",
		"direct",
		false,
		false,
		false,
		false,
		nil)

	err := channel.Publish("derectEx", "direct_key", false, false, amqp.Publishing{
		Body: []byte(msg),
	})
	if err != nil {
		log.Fatal("send msg failed! %v\n", err)

	}
}
