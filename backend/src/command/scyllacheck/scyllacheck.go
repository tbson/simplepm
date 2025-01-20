package main

import (
	"fmt"

	"src/client/skyllaclient"
)

func main() {
	defer skyllaclient.Close()
	user_id := 1
	rows, err := skyllaclient.Query("SELECT * FROM event.messages WHERE user_id = ?", user_id)
	if err != nil {
		panic(err)
	}

	for _, row := range rows {
		fmt.Println(row)
	}
}
