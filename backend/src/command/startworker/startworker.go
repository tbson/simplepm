package main

import (
	"log"

	"src/client/queueclient"
	"src/queue"
)

func main() {
	log.Println("[+] Waiting for messages...")
	client := queueclient.NewClient()
	client.Consumes(queue.Handlers)
}
