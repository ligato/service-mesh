// +build recover

package nsmd_integration_tests

import (
	"context"
	"fmt"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/registry"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/networkservicemesh/networkservicemesh/test/integration/utils"
	"github.com/networkservicemesh/networkservicemesh/test/kube_testing"
	"github.com/networkservicemesh/networkservicemesh/test/kube_testing/pods"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

func TestNSMHealLocalDieNSMD(t *testing.T) {
	RegisterTestingT(t)

	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}

	k8s, err := kube_testing.NewK8s(true)
	defer k8s.Cleanup()

	Expect(err).To(BeNil())

	// Deploy open tracing to see what happening.
	nodes_setup := utils.SetupNodes(k8s, 2, defaultTimeout)

	// Run ICMP on latest node
	icmpPod := utils.DeployICMP(k8s, nodes_setup[1].Node, "icmp-responder-nse-1", defaultTimeout)
	Expect(icmpPod).ToNot(BeNil())

	nscPodNode := utils.DeployNSC(k8s, nodes_setup[0].Node, "nsc-1", defaultTimeout)
	var nscInfo *utils.NSCCheckInfo
	failures := InterceptGomegaFailures(func() {
		nscInfo = utils.CheckNSC(k8s, t, nscPodNode)
	})
	// Do dumping of container state to dig into what is happened.
	utils.PrintErrors(failures, k8s, nodes_setup, nscInfo, t)

	logrus.Infof("Delete Local NSMD")
	k8s.DeletePods(nodes_setup[0].Nsmd)

	logrus.Infof("Waiting for NSE with network service")
	k8s.WaitLogsContains(nodes_setup[1].Nsmd, "nsmd", "NSM: Remote opened connection is not monitored and put into Healing state", defaultTimeout)
	// Now are are in dataplane dead state, and in Heal procedure waiting for dataplane.
	nsmdName := fmt.Sprintf("%s-recovered", nodes_setup[0].Nsmd.Name)

	logrus.Infof("Starting recovered NSMD...")
	startTime := time.Now()
	nodes_setup[0].Nsmd = k8s.CreatePod(pods.NSMgrPodWithConfig(nsmdName, nodes_setup[0].Node, &pods.NSMgrPodConfig{Namespace: k8s.GetK8sNamespace()})) // Recovery NSEs
	logrus.Printf("Started new NSMD: %v on node %s", time.Since(startTime), nodes_setup[0].Node.Name)

	failures = InterceptGomegaFailures(func() {
		logrus.Infof("Waiting for connection recovery...")
		k8s.WaitLogsContains(nodes_setup[0].Nsmd, "nsmd", "Heal: Connection recovered:", defaultTimeout)
		logrus.Infof("Waiting for connection recovery Done...")

		nscInfo = utils.CheckNSC(k8s, t, nscPodNode)
	})
	utils.PrintErrors(failures, k8s, nodes_setup, nscInfo, t)
}

func TestNSMHealLocalDieNSMDOneNode(t *testing.T) {
	RegisterTestingT(t)

	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}
	testNSMHealLocalDieNSMDOneNode(t, utils.DeployNSC, utils.DeployICMP, utils.CheckNSC, false)
}

func TestNSMHealLocalDieNSMDOneNodeMemif(t *testing.T) {
	RegisterTestingT(t)

	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}
	testNSMHealLocalDieNSMDOneNode(t, utils.DeployVppAgentNSC, utils.DeployVppAgentICMP, utils.CheckVppAgentNSC, false)
}

func TestNSMHealLocalDieNSMDOneNodeCleanedEndpoints(t *testing.T) {
	RegisterTestingT(t)

	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}
	testNSMHealLocalDieNSMDOneNode(t, utils.DeployNSC, utils.DeployICMP, utils.CheckNSC, true)
}

func testNSMHealLocalDieNSMDOneNode(t *testing.T, deployNsc, deployNse utils.PodSupplier, nscCheck utils.NscChecker, cleanupEndpointsCRDs bool) {
	k8s, err := kube_testing.NewK8s(true)
	defer k8s.Cleanup()

	Expect(err).To(BeNil())

	nodes_setup := utils.SetupNodes(k8s, 1, defaultTimeout)

	// Run ICMP on latest node
	icmpPod := deployNse(k8s, nodes_setup[0].Node, "icmp-responder-nse-1", defaultTimeout)
	Expect(icmpPod).ToNot(BeNil())

	nscPodNode := deployNsc(k8s, nodes_setup[0].Node, "nsc-1", defaultTimeout)
	var nscInfo *utils.NSCCheckInfo
	failures := InterceptGomegaFailures(func() {
		nscInfo = nscCheck(k8s, t, nscPodNode)
	})
	// Do dumping of container state to dig into what is happened.
	utils.PrintErrors(failures, k8s, nodes_setup, nscInfo, t)

	logrus.Infof("Delete Local NSMD")
	k8s.DeletePods(nodes_setup[0].Nsmd)

	nsmdName := fmt.Sprintf("%s-recovered", nodes_setup[0].Nsmd.Name)

	if cleanupEndpointsCRDs {
		logrus.Infof("Cleanup Endpoints...")
		k8s.CleanupEndpointsCRDs()
	}

	logrus.Infof("Starting recovered NSMD...")
	startTime := time.Now()
	nodes_setup[0].Nsmd = k8s.CreatePod(pods.NSMgrPodWithConfig(nsmdName, nodes_setup[0].Node, &pods.NSMgrPodConfig{Namespace: k8s.GetK8sNamespace()})) // Recovery NSEs
	logrus.Printf("Started new NSMD: %v on node %s", time.Since(startTime), nodes_setup[0].Node.Name)

	failures = InterceptGomegaFailures(func() {
		logrus.Infof("Waiting for connection recovery...")
		k8s.WaitLogsContains(nodes_setup[0].Nsmd, "nsmd", "Heal: Connection recovered:", defaultTimeout)
		logrus.Infof("Waiting for connection recovery Done...")

		nscInfo = nscCheck(k8s, t, nscPodNode)
	})
	utils.PrintErrors(failures, k8s, nodes_setup, nscInfo, t)
}

func TestNSMHealLocalDieNSMDOneNodeFakeEndpoint(t *testing.T) {
	RegisterTestingT(t)

	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}
	testNSMHealLocalDieNSMDTwoNodes(t, utils.DeployNSC, utils.DeployICMP, utils.CheckNSC)
}

func testNSMHealLocalDieNSMDTwoNodes(t *testing.T, deployNsc, deployNse utils.PodSupplier, nscCheck utils.NscChecker) {
	k8s, err := kube_testing.NewK8s(true)
	defer k8s.Cleanup()

	Expect(err).To(BeNil())

	nodes_setup := utils.SetupNodes(k8s, 2, defaultTimeout)

	// Run ICMP on latest node
	icmpPod := deployNse(k8s, nodes_setup[0].Node, "icmp-responder-nse-1", defaultTimeout)
	Expect(icmpPod).ToNot(BeNil())

	nscPodNode := deployNsc(k8s, nodes_setup[0].Node, "nsc-1", defaultTimeout)
	var nscInfo *utils.NSCCheckInfo
	failures := InterceptGomegaFailures(func() {
		nscInfo = nscCheck(k8s, t, nscPodNode)
	})
	// Do dumping of container state to dig into what is happened.
	utils.PrintErrors(failures, k8s, nodes_setup, nscInfo, t)

	// Remember nse name
	_, nsm1RegistryClient, fwd1Close := utils.PrepareRegistryClients(k8s, nodes_setup[0].Nsmd)
	nseList, err := nsm1RegistryClient.GetEndpoints(context.Background(), &empty.Empty{})
	fwd1Close()

	Expect(err).To(BeNil())
	Expect(len(nseList.NetworkServiceEndpoints)).To(Equal(1))
	nseName := nseList.NetworkServiceEndpoints[0].EndpointName

	logrus.Info(nseName)
	logrus.Infof("Delete Local NSMD")
	k8s.DeletePods(nodes_setup[0].Nsmd)

	nsmdName := fmt.Sprintf("%s-recovered", nodes_setup[0].Nsmd.Name)

	logrus.Infof("Cleanup Endpoints CRDs...")
	k8s.CleanupEndpointsCRDs()

	nse2RegistryClient, nsm2RegistryClient, fwd2Close := utils.PrepareRegistryClients(k8s, nodes_setup[1].Nsmd)
	defer fwd2Close()

	_, err = nse2RegistryClient.RegisterNSE(context.Background(), &registry.NSERegistration{
		NetworkService: &registry.NetworkService{
			Name:    "icmp-responder",
			Payload: "IP",
		},
		NetworkserviceEndpoint: &registry.NetworkServiceEndpoint{
			NetworkServiceName: "icmp-responder",
			EndpointName:       nseName,
		},
	})
	Expect(err).To(BeNil())
	nseList2, err := nsm2RegistryClient.GetEndpoints(context.Background(), &empty.Empty{})
	Expect(err).To(BeNil())
	Expect(len(nseList2.NetworkServiceEndpoints)).To(Equal(1))

	logrus.Infof("Starting recovered NSMD...")
	startTime := time.Now()
	nodes_setup[0].Nsmd = k8s.CreatePod(pods.NSMgrPodWithConfig(nsmdName, nodes_setup[0].Node, &pods.NSMgrPodConfig{Namespace: k8s.GetK8sNamespace()})) // Recovery NSEs
	logrus.Printf("Started new NSMD: %v on node %s", time.Since(startTime), nodes_setup[0].Node.Name)

	failures = InterceptGomegaFailures(func() {
		logrus.Infof("Waiting for connection recovery...")
		k8s.WaitLogsContains(nodes_setup[0].Nsmd, "nsmd", "Heal: Connection recovered:", defaultTimeout)
		logrus.Infof("Waiting for connection recovery Done...")

		nscInfo = nscCheck(k8s, t, nscPodNode)
	})
	utils.PrintErrors(failures, k8s, nodes_setup, nscInfo, t)
}
