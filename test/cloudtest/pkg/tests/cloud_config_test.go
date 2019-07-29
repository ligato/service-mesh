package tests

import (
	"io/ioutil"
	"testing"

	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"

	"github.com/networkservicemesh/networkservicemesh/test/cloudtest/pkg/config"
)

func TestClusterConfiguration(t *testing.T) {
	RegisterTestingT(t)

	var testConfig config.CloudTestConfig

	file1, err := ioutil.ReadFile("./config1.yaml")
	Expect(err).To(BeNil())

	err = yaml.Unmarshal(file1, &testConfig)
	Expect(err).To(BeNil())
	Expect(len(testConfig.Providers)).To(Equal(3))
	Expect(testConfig.Reporting.JUnitReportFile).To(Equal("./.tests/junit.xml"))

}
