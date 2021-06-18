/*
 * Copyright (c) 2020 - present Kurtosis Technologies LLC.
 * All Rights Reserved.
 */

package testsuite_impl

import (
	"github.com/galenmarchetti/kurtosis-onboarding-test/testsuite/testsuite_impl/cassandra_test"
	"github.com/kurtosis-tech/kurtosis-libs/golang/lib/testsuite"
)

type CassandraTestsuite struct {
	cassandraServiceImage string
}

func NewCassandraTestsuite(cassandraImage string) *CassandraTestsuite {
	return &CassandraTestsuite{cassandraServiceImage: cassandraImage}
}

func (suite CassandraTestsuite) GetTests() map[string]testsuite.Test {
	tests := map[string]testsuite.Test{
		"cassandraTest": cassandra_test.NewCassandraTest(suite.cassandraServiceImage),
	}

	return tests
}

func (suite CassandraTestsuite) GetNetworkWidthBits() uint32 {
	return 8
}

