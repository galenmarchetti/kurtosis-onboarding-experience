package cassandra_service

import (
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
	return true
}