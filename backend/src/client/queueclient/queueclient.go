package queueclient

import (
	"fmt"
	"src/client/rabbitclient"
	"src/common/ctype"
	"src/common/setting"
)

type Client interface {
	Publish(queueName string, body ctype.Dict)
	Consumes(queues map[string]func([]byte))
}

func NewClient() Client {
	switch setting.QUEUE_BACKEND() {
	case "rabbitmq":
		return rabbitclient.NewClient()
	default:
		panic(fmt.Sprintf("unsupported queue backend: %s", setting.QUEUE_BACKEND()))
	}
}
