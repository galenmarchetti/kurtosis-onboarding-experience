package cassandra_service
/*
	NEW USER ONBOARDING:
	- Rename this package, this file, and the containing directory after your custom service.
	- Rename all structs and functions in this file after your custom service.
*/

import (
	"fmt"
	"github.com/kurtosis-tech/kurtosis-libs/golang/lib/services"
)

const (
	testVolumeMountpoint = "/test-volume"

	clusterCommunicationPort = 7000
	nativeProtocolClientPort = 9042
	jmxPort = 7199
)

type CassandraServiceConfigFactory struct {
	image     string
	port	  int
}

func NewCassandraServiceConfigFactory(image string) *CassandraServiceConfigFactory {
	return &CassandraServiceConfigFactory{image: image}
}


func (factory CassandraServiceConfigFactory) GetCreationConfig(containerIpAddr string) (*services.ContainerCreationConfig, error) {
	result := services.NewContainerCreationConfigBuilder(
		factory.image,
		testVolumeMountpoint,
		func(serviceCtx *services.ServiceContext) services.Service { return NewCassandraService(serviceCtx, nativeProtocolClientPort) },
	).WithUsedPorts(map[string]bool{
		/*
			NEW USER ONBOARDING:
			- Add any other ports that your service needs to have open to other services, or to the
		*/
		fmt.Sprintf("%v/tcp", nativeProtocolClientPort): true,
	}).Build()

	return result, nil
}

func (factory CassandraServiceConfigFactory) GetRunConfig(containerIpAddr string, generatedFileFilepaths map[string]string) (*services.ContainerRunConfig, error) {
	result := services.NewContainerRunConfigBuilder().Build()
	return result, nil
}
