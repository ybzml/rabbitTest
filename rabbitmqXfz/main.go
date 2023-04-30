package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"sync"
)

var wg sync.WaitGroup

func main() {
	connction, err1 := amqp.Dial("amqp://admin:admin@192.168.192.129:5672/")
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
	wg.Add(2)
	go func() {
		for msg1 := range createAndCusQue(channel, "topicwwww.#") {
			fmt.Printf("ms1:%s\n", msg1.Body)
			msg1.Ack(false)
		}
	}()
	go func() {
		for msg2 := range createAndCusQue(channel, "topickkkk.#") {
			fmt.Printf("ms2:%s\n", msg2.Body)
			msg2.Ack(false)
		}
	}()
	wg.Wait()
}
func createAndCusQue(channel *amqp.Channel, routing_key string) <-chan amqp.Delivery {
	queue, err := channel.QueueDeclare("", false, false, false, false, nil)
	channel.QueueBind(queue.Name, routing_key, "topicEx", false, nil)
	if err != nil {
		log.Fatal("queue  failed! %v\n", err)
	}
	consume, _ := channel.Consume(queue.Name, "", false, false, false, false, nil)
	return consume
}
