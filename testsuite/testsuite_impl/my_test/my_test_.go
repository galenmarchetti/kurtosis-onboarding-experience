package my_test

import (
	"bytes"
	"fmt"
	"github.com/galenmarchetti/kurtosis-onboarding-test/testsuite/services_impl/my_service"
	"github.com/kurtosis-tech/kurtosis-client/golang/networks"
	"github.com/kurtosis-tech/kurtosis-client/golang/services"
	"github.com/kurtosis-tech/kurtosis-libs/golang/lib/testsuite"
	"github.com/palantir/stacktrace"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	waitForStartupTimeBetweenPolls = 1 * time.Second
	waitForStartupMaxPolls = 90

	rpcPort = 8545
)

var serviceIDs = []services.ServiceID{
	"service-0",
}

type MyTest struct {}

func (test MyTest) Configure(builder *testsuite.TestConfigurationBuilder) {
	builder.WithSetupTimeoutSeconds(360).WithRunTimeoutSeconds(360)
}

func (test MyTest) Setup(networkCtx *networks.NetworkContext) (networks.Network, error) {
	logrus.Infof("Setting up test.")

	// ================== BEGIN ETH CODE ================
	configFactory := my_service.NewMyServiceConfigFactory("ethereum/client-go", "")


	serviceCtx, _, err := networkCtx.AddService(serviceIDs[0], configFactory)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred adding the service")
	}

	// TODO check for availability?????

	logrus.Infof("Added service with IP address: %v", serviceCtx.GetIPAddress())

	return networkCtx, nil
}

func (test MyTest) Run(uncastedNetwork networks.Network) error {
	logrus.Infof("Running test.")
	// Necessary because Go doesn't have generics
	castedNetwork := uncastedNetwork.(*networks.NetworkContext)

	serviceCtx, err := castedNetwork.GetServiceContext(serviceIDs[0])
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred getting the datastore service context")
	}
	logrus.Infof("Got service context for datastore service '%v'", serviceCtx.GetServiceID())

	/*
		NEW USER ONBOARDING:
		- Fill in the logic necessary to run your custom test.
	*/
	return nil
}


// ================== BEGIN ETH CODE ================

func sendRpcCall(ipAddress string, rpcJsonString string, targetStruct interface{}) error {
	url := fmt.Sprintf("http://%v:%v", ipAddress, rpcPort)
	var jsonByteArray = []byte(rpcJsonString)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonByteArray))
	if err != nil {
		return stacktrace.Propagate(err, "Failed to send RPC request to geth node")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return stacktrace.NewError("Received non-200 status code rom admin RPC api: %v", resp.StatusCode)
	}
	return nil
}

