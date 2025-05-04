package main

import (
	"fmt"

	"src/client/scylla"
)

func main() {
	client := scylla.New()
	defer client.Close()
	user_id := 1
	rows, err := client.Query("SELECT * FROM event.messages WHERE user_id = ?", user_id)
	if err != nil {
		panic(err)
	}

	for _, row := range rows {
		fmt.Println(row)
	}
}
