package my_service

import (
	"fmt"
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
	).WithUsedPorts(map[string]bool{"8565/tcp": true}).Build()

	return result, nil
}

func (factory MyServiceConfigFactory) GetRunConfig(
		containerIpAddr string,
		generatedFileFilepaths map[string]string,
		staticFileFilepaths map[services.StaticFileID]string) (*services.ContainerRunConfig, error) {
	entrypointCommand := fmt.Sprintf("geth --dev -http --http.api admin,eth,net,rpc --http.addr %v --http.corsdomain '*' --nat extip:%v",
		containerIpAddr,
		containerIpAddr)
	entrypointArgs := []string{
		"/bin/sh",
		"-c",
		entrypointCommand,
	}
	result := services.NewContainerRunConfigBuilder().WithEntrypointOverride(entrypointArgs).Build()
	return result, nil
}
