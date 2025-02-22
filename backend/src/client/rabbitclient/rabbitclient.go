package rabbitclient

import (
	"encoding/json"
	"fmt"
	"log"
	"src/common/ctype"
	"src/common/setting"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

const AUDIT_LOG_QUEUE = "AUDIT_LOG"
const BROADCAST_MESSAGE_QUEUE = "BROADCAST_MESSAGE"

type RabbitClient struct {
	Conn *amqp.Connection
}

var conn *amqp.Connection
var err error
var client = RabbitClient{
	Conn: conn,
}

// InitClient connects to RabbitMQ and returns a RabbitClient instance.
func InitClient() RabbitClient {
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

func GetClient() RabbitClient {
	return client
}

func (rc RabbitClient) GetQueue(ch *amqp.Channel, queueName string) amqp.Queue {
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

// Publish opens a channel, declares the queue, serializes the body to JSON, and publishes it.
func (rc RabbitClient) Publish(queueName string, body ctype.Dict) {
	ch, err := rc.Conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q := rc.GetQueue(ch, queueName)

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
	log.Printf(" [x] Sent %s to queue %s\n", jsonBody, queueName)
}

// Consume opens a new channel, declares the queue, and returns a channel of deliveries.
func (rc RabbitClient) Consume(queueName string) <-chan amqp.Delivery {
	ch, err := rc.Conn.Channel()
	failOnError(err, "Failed to open a channel for consuming")

	q := rc.GetQueue(ch, queueName)

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
