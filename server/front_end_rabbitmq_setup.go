package server

import (
	"fmt"
	"front-end-httpserver/config"
	"github.com/streadway/amqp"
	"log"
)

func InitRabbitMQConnection() (*amqp.Connection, error) {
	// Creating connection for Rabbit MQ
	rabbitMQUrl := fmt.Sprintf("amqp://guest:guest@%s:%s/", config.DefaultIP, config.DefaultRabbitMQPort)
	conn, err := amqp.Dial(rabbitMQUrl)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect tp rabbitmq", err)
		return nil, err
	}
	return conn, nil
}

func InitRabbitMQChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	// Creating channel for Rabbit MQ
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "Failed to open a channel", err)
		return nil, err
	}
	// Declaring Exchange
	err = ch.ExchangeDeclare(
		config.ExchangeName, // name
		"topic",             // type
		true,                // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare an exchange", err)
	}
	// Creating the Queue
	queue, err := ch.QueueDeclare(
		config.QueueName, // queue name
		false,            // durable
		false,            // delete when unused
		true,             // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a queue", err)
	}
	// Setting the binding key
	err = ch.QueueBind(
		queue.Name,                    // queue name
		"*.*."+config.SenderName+".*", // routing key
		config.ExchangeName,           // exchange
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to bind a queue", err)
	}
	return ch, err
}

func (rabbitMQ *FrontEndServer) CloseRabbitmqConnection() {
	err := rabbitMQ.amqpChannel.Close()
	err = rabbitMQ.amqpConnection.Close()
	if err != nil {
		log.Fatalf("%s: %s", "Failed to close connection", err)
	}
}

/**
Function to publish a message to the exchange
*/
//func publishMessage(routingKey string, body []byte, ch *amqp.Channel) {
//	err := ch.Publish(
//		exchangeName, // exchange
//		routingKey,   // routing key
//		false,        // mandatory
//		false,        // immediate
//		amqp.Publishing{
//			ContentType: "text/plain",
//			Body:        body,
//		})
//	failOnError(err, "Failed to publish message: "+routingKey)
//	log.Printf("Sent event to %s with message: %s \n", routingKey, string(body))
//}
//

//
