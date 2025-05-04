package queue

import (
	"fmt"
	"src/adapter/queue/rabbit"
	"src/common/ctype"
	"src/common/setting"
)

type Client interface {
	Publish(queueName string, body ctype.Dict)
	Consumes(queues map[string]func([]byte))
}

func New() Client {
	switch setting.QUEUE_BACKEND() {
	case "rabbitmq":
		return rabbit.New()
	default:
		panic(fmt.Sprintf("unsupported queue backend: %s", setting.QUEUE_BACKEND()))
	}
}
