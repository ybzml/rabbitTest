package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {
	connction, _ := amqp.Dial("amqp://admin:123@39.107.239.138:5672/")
	channel, _ := connction.Channel()
	channel.ExchangeDeclare(
		"topicEx",
		"topic",
		false,
		false,
		false,
		false,
		nil)

	err1 := channel.Publish("topicEx", "topicwwww.ag", false, false, amqp.Publishing{
		Body: []byte("wsk1"),
	})
	if err1 != nil {
		log.Fatal("send msg failed! %v\n", err1)
	}
	err2 := channel.Publish("topicEx", "topicwwww.ns", false, false, amqp.Publishing{
		Body: []byte("nsk2"),
	})
	if err2 != nil {
		log.Fatal("send msg failed! %v\n", err1)
	}
	err3 := channel.Publish("topicEx", "topickkkk.ns", false, false, amqp.Publishing{
		Body: []byte("kkkkkkkkkkkk"),
	})
	if err2 != nil {
		log.Fatal("send msg failed! %v\n", err3)
	}
}
