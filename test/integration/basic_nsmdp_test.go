// +build basic

package nsmd_integration_tests

import (
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/nsmd"
	"github.com/networkservicemesh/networkservicemesh/test/integration/utils"
	"github.com/networkservicemesh/networkservicemesh/test/kube_testing"
	"github.com/networkservicemesh/networkservicemesh/test/kube_testing/pods"
	. "github.com/onsi/gomega"
	"testing"
)

func TestNSMDDP(t *testing.T) {
	RegisterTestingT(t)

	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}

	k8s, err := kube_testing.NewK8s(true)
	defer k8s.Cleanup()

	Expect(err).To(BeNil())

	nodes, err := utils.SetupNodes(k8s, 1, defaultTimeout)
	Expect(err).To(BeNil())
	icmpPod := nsmd_test_utils.DeployICMP(k8s, nodes[0].Node, "icmp-responder-nse-1", defaultTimeout)

	nsmdName := nodes[0].Nsmd.Name
	k8s.DeletePods(nodes[0].Nsmd, icmpPod)
	nodes[0].Nsmd = k8s.CreatePod(pods.NSMgrPod(nsmdName, nodes[0].Node, k8s.GetK8sNamespace())) // Recovery NSEs
	icmpPod = utils.DeployICMP(k8s, nodes[0].Node, "icmp-responder-nse-2", defaultTimeout)
	Expect(icmpPod).ToNot(BeNil())
}

func TestNSMDRecoverNSE(t *testing.T) {
	RegisterTestingT(t)

	if testing.Short() {
		t.Skip("Skip, please run without -short")
		return
	}

	k8s, err := kube_testing.NewK8s(true)
	defer k8s.Cleanup()

	Expect(err).To(BeNil())

	nodes, err := utils.SetupNodesConfig(k8s, 1, defaultTimeout, []*pods.NSMgrPodConfig{
		&pods.NSMgrPodConfig{
			Variables: map[string]string{
				nsmd.NsmdDeleteLocalRegistry: "true",
			},
			Namespace:          k8s.GetK8sNamespace(),
			DataplaneVariables: utils.DefaultDataplaneVariables(),
		},
	}, k8s.GetK8sNamespace())
	Expect(err).To(BeNil())
	icmpPod := utils.DeployICMP(k8s, nodes[0].Node, "icmp-responder-nse-1", defaultTimeout)

	nsmdName := nodes[0].Nsmd.Name

	k8s.DeletePods(nodes[0].Nsmd, icmpPod)

	nodes[0].Nsmd = k8s.CreatePod(pods.NSMgrPodWithConfig(nsmdName, nodes[0].Node, &pods.NSMgrPodConfig{
		Namespace: k8s.GetK8sNamespace(),
	})) // Recovery NSEs
	icmpPod = utils.DeployICMP(k8s, nodes[0].Node, "icmp-responder-nse-2", defaultTimeout)
	Expect(icmpPod).ToNot(BeNil())
}
