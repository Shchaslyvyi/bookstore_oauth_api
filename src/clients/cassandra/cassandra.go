package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	cluster *gocql.ClusterConfig
)

func init() {
	// Connect to Cassandra cluster:
	cluster = gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
}

// GetSession func returns a session
func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}
