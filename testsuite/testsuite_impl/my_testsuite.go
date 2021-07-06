package testsuite_impl

import (
	"github.com/galenmarchetti/kurtosis-onboarding-test/testsuite/testsuite_impl/my_test"
	"github.com/kurtosis-tech/kurtosis-libs/golang/lib/testsuite"
)

type MyTestsuite struct {}

func (suite MyTestsuite) GetTests() map[string]testsuite.Test {
	tests := map[string]testsuite.Test{
		"myTest": &my_test.MyTest{},
	}
	return tests
}

func (suite MyTestsuite) GetNetworkWidthBits() uint32 {
	return 8
}
