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
	return true
}