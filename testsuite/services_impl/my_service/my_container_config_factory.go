package my_service

import (
	"github.com/kurtosis-tech/kurtosis-client/golang/services"
)

const (
	testVolumeMountpoint = "/test-volume"
)

type MyServiceConfigFactory struct {
	image             string
	existingServiceIP string
}

func NewMyServiceConfigFactory(image string, existingServiceIP string) *MyServiceConfigFactory {
	return &MyServiceConfigFactory{image: image, existingServiceIP: existingServiceIP}
}


func (factory MyServiceConfigFactory) GetCreationConfig(containerIpAddr string) (*services.ContainerCreationConfig, error) {
	result := services.NewContainerCreationConfigBuilder(
		factory.image,
		testVolumeMountpoint,
	).Build()

	return result, nil
}

func (factory MyServiceConfigFactory) GetRunConfig(
		containerIpAddr string,
		generatedFileFilepaths map[string]string,
		staticFileFilepaths map[services.StaticFileID]string) (*services.ContainerRunConfig, error) {
	result := services.NewContainerRunConfigBuilder().Build()
	return result, nil
}
