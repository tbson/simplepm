package skyllaclient

import (
	"fmt"
	"sync"

	"src/common/setting"

	"github.com/gocql/gocql"
)

// We'll keep a single session instance
var (
	once    sync.Once
	session *gocql.Session
	initErr error
)

// initSession is called only once (thread-safe) to create the session.
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

// getSession returns the shared session instance.
func getSession() (*gocql.Session, error) {
	once.Do(initSession)
	return session, initErr
}

// Query executes a CQL query and returns the rows in []map[string]interface{} form.
// Example usage: rows, err := repository.Query("SELECT * FROM event.messages")
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

// Close terminates the session.
// Usually called from main() when the application is shutting down.
func Close() {
	if session != nil {
		session.Close()
	}
}
