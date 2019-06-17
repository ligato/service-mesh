package tests

import (
	"github.com/networkservicemesh/networkservicemesh/test/cloudtest/pkg/commands"
	"github.com/networkservicemesh/networkservicemesh/test/cloudtest/pkg/config"
	"github.com/networkservicemesh/networkservicemesh/test/cloudtest/pkg/k8s"
	"github.com/networkservicemesh/networkservicemesh/test/cloudtest/pkg/providers/shell"
	"github.com/networkservicemesh/networkservicemesh/test/cloudtest/pkg/utils"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"os"
	"testing"
)

type testValidationFactory struct {
}

type testValidator struct {
	location string
	config   *config.ClusterProviderConfig
}

func (v *testValidator) Validate() error {
	// Validation is passed for now
	return nil
}

func (*testValidationFactory) CreateValidator(config *config.ClusterProviderConfig, location string) (k8s.KubernetesValidator, error) {
	return &testValidator{
		config:   config,
		location: location,
	}, nil
}

func TestShellProvider(t *testing.T) {
	RegisterTestingT(t)

	testConfig := &config.CloudTestConfig{
	}

	testConfig.Timeout = 300

	tmpDir, err := ioutil.TempDir(os.TempDir(), "cloud-test-temp")
	defer utils.ClearFolder(tmpDir, false)
	Expect(err).To(BeNil())

	testConfig.ConfigRoot = tmpDir
	createProvider(testConfig, "a_provider")
	createProvider(testConfig, "b_provider")

	testConfig.Executions = append(testConfig.Executions, &config.ExecutionConfig{
		Name:        "simple",
		Timeout:     15,
		PackageRoot: "./sample",
	})

	testConfig.Executions = append(testConfig.Executions, &config.ExecutionConfig{
		Name:        "simple_tagged",
		Timeout:     15,
		Tags:        []string{"basic"},
		PackageRoot: "./sample",
	})

	err, report := commands.PerformTesting(testConfig, &testValidationFactory{}, &commands.Arguments{})
	Expect(err.Error()).To(Equal("There is failed tests 4"))

	Expect(report).NotTo(BeNil())

	Expect(len(report.Suites)).To(Equal(2))
	Expect(report.Suites[0].Failures).To(Equal(2))
	Expect(report.Suites[0].Tests).To(Equal(6))
	Expect(len(report.Suites[0].TestCases)).To(Equal(6))

	// Do assertions
}

func createProvider(testConfig *config.CloudTestConfig, name string) {
	testConfig.Providers = append(testConfig.Providers, &config.ClusterProviderConfig{
		Timeout:    100,
		Name:       name,
		NodeCount:  1,
		Kind:       "shell",
		RetryCount: 1,
		Instances:  2,
		Scripts: map[string]string{
			shell.ConfigScript:  "echo ./.tests/config",
			shell.StartScript:   "echo started",
			shell.PrepareScript: "echo prepared",
			shell.InstallScript: "echo installed",
			shell.StopScript:    "echo stopped",
		},
		Enabled: true,
	})
}

func TestInvalidProvider(t *testing.T) {
	RegisterTestingT(t)

	testConfig := &config.CloudTestConfig{
	}

	testConfig.Timeout = 300

	tmpDir, err := ioutil.TempDir(os.TempDir(), "cloud-test-temp")
	defer utils.ClearFolder(tmpDir, false)
	Expect(err).To(BeNil())

	testConfig.ConfigRoot = tmpDir
	createProvider(testConfig, "a_provider")
	delete(testConfig.Providers[0].Scripts, shell.StartScript)

	testConfig.Executions = append(testConfig.Executions, &config.ExecutionConfig{
		Name:        "simple",
		Timeout:     2,
		PackageRoot: "./sample",
	})

	err, report := commands.PerformTesting(testConfig, &testValidationFactory{}, &commands.Arguments{})
	Expect(err.Error()).To(Equal("Failed to create cluster instance. Error Invalid start script"))

	Expect(report).To(BeNil())
	// Do assertions
}

func TestRequireEnvVars(t *testing.T) {
	RegisterTestingT(t)

	testConfig := &config.CloudTestConfig{
	}

	testConfig.Timeout = 300

	tmpDir, err := ioutil.TempDir(os.TempDir(), "cloud-test-temp")
	defer utils.ClearFolder(tmpDir, false)
	Expect(err).To(BeNil())

	testConfig.ConfigRoot = tmpDir

	createProvider(testConfig, "a_provider")

	testConfig.Providers[0].EnvCheck = append(testConfig.Providers[0].EnvCheck, []string{
		"KUBECONFIG", "QWE",
	}...)

	testConfig.Executions = append(testConfig.Executions, &config.ExecutionConfig{
		Name:        "simple",
		Timeout:     2,
		PackageRoot: "./sample",
	})

	err, report := commands.PerformTesting(testConfig, &testValidationFactory{}, &commands.Arguments{})
	Expect(err.Error()).To(Equal("Failed to create cluster instance. Error Environment variable are not specified  Required variables: [KUBECONFIG QWE]"))

	Expect(report).To(BeNil())
	// Do assertions
}

func TestRequireEnvVars_DEPS(t *testing.T) {
	RegisterTestingT(t)

	testConfig := &config.CloudTestConfig{
	}

	testConfig.Timeout = 300

	tmpDir, err := ioutil.TempDir(os.TempDir(), "cloud-test-temp")
	defer utils.ClearFolder(tmpDir, false)
	Expect(err).To(BeNil())

	testConfig.ConfigRoot = tmpDir

	createProvider(testConfig, "a_provider")

	testConfig.Providers[0].EnvCheck = append(testConfig.Providers[0].EnvCheck, "PACKET_AUTH_TOKEN")
	testConfig.Providers[0].EnvCheck = append(testConfig.Providers[0].EnvCheck, "PACKET_PROJECT_ID")

	_ = os.Setenv("PACKET_AUTH_TOKEN", "token")
	_ = os.Setenv("PACKET_PROJECT_ID", "id")

	testConfig.Providers[0].Env = append(testConfig.Providers[0].Env, []string{
		"CLUSTER_RULES_PREFIX=packet",
		"CLUSTER_NAME=$(cluster-name)-$(uuid)",
		"KUBECONFIG=$(tempdir)/config",
		"TERRAFORM_ROOT=$(tempdir)/terraform",
		"TF_VAR_auth_token=${PACKET_AUTH_TOKEN}",
		"TF_VAR_master_hostname=devci-${CLUSTER_NAME}-master",
		"TF_VAR_worker1_hostname=ci-${CLUSTER_NAME}-worker1",
		"TF_VAR_project_id=${PACKET_PROJECT_ID}",
		"TF_VAR_public_key=${TERRAFORM_ROOT}/sshkey.pub",
		"TF_VAR_public_key_name=key-${CLUSTER_NAME}",
		"TF_LOG=DEBUG",
	}...)

	testConfig.Executions = append(testConfig.Executions, &config.ExecutionConfig{
		Name:        "simple",
		Timeout:     2,
		PackageRoot: "./sample",
	})

	err, report := commands.PerformTesting(testConfig, &testValidationFactory{}, &commands.Arguments{})
	Expect(err.Error()).To(Equal("There is failed tests 2"))

	Expect(report).ToNot(BeNil())
	// Do assertions
}

