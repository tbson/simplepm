package main

import (
	"src/client/queueclient"
	"src/common/ctype"
	"src/queue"
)

func main() {
	client := queueclient.NewClient()

	testPayload := ctype.Dict{
		"user_id": 1,
		"action":  "ACTION_1",
	}

	client.Publish(queue.TEST_MESSAGE, testPayload)
}
