package cassandra_service

import (
	"github.com/kurtosis-tech/kurtosis-libs/golang/lib/services"
)


type CassandraService struct {
	serviceCtx *services.ServiceContext
}

func NewCassandraService(serviceCtx *services.ServiceContext) *CassandraService {
	return &CassandraService{serviceCtx: serviceCtx}
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