package my_service

import (
	"github.com/kurtosis-tech/kurtosis-libs/golang/lib/services"
)


type MyService struct {
	serviceCtx *services.ServiceContext
}

func NewMyService(serviceCtx *services.ServiceContext) *MyService {
	return &MyService{serviceCtx: serviceCtx}
}

// ===========================================================================================
//                              Service interface methods
// ===========================================================================================
func (service MyService) IsAvailable() bool {
	/*
		NEW USER ONBOARDING:
		- Write logic to verify that Cassandra is available and ready to be tested.

		NOTE: Cassandra containers create a default database on startup, and do not accept connections until
	          the database is initialized. Therefore, a sensible IsAvailable() function would poll the Cassandra
	          native protocol port until the connection is accepted.
	*/
	return true
}