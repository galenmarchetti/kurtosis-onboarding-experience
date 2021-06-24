package cassandra_test

import (
	"github.com/galenmarchetti/kurtosis-onboarding-test/testsuite/services_impl/cassandra_service"
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

var cassandraIds = []services.ServiceID{
	"cassandra-0",
}

type CassandraTest struct {
	CassandraServiceImage string
}

func NewCassandraTest(image string) *CassandraTest {
	return &CassandraTest{CassandraServiceImage: image}
}

func (test CassandraTest) Configure(builder *testsuite.TestConfigurationBuilder) {
	builder.WithSetupTimeoutSeconds(360).WithRunTimeoutSeconds(360)
}

func (test CassandraTest) Setup(networkCtx *networks.NetworkContext) (networks.Network, error) {
	logrus.Infof("Setting up cassandra test.")
	/*
		NEW USER ONBOARDING:
		- Add services multiple times using the below logic in order to have more than one service.
	*/
	configFactory := cassandra_service.NewCassandraServiceConfigFactory(test.CassandraServiceImage, "")
	service, _, availabilityChecker, err := networkCtx.AddService(cassandraIds[0], configFactory)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred adding the service")
	}
	if err := availabilityChecker.WaitForStartup(waitForStartupTimeBetweenPolls, waitForStartupMaxPolls); err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred waiting for the service to become available")
	}
	logrus.Infof("Added Cassandra service with IP address: %v", service)

	return networkCtx, nil
}

func (test CassandraTest) Run(uncastedNetwork networks.Network) error {
	logrus.Infof("Running cassandra test.")
	// Necessary because Go doesn't have generics
	castedNetwork := uncastedNetwork.(*networks.NetworkContext)

	uncastedService, err := castedNetwork.GetService(cassandraIds[0])
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred getting the datastore service")
	}

	// Necessary again due to no Go generics
	castedService := uncastedService.(*cassandra_service.CassandraService)
	logrus.Infof("Service is available: %v", castedService.IsAvailable())

	/*
		NEW USER ONBOARDING:
		- Fill in the logic necessary to run your custom test.
	*/
	return nil
}