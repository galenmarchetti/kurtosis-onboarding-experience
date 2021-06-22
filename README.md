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
    2. Add the test logic in [this Gist](https://gist.github.com/galenmarchetti/118a2555749c7c47760cb58faa251795) to the Run() function in the Cassandra test, replacing the final return nil line with the return line from the code in the Gist.
    3. Verify that running `bash scripts/build-and-run.sh all` generates output indicating that one test ran (cassandraTest), and that the test contained business logic for a Cassandra test, and that it passed.

## Implement the Advanced Cassandra Test

1. Add a second cassandra node to your Kurtosis-deployed cassandra cluster.
    1. Add the "cluster communication port" as a used port to the container creation configuration builder in the GetCreationConfig method of `testsuite/services_impl/cassandra_service/cassandra_container_config_factory.go`.
    2. 
2. Add a third cassandra node to your Kurtosis-deployed cassandra cluster.
    1. TODO TODO TODO Fill this in
