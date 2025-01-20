package main

import (
	"fmt"

	"src/client/skyllaclient"
)

func main() {
	defer skyllaclient.Close()

	rows, err := skyllaclient.Query("SELECT * FROM event.messages")
	if err != nil {
		panic(err)
	}

	for _, row := range rows {
		fmt.Println(row)
	}
}
