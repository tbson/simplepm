package scylla

import (
	"fmt"
	"sync"

	"src/common/setting"

	"github.com/gocql/gocql"
)

type ScyllaProvider interface {
	GenerateID() gocql.UUID
	Query(cql string, args ...interface{}) ([]map[string]interface{}, error)
	QueryWithPaging(cql string, pageSize int, pageState []byte, args ...interface{}) *gocql.Query
	Exec(cql string, args ...interface{}) error
	Close()
}

type client struct {
	session *gocql.Session
}

var (
	once     sync.Once
	instance *client
	initErr  error
)

func newInstance() (*client, error) {
	host := setting.NOSQL_HOST()
	port := setting.NOSQL_PORT()

	url := fmt.Sprintf("%s:%s", host, port)
	cluster := gocql.NewCluster(url)

	// Optional: Set default consistency
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	return &client{
		session: session,
	}, nil
}

func New() *client {
	once.Do(func() {
		instance, initErr = newInstance()
		if initErr != nil {
			// In production, you might want to handle this differently,
			// but panic is common in init if the DB connection cannot be established.
			panic(fmt.Sprintf("Failed to create scylla client: %v", initErr))
		}
	})
	return instance
}

func (c *client) GenerateID() gocql.UUID {
	return gocql.TimeUUID()
}

// Query executes a SELECT (or similar) and returns rows in []map[string]interface{} form.
func (c *client) Query(cql string, args ...interface{}) ([]map[string]interface{}, error) {
	iter := c.session.Query(cql, args...).Iter()
	rows, err := iter.SliceMap()
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	return rows, nil
}

func (c *client) QueryWithPaging(
	cql string,
	pageSize int,
	pageState []byte,
	args ...interface{},
) *gocql.Query {
	q := c.session.Query(cql, args...).PageSize(pageSize)
	if len(pageState) > 0 {
		q = q.PageState(pageState)
	}
	return q
}

// Exec runs an INSERT, UPDATE, DELETE, or any CQL statement that does not return rows.
func (c *client) Exec(cql string, args ...interface{}) error {
	if err := c.session.Query(cql, args...).Exec(); err != nil {
		return fmt.Errorf("exec error: %w", err)
	}
	return nil
}

// Close terminates the session and frees up resources.
// Call this in main() before your application exits.
func (c *client) Close() {
	if c.session != nil {
		c.session.Close()
	}
}
