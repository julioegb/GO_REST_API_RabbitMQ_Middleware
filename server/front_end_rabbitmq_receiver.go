package server

import (
	"front-end/config"
	"log"
)

func (rabbitmq *FrontEndServer) StartConsume() error {
	msgs, err := rabbitmq.amqpChannel.Consume(
		config.QueueName,  // queue
		config.SenderName, // consumer
		true,              // auto ack
		false,             // exclusive
		false,             // no local
		false,             // no wait
		nil,               // args
	)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to register a consumer", err)
		return err
	}
	log.Println("Start consuming the Queue...")
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			// Implement here the switch caso to decode the receive messages

			log.Printf("Message Received: %v", msg.Body)
		}
	}()
	<-forever
	return nil
}
