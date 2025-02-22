package main

import (
	"log"

	"src/client/rabbitclient"
	"src/util/localeutil"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Handler for audit log messages.
func handleAuditLog(msg amqp.Delivery) {
	log.Printf("[AuditLog] Received message: %s", msg.Body)
	// TODO: Process audit log message (e.g. write to database)
}

// Handler for broadcast messages.
func handleBroadcast(msg amqp.Delivery) {
	log.Printf("[Broadcast] Received message: %s", msg.Body)
	// TODO: Process broadcast message (e.g. forward to Centrifugo)
}

func main() {
	// Initialize localization and RabbitMQ client.
	localeutil.Init("en")
	client := rabbitclient.InitClient()

	// Define the queues you want to consume from and their handlers.
	queues := map[string]func(amqp.Delivery){
		rabbitclient.AUDIT_LOG_QUEUE:         handleAuditLog,
		rabbitclient.BROADCAST_MESSAGE_QUEUE: handleBroadcast,
	}

	// For each queue, start a consumer in a separate goroutine.
	for queueName, handler := range queues {
		msgs := client.Consume(queueName)
		go func(q string, msgs <-chan amqp.Delivery, handler func(amqp.Delivery)) {
			for d := range msgs {
				handler(d)
			}
			log.Printf("Consumer for queue %s has closed", q)
		}(queueName, msgs, handler)
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	// Block forever.
	select {}
}
