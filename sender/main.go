package main

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}
	fmt.Println("RabbitMQ running")
	// Let's start by opening a channel to our RabbitMQ instance
	// over the connection we have already established
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	fmt.Println(q)

	if err != nil {
		fmt.Println(err)
	}

	message := Message{
		Name:  "vanhung",
		Age:   10,
		Email: "vanhung@gmail.com",
	}
	emessage, _ := json.Marshal(message)
	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        emessage,
		},
	)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Published Message to Queue")
}

type Message struct {
	Name  string
	Age   int
	Email string
}
