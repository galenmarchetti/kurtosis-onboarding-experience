package cassandra_service

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
	image string
	cassandraSeed string
}

func NewCassandraServiceConfigFactory(image string, cassandraSeed string) *CassandraServiceConfigFactory {
	return &CassandraServiceConfigFactory{image: image, cassandraSeed: cassandraSeed}
}


func (factory CassandraServiceConfigFactory) GetCreationConfig(containerIpAddr string) (*services.ContainerCreationConfig, error) {
	result := services.NewContainerCreationConfigBuilder(
		factory.image,
		testVolumeMountpoint,
		func(serviceCtx *services.ServiceContext) services.Service { return NewCassandraService(serviceCtx, nativeProtocolClientPort) },
	).WithUsedPorts(map[string]bool{
		/*
			NEW USER ONBOARDING:
			- Add any other ports that your service needs to have open to other services to this "used ports" map
		*/
		fmt.Sprintf("%v/tcp", nativeProtocolClientPort): true,
		fmt.Sprintf("%v/tcp", clusterCommunicationPort): true,
	}).Build()

	return result, nil
}

func (factory CassandraServiceConfigFactory) GetRunConfig(containerIpAddr string, generatedFileFilepaths map[string]string) (*services.ContainerRunConfig, error) {
	result := services.NewContainerRunConfigBuilder().WithEnvironmentVariableOverrides(map[string]string{
		"CASSANDRA_SEEDS": factory.cassandraSeed,
		"CASSANDRA_LISTEN_ADDRESS": containerIpAddr,
	}).Build()
	return result, nil
}
