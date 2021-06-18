/*
 * Copyright (c) 2021 - present Kurtosis Technologies LLC.
 * All Rights Reserved.
 */

package execution_impl

import (
	"encoding/json"
	"github.com/kurtosis-tech/kurtosis-libs/golang/lib/testsuite"
	"github.com/galenmarchetti/kurtosis-onboarding-test/testsuite/testsuite_impl"
	"github.com/palantir/stacktrace"
	"github.com/sirupsen/logrus"
	"strings"
)

type CassandraTestsuiteConfigurator struct {}

func NewCassandraTestsuiteConfigurator() *CassandraTestsuiteConfigurator {
	return &CassandraTestsuiteConfigurator{}
}

func (t CassandraTestsuiteConfigurator) SetLogLevel(logLevelStr string) error {
	level, err := logrus.ParseLevel(logLevelStr)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred parsing loglevel string '%v'", logLevelStr)
	}
	logrus.SetLevel(level)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	return nil
}

func (t CassandraTestsuiteConfigurator) ParseParamsAndCreateSuite(paramsJsonStr string) (testsuite.TestSuite, error) {
	paramsJsonBytes := []byte(paramsJsonStr)
	var args CassandraTestsuiteArgs
	if err := json.Unmarshal(paramsJsonBytes, &args); err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred deserializing the testsuite params JSON")
	}

	if err := validateArgs(args); err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred validating the deserialized testsuite params")
	}

	/*
		NEW USER ONBOARDING:
		- Change the "CassandraImage" argument here to your own actual custom service image.
	*/
	suite := testsuite_impl.NewCassandraTestsuite(
		args.CassandraImage)
	return suite, nil
}

func validateArgs(args CassandraTestsuiteArgs) error {
	if strings.TrimSpace(args.CassandraImage) == "" {
		return stacktrace.NewError("Custom service image is empty.")
	}
	return nil
}
