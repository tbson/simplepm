package skyllaclient

import (
	"fmt"
	"sync"

	"src/common/setting"

	"github.com/gocql/gocql"
)

var (
	once     sync.Once
	instance *Client
	initErr  error
)

// Client wraps a gocql.Session and its ClusterConfig.
type Client struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

func GenerateID() gocql.UUID {
	return gocql.TimeUUID()
}

// newClient initializes the cluster and session, returning a new Client.
func newClient() (*Client, error) {
	host := setting.NOSQL_HOST
	port := setting.NOSQL_PORT

	url := fmt.Sprintf("%s:%s", host, port)
	cluster := gocql.NewCluster(url)

	// Optional: Set default consistency
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	return &Client{
		cluster: cluster,
		session: session,
	}, nil
}

// Client provides a single shared Client instance (singleton).
// If you want multiple independent connections, you can create
// your own version of newClient() calls instead.
func NewClient() *Client {
	once.Do(func() {
		instance, initErr = newClient()
		if initErr != nil {
			// In production, you might want to handle this differently,
			// but panic is common in init if the DB connection cannot be established.
			panic(fmt.Sprintf("Failed to create scylla client: %v", initErr))
		}
	})
	return instance
}

// Query executes a SELECT (or similar) and returns rows in []map[string]interface{} form.
func (c *Client) Query(cql string, args ...interface{}) ([]map[string]interface{}, error) {
	iter := c.session.Query(cql, args...).Iter()
	rows, err := iter.SliceMap()
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	return rows, nil
}

// Exec runs an INSERT, UPDATE, DELETE, or any CQL statement that does not return rows.
func (c *Client) Exec(cql string, args ...interface{}) error {
	if err := c.session.Query(cql, args...).Exec(); err != nil {
		return fmt.Errorf("exec error: %w", err)
	}
	return nil
}

// Close terminates the session and frees up resources.
// Call this in main() before your application exits.
func (c *Client) Close() {
	if c.session != nil {
		c.session.Close()
	}
}
