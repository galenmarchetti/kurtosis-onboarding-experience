package cassandra_service

import (
	"github.com/gocql/gocql"
	"github.com/kurtosis-tech/kurtosis-libs/golang/lib/services"
	"github.com/palantir/stacktrace"
)


type CassandraService struct {
	serviceCtx *services.ServiceContext
}

func NewCassandraService(serviceCtx *services.ServiceContext) *CassandraService {
	return &CassandraService{serviceCtx: serviceCtx}
}

func (service CassandraService) GetIPAddress() string {
	return service.serviceCtx.GetIPAddress()
}

/*
	Creates and returns an open session for the Cassandra service.
	NOTE: This session is not automatically closed. After calling, make sure to call session.Close()
	on the returned object.
*/
func (service CassandraService) CreateSession() (*gocql.Session, error) {
	cluster := gocql.NewCluster(service.GetIPAddress())
	cluster.Consistency = gocql.One
	cluster.ProtoVersion = 4

	// Define object used to send queries to local Cassandra cluster
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, stacktrace.Propagate(err, "Failed to initiate session.")
	}

	return session, nil
}

/*
	All Kurtosis services must implement the IsAvailable() method, which indicates that the service
	has fully started up and is ready to interact with the world.
	Without this method, many tests would fail because they would start querying services that aren't
	yet ready to be queried.
*/

func (service CassandraService) IsAvailable() bool {
	session, err := service.CreateSession()
	if err != nil {
		stacktrace.Propagate(err, "Failed to initiate session on Cassandra service.")
		return false
	}
	defer session.Close()

	return true
}