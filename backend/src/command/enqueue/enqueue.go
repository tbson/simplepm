package main

import (
	"src/client/rabbitclient"
	"src/common/ctype"
	"src/module/queue"
	"src/util/localeutil"
)

func main() {
	localeutil.Init("en")
	client := rabbitclient.NewClient()

	auditPayload := ctype.Dict{
		"user_id": 1,
		"action":  "AUDIT LOG 1!",
	}

	broadcastPayload := ctype.Dict{
		"channel": "channel1",
		"message": "BROADCAST LOG 1!",
	}

	client.Publish(queue.AUDIT_LOG_QUEUE, auditPayload)
	client.Publish(queue.BROADCAST_MESSAGE_QUEUE, broadcastPayload)
}
