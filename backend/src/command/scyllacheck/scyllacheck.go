package main

import (
	"fmt"

	"src/common/setting"

	"github.com/gocql/gocql"
)

func main() {
	host := setting.NOSQL_HOST
	port := setting.NOSQL_PORT
	url := fmt.Sprintf("%s:%s", host, port)
	var cluster = gocql.NewCluster(url)

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
