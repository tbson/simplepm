package main

import (
	queueadapter "src/adapter/queue"
	"src/common/ctype"
	"src/worker/queueworker"
)

func main() {
	client := queueadapter.New()

	testPayload := ctype.Dict{
		"user_id": 1,
		"action":  "ACTION_1",
	}

	client.Publish(queueworker.TEST_MESSAGE, testPayload)
}
