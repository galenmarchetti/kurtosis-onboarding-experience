package my_test

import (
	"github.com/galenmarchetti/kurtosis-onboarding-test/testsuite/services_impl/my_service"
	"github.com/kurtosis-tech/kurtosis-libs/golang/lib/networks"
	"github.com/kurtosis-tech/kurtosis-libs/golang/lib/services"
	"github.com/kurtosis-tech/kurtosis-libs/golang/lib/testsuite"
	"github.com/palantir/stacktrace"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	waitForStartupTimeBetweenPolls = 1 * time.Second
	waitForStartupMaxPolls = 90
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
	/*
		NEW USER ONBOARDING:
		- Add services multiple times using the below logic in order to have more than one service.
	*/
	configFactory := my_service.NewMyServiceConfigFactory("ethereum/go-ethereum", "")


	uncastedService, _, availabilityChecker, err := networkCtx.AddService(serviceIDs[0], configFactory)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred adding the service")
	}
	if err := availabilityChecker.WaitForStartup(waitForStartupTimeBetweenPolls, waitForStartupMaxPolls); err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred waiting for the service to become available")
	}
	castedService := uncastedService.(*my_service.MyService)
	logrus.Infof("Added service with IP address: %v", castedService)

	return networkCtx, nil
}

func (test MyTest) Run(uncastedNetwork networks.Network) error {
	logrus.Infof("Running test.")
	// Necessary because Go doesn't have generics
	castedNetwork := uncastedNetwork.(*networks.NetworkContext)

	uncastedService, err := castedNetwork.GetService(serviceIDs[0])
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred getting the datastore service")
	}

	// Necessary again due to no Go generics
	castedService := uncastedService.(*my_service.MyService)
	logrus.Infof("Service is available: %v", castedService.IsAvailable())

	/*
		NEW USER ONBOARDING:
		- Fill in the logic necessary to run your custom test.
	*/
	return nil
}