package main

import (
	"log"

	queueadapter "src/adapter/queue"
	"src/queue"
	"src/util/dbutil"
)

func main() {
	log.Println("[+] Waiting for messages...")
	dbutil.InitDb()
	client := queueadapter.New()
	client.Consumes(queue.Handlers)
}
