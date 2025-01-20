package skyllaclient

import (
	"fmt"
	"sync"

	"src/common/setting"

	"github.com/gocql/gocql"
)

var (
	once    sync.Once
	session *gocql.Session
	initErr error
)

func initSession() {
	host := setting.NOSQL_HOST
	port := setting.NOSQL_PORT

	url := fmt.Sprintf("%s:%s", host, port)
	cluster := gocql.NewCluster(url)

	session, initErr = cluster.CreateSession()
	if initErr != nil {
		fmt.Printf("Failed to create session: %v\n", initErr)
	}
}

func getSession() (*gocql.Session, error) {
	once.Do(initSession)
	return session, initErr
}

func Query(cql string, args ...interface{}) ([]map[string]interface{}, error) {
	s, err := getSession()
	if err != nil {
		return nil, fmt.Errorf("failed to get session: %w", err)
	}

	iter := s.Query(cql, args...).Iter()
	rows, err := iter.SliceMap()
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}

	return rows, nil
}

func Exec(cql string, args ...interface{}) error {
	s, err := getSession()
	if err != nil {
		return fmt.Errorf("failed to get session: %w", err)
	}

	if err := s.Query(cql, args...).Exec(); err != nil {
		return fmt.Errorf("exec error: %w", err)
	}

	return nil
}

func Close() {
	if session != nil {
		session.Close()
	}
}
