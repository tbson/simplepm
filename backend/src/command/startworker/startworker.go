package main

import (
	"log"

	"src/client/queueclient"
	"src/queue"
	"src/util/dbutil"
)

func main() {
	log.Println("[+] Waiting for messages...")
	dbutil.InitDb()
	client := queueclient.NewClient()
	client.Consumes(queue.Handlers)
}
