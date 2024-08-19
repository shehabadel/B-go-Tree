package main

import (
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	wait := make(chan string)
	conn, err := amqp091.Dial("amqp://localhost:5672")
	if err != nil {
		log.Default().Fatalln("connection err", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Default().Fatal("ch err", err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare("ft.customer", amqp091.ExchangeDirect, true, false, false, false, nil)

	if err != nil {
		log.Default().Fatalln("exchange error", err)
	}

	body := []byte("hello world")

	err = ch.Publish("ft.customer", "ft.egbank", false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        body,
	})
	if err != nil {
		log.Default().Println("err publish: ", err)
	}

	go consume(ch, wait)

	for {
		select {
		case res := <-wait:
			fmt.Println(res)
			close(wait)
			return
		}
	}
}

func consume(ch *amqp091.Channel, wait chan string) {

	q, err := ch.QueueDeclare("ft.customer.tez", true, false, false, false, nil)
	if err != nil {
		log.Default().Println("err create queue ", err)
	}

	err = ch.QueueBind(q.Name, "ft.egbank", "ft.customer", false, nil)
	if err != nil {
		log.Default().Println("err bind queue ", err)
	}

	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Default().Println("error consume ", err)
	}
	for msg := range msgs {
		fmt.Println("Consumed the following: ", string(msg.Body))
		err := msg.Ack(false)
		if err != nil {
			log.Default().Println("ack error")
		}
		wait <- "done"
	}
}
