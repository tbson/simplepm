package rabbitclient

import (
	"encoding/json"
	"fmt"
	"log"
	"src/common/ctype"
	"src/common/setting"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitClient struct {
	Conn *amqp.Connection
}

var conn *amqp.Connection
var err error
var client = RabbitClient{
	Conn: conn,
}

// InitClient connects to RabbitMQ and returns a RabbitClient instance.
func NewClient() RabbitClient {
	if client.Conn != nil {
		return client
	}
	dialUrl := fmt.Sprintf(
		"amqp://%s:%s@%s:%d/",
		setting.RABBITMQ_USER,
		setting.RABBITMQ_PASSWORD,
		setting.RABBITMQ_HOST,
		setting.RABBITMQ_PORT,
	)
	conn, err = amqp.Dial(dialUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
	client.Conn = conn
	return client
}

// Publish opens a channel, declares the queue, serializes the body to JSON, and publishes it.
func (rc RabbitClient) Publish(queueName string, body ctype.Dict) {
	ch, err := rc.Conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q := getQueue(ch, queueName)

	// Serialize the message body to JSON
	jsonBody, err := json.Marshal(body)
	failOnError(err, "Failed to marshal message body to JSON")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonBody,
		})
	failOnError(err, "Failed to publish a message")
}

// Consume opens a new channel, declares the queue, and returns a channel of deliveries.
func (rc RabbitClient) Consume(queueName string) <-chan amqp.Delivery {
	ch, err := rc.Conn.Channel()
	failOnError(err, "Failed to open a channel for consuming")

	q := getQueue(ch, queueName)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer tag
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // arguments
	)
	failOnError(err, "Failed to register a consumer")
	return msgs
}

func (rc RabbitClient) Consumes(queues map[string]func([]byte)) {
	// For each queue, start a consumer in a separate goroutine.
	for queueName, handler := range queues {
		msgs := client.Consume(queueName)
		go func(q string, msgs <-chan amqp.Delivery, handler func([]byte)) {
			for d := range msgs {
				handler(d.Body)
			}
		}(queueName, msgs, handler)
	}
	// Block forever.
	select {}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func getQueue(ch *amqp.Channel, queueName string) amqp.Queue {
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")
	return q
}
