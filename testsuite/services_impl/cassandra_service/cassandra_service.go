package cassandra_service
/*
	NEW USER ONBOARDING:
	- Rename this package, this file, and the containing directory to reflect the functionality of your custom test.
	- Rename all structs and functions within to reflect your custom service.
*/

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