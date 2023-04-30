package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	connction, err1 := amqp.Dial("amqp://admin:123@39.107.239.138:5672/")
	defer connction.Close()
	if err1 != nil {
		log.Fatal("连接  failed! %v\n", err1)
	}
	channel, _ := connction.Channel()
	defer channel.Close()
	channel.ExchangeDeclare(
		"topicEx",
		"topic",
		false,
		false,
		false,
		false,
		nil)

	queue, err := channel.QueueDeclare("", false, false, false, false, nil)

	channel.QueueBind(queue.Name, "direct_key", "derectEx", false, nil)

	if err != nil {
		log.Fatal("queue  failed! %v\n", err)
	}
	consume, _ := channel.Consume(queue.Name, "", false, false, false, false, nil)
	msg := <-consume
	fmt.Printf("received msg:%s\n", msg.Body)
}
