package main

import (
	"log"

	"src/adapter/queue"
	"src/util/dbutil"
	"src/worker/queueworker"
)

func main() {
	log.Println("[+] Waiting for messages...")
	dbutil.InitDb()
	client := queue.New()
	client.Consumes(queueworker.Handlers)
}
