package main

import (
	"log"
	"src/client/rabbitclient"
	"src/common/ctype"
	"src/util/localeutil"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	localeutil.Init("en")
	client := rabbitclient.InitClient()
	auditPayload := ctype.Dict{
		"message": "AUDIT LOG 1!",
	}

	broadcastPayload := ctype.Dict{
		"message": "BROADCAST LOG 1!",
	}

	client.Publish(rabbitclient.AUDIT_LOG_QUEUE, auditPayload)
	client.Publish(rabbitclient.BROADCAST_MESSAGE_QUEUE, broadcastPayload)
}
