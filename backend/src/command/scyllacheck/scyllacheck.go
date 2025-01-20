package main

import (
	"fmt"

	"github.com/gocql/gocql"
)

func main() {
	var cluster = gocql.NewCluster("simplepm_nosql:9042")

	var session, err = cluster.CreateSession()
	if err != nil {
		panic("Failed to connect to cluster")
	}

	defer session.Close()

	var query = session.Query("SELECT * FROM event.messages")

	if rows, err := query.Iter().SliceMap(); err == nil {
		for _, row := range rows {
			fmt.Printf("%v\n", row)
		}
	} else {
		panic("Query error: " + err.Error())
	}
}
