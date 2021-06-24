Cassandra Onboarding Testsuite
=====================

## Implement the Basic Cassandra Test

1. Verify that the Docker daemon is running on your local machine.
2. Clone this repository by running `git clone https://github.com/galenmarchetti/kurtosis-onboarding-experience.git`
3. Run the empty Cassandra test to verify that everything works on your local machine.
    1. Run `bash scripts/build-and-run.sh all`
    2. Verify that the output of the build-and-run.sh script indicates that one test ran (cassandraTest) and that it passed.
4. Import the dependencies that are used in this example test suite.
    1. Run `go get github.com/gocql/gocql`
    2. Run `go get github.com/palantir/stacktrace`
5. Write the IsAvailable() function on the cassandra service to handle Cassandra’s availability condition.
    1. In your preferred IDE, open the Cassandra service definition at `testsuite/services_impl/cassandra_service/cassandra_service.go`
    2. Implement a GetIPAddress function on the Cassandra service to propagate the IPAddress forward.
        1. Add the GetIPAddress() function in [this Gist](https://gist.github.com/galenmarchetti/7958b0973f63081425091563578db1e9) to the Cassandra service file, between the constructor NewCassandraService() and the IsAvailable() function.
    3. Implement a “CreateSession” utility function on the Cassandra service to create a querying session so that later you can check if the service is up, or make queries against it.
        1. Add the CreateSession() in [this Gist](https://gist.github.com/galenmarchetti/03e41a50996279233f0d60cf23bfe331) to the Cassandra service file.
    4. Replace the default empty IsAvailable() function with the Cassandra-specific logic in [this Gist](https://gist.github.com/galenmarchetti/3f14080949a131d16a7b6204390a13ee)
    5. Verify that running `bash scripts/build-and-run.sh all` generates output indicating that one test ran (cassandraTest) and that it passed
6. Write the basic cassandra test logic to write a row to the cassandra node, and then read it back.
    1. In your preferred IDE, open the Cassandra test definition at `testsuite/testsuite_impl/cassandra_test/cassandra_test_.go`
    2. Add the helper functions `writeTweet` and `readAndConfirmTweet` in [this Gist](https://gist.github.com/galenmarchetti/98252fec7b92d2174d71ee7c72261bd3) to the bottom of the file, so they can be used later.
    3. Add the test logic in [this Gist](https://gist.github.com/galenmarchetti/118a2555749c7c47760cb58faa251795) to the Run() function in the Cassandra test, replacing the final return nil line with the return line from the code in the Gist.
    4. Verify that running `bash scripts/build-and-run.sh all` generates output indicating that one test ran (cassandraTest), and that the test contained business logic for a Cassandra test, and that it passed.

## Implement the Advanced Cassandra Test

   1. Modify the Cassandra Container Config Factor GetCreationConfig to expose the cluster communication port in addition to the native protocol client port using the WithUsedPorts() method on the NewContainerCreationConfigBuilder.
   2. Modify the Cassandra Container Config Factory GetRunConfig method to set the environment variable of a cassandra node appropriately to instruct a cassandra node to connect to an existing cluster.
       1. Modify the GetRunConfig method, using the WithEnvironmentVariableOverrides() function on the NewContainerRunConfigBuilder, to set an environment variable called CASSANDRA_SEEDS to the value of the "Cassandra Seed Node IP Address" on the factory struct.
   3. Modify the Setup() method of the cassandra_test file so that three nodes are set up instead of just one, with the second and third nodes pointing to the first as the “seed node”.
       1. Add two more service ids to the  []services.ServiceID to store the total three service ids of the services in your network.
       2. After creating the seed node, store the seed node IP address in a variable so that it can be used later to create the second and third nodes with the ability to connect to the existing cluster.
       3. Iterate through the remaining serviceIDs (second and third) with a for loop, adding them to the network, and waiting for availability before continuing to the next.
       4. Verify that running your testsuite still runs a passing test, although the setup method will take a lot longer now given the time it takes for cassandra nodes to sequentially enter a cluster.
   4. Modify the Run() method of the cassandra_test file so that the tweet is written to one node, and then read from all three nodes to verify that the data has been persisted across the entire cluster.
       1. Verify that running your testsuite returns a passing test where the tweet is read and confirmed from each node in the cluster.
