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
	fmt.Println("sender topic running")
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"sender3", // name
		"topic",   // type
		true,      // durable
		false,     // auto-deleted
		false,     // internal
		false,     // no-wait
		nil,       // arguments
	)

	message := Message3{
		Name:  "vanhung",
		Age:   10,
		Email: "vanhung@gmail.com",
	}
	emessage, _ := json.Marshal(message)
	err = ch.Publish(
		"sender3",
		"vanhung",
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

type Message3 struct {
	Name  string
	Age   int
	Email string
}
