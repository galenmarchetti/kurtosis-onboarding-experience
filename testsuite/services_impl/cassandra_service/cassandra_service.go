package cassandra_service

import (
	"github.com/gocql/gocql"
	"github.com/kurtosis-tech/kurtosis-libs/golang/lib/services"
)


type CassandraService struct {
	serviceCtx *services.ServiceContext
	port       int
}

func NewCassandraService(serviceCtx *services.ServiceContext, port int) *CassandraService {
	return &CassandraService{serviceCtx: serviceCtx, port: port}
}

func (service CassandraService) GetIPAddress() string {
	return service.serviceCtx.GetIPAddress()
}

// ===========================================================================================
//                              Service interface methods
// ===========================================================================================
func (service CassandraService) IsAvailable() bool {
	/*
		NEW USER ONBOARDING:
		- Write logic to verify that Cassandra is available and ready to be tested.

		NOTE: Cassandra containers create a default database on startup, and do not accept connections until
	          the database is initialized. Therefore, a sensible IsAvailable() function would poll the Cassandra
	          native protocol port until the connection is accepted.
	*/
	// Define object used to represent local Cassandra cluster
	cluster := gocql.NewCluster(service.GetIPAddress())
	cluster.Consistency = gocql.One
	cluster.ProtoVersion = 4

	// Define object used to send queries to local Cassandra cluster
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	return true
}