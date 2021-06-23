package cassandra_test

import (
	"github.com/galenmarchetti/kurtosis-onboarding-test/testsuite/services_impl/cassandra_service"
	"github.com/gocql/gocql"
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
	"cassandra-1",
	"cassandra-2",
}

type CassandraTest struct {
	CassandraServiceImage string
}

func NewCassandraTest(image string) *CassandraTest {
	return &CassandraTest{CassandraServiceImage: image}
}

func (test *CassandraTest) Configure(builder *testsuite.TestConfigurationBuilder) {
	builder.WithSetupTimeoutSeconds(360).WithRunTimeoutSeconds(360)
}

func (test *CassandraTest) Setup(networkCtx *networks.NetworkContext) (networks.Network, error) {
	logrus.Infof("Setting up cassandra test.")

	// Start the seed node, which doesn't have a cluster to join.
	configFactory := cassandra_service.NewCassandraServiceConfigFactory(test.CassandraServiceImage, "")
	service, hostPortBindings, availabilityChecker, err := networkCtx.AddService(cassandraIds[0], configFactory)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred adding the service")
	}
	// Get the seed node IP address, so that subsequent nodes can join its cluster.
	seedIP := service.(*cassandra_service.CassandraService).GetIPAddress()

	// Add the remaining nodes, waiting for each to become available before starting the next.
	for i := 1; i < 3; i++ {
		configFactory = cassandra_service.NewCassandraServiceConfigFactory(test.CassandraServiceImage, seedIP)
		_, hostPortBindings, availabilityChecker, err = networkCtx.AddService(cassandraIds[i], configFactory)
		if err != nil {
			return nil, stacktrace.Propagate(err, "An error occurred adding the service")
		}
		if err := availabilityChecker.WaitForStartup(waitForStartupTimeBetweenPolls, waitForStartupMaxPolls); err != nil {
			return nil, stacktrace.Propagate(err, "An error occurred waiting for the service to become available")
		}
		logrus.Infof("Added service with host port bindings: %+v", hostPortBindings)
	}
	return networkCtx, nil
}

func (test *CassandraTest) Run(uncastedNetwork networks.Network) error {
	logrus.Infof("Running cassandra tests.")
	logrus.Infof("Test object: %+v", test)
	// Necessary because Go doesn't have generics
	castedNetwork := uncastedNetwork.(*networks.NetworkContext)
	logrus.Infof("casted network")

	uncastedService, err := castedNetwork.GetService(cassandraIds[0])
	logrus.Infof("got uncasted service")
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred getting the cassandra service")
	}
	// Necessary again due to no Go generics
	castedService := uncastedService.(*cassandra_service.CassandraService)

	logrus.Infof("About to get a cassandra seedSession.")

	seedSession, err := castedService.CreateSession()
	if err != nil {
		return stacktrace.Propagate(err, "Failed to create seedSession on the cassandra service.")
	}
	defer seedSession.Close()

	err = writeTweet(seedSession)
	if err != nil {
		return stacktrace.Propagate(err, "Failed to write tweet")
	}

	for i := 0; i < 3; i++ {
		logrus.Infof("Reading tweet from node %+v", cassandraIds[i])
		err := readAndConfirmTweet(castedNetwork, cassandraIds[i])
		if err != nil {
			return stacktrace.Propagate(err, "Test failed reading tweet from node %v", cassandraIds[i])
		}
	}
	return nil
}

func writeTweet(session *gocql.Session) error {
	// Create a keyspace "test" to use for testing purposes
	logrus.Infof("Creating a keyspace in Cassandra.")
	err := session.Query(`CREATE KEYSPACE test WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 3 }`).Exec()
	if err != nil {
		return stacktrace.Propagate(err, "Failed to create cassandra keyspace.")
	}

	// Create a table "tweet" to use for testing purposes
	logrus.Infof("Creating a 'tweet' table in Cassandra.")
	err = session.Query(`CREATE TABLE test.tweet(timeline text, id timeuuid PRIMARY KEY, text text)`).Exec()
	if err != nil {
		return stacktrace.Propagate(err, "Failed to create tweet table.")
	}

	// Insert a tweet into "tweet" table
	logrus.Infof("Inserting a tweet into the 'tweet' table in Cassandra.")
	err = session.Query(`INSERT INTO test.tweet (timeline, id, text) VALUES (?, ?, ?)`, "me", gocql.TimeUUID(), "hello world").Exec()
	if err != nil {
		return stacktrace.Propagate(err, "Failed to insert tweet into table.")
	}
	return nil
}


func readAndConfirmTweet(network *networks.NetworkContext, nodeId services.ServiceID) error {
	uncastedService, err := network.GetService(nodeId)
	logrus.Infof("got uncasted service")
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred getting the cassandra service")
	}
	// Necessary again due to no Go generics
	castedService := uncastedService.(*cassandra_service.CassandraService)

	logrus.Infof("About to get a cassandra session.")

	session, err := castedService.CreateSession()
	if err != nil {
		return stacktrace.Propagate(err, "Failed to create session on the cassandra service.")
	}
	defer session.Close()
	// Read a tweet from "tweet" table
	logrus.Infof("Reading a tweet from the 'tweet' table in Cassandra.")
	var (
		id   gocql.UUID
		text string
	)

	err = session.Query(`SELECT id, text FROM test.tweet WHERE timeline = ? LIMIT 1 ALLOW FILTERING`, "me").Scan(&id, &text)
	if err != nil {
		return stacktrace.Propagate(err, "Failed to read lines from tweet table.")
	}

	logrus.Infof("Verifying that the read tweet is identical to the written tweet")
	if text == "hello world" {
		logrus.Infof("Tweet verified, test passed.")
		return nil
	}
	return stacktrace.NewError("Cassandra test failed, text is: %v", text)
}