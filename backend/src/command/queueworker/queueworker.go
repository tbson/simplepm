package main

import (
	"log"

	"src/client/rabbitclient"
	"src/module/queue"
	"src/util/localeutil"
	"src/util/vldtutil"
)

type AuditLog struct {
	UserID uint   `json:"user_id"`
	Action string `json:"action"`
}

type BroadcastMessage struct {
	Channel string `json:"channel"`
	Message string `json:"message"`
}

// Handler for audit log messages.
func handleAuditLog(msg []byte) {
	data, err := vldtutil.BytesToStruct(msg, &AuditLog{})
	if err != nil {
		return
	}
	log.Printf("[AuditLog] User ID: %d, Action: %s", data.UserID, data.Action)
}

// Handler for broadcast messages.
func handleBroadcast(msg []byte) {
	data, err := vldtutil.BytesToStruct(msg, &BroadcastMessage{})
	if err != nil {
		return
	}
	log.Printf("[Broadcast] Channel: %s, Message: %s", data.Channel, data.Message)
}

func main() {
	localeutil.Init("en")
	client := rabbitclient.NewClient()

	queues := map[string]func([]byte){
		queue.AUDIT_LOG_QUEUE:         handleAuditLog,
		queue.BROADCAST_MESSAGE_QUEUE: handleBroadcast,
	}

	log.Println("[+] Waiting for messages...")
	client.Consumes(queues)
}
