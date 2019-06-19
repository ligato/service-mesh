package nsm

import (
	"context"
	"fmt"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	net_context "golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/crossconnect"
	local_connection "github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/local/connection"
	local_networkservice "github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/local/networkservice"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/nsm"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/nsm/connection"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/nsm/networkservice"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/registry"
	remote_connection "github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/remote/connection"
	remote_networkservice "github.com/networkservicemesh/networkservicemesh/controlplane/pkg/apis/remote/networkservice"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/model"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/nsmd"
	"github.com/networkservicemesh/networkservicemesh/controlplane/pkg/serviceregistry"
	test_utils "github.com/networkservicemesh/networkservicemesh/controlplane/pkg/tests/utils"
)

const (
	networkServiceName = "golden_network"

	localNSMName  = "nsm-local"
	remoteNSMName = "nsm-remote"

	dataplane1Name = "dataplane-1"

	nse1Name = "nse-1"
	nse2Name = "nse-2"
)

type healTestData struct {
	model model.Model

	serviceRegistry   *serviceRegistryStub
	connectionManager *connectionManagerStub
	nseManager        *nseManagerStub

	healProcessor *healProcessor
}

func newHealTestData() *healTestData {
	var data = &healTestData{
		model: model.NewModel(),
	}

	data.serviceRegistry = &serviceRegistryStub{
		discoveryClient: &discoveryClientStub{
			response: data.createFindNetworkServiceResponse(),
		},
	}
	data.connectionManager = &connectionManagerStub{
		model: data.model,
	}
	data.nseManager = &nseManagerStub{
		model: data.model,

		nseClients: map[string]*nseClientStub{},
		nses:       []*registry.NSERegistration{},
	}

	data.healProcessor = &healProcessor{
		serviceRegistry: data.serviceRegistry,
		model:           data.model,
		properties: &nsm.NsmProperties{
			HealEnabled: true,
		},

		conManager: data.connectionManager,
		nseManager: data.nseManager,
	}

	data.model.SetNsm(&registry.NetworkServiceManager{
		Name: localNSMName,
	})
	data.model.AddOrUpdateDataplane(&model.Dataplane{
		RegisteredName:       dataplane1Name,
		MechanismsConfigured: true,
	})

	return data
}

func TestHealDstDown_RemoteClientLocalEndpoint(t *testing.T) {
	RegisterTestingT(t)
	data := newHealTestData()

	nse1 := data.createEndpoint(nse1Name, localNSMName)
	data.model.AddOrUpdateEndpoint(&model.Endpoint{
		Endpoint: nse1,
	})

	xcon := data.createCrossConnection(true, false, "src", "dst")
	request := data.createRequest(true)
	clientConnection := data.createClientConnection(xcon, nse1, remoteNSMName, request)
	editor, _ := data.model.AddClientConnection("id", model.ClientConnectionHealing, clientConnection)

	healed := data.healProcessor.healDstDown("", editor.ClientConnection)
	Expect(healed).To(BeFalse())

	test_utils.NewModelVerifier(data.model).
		EndpointExists(nse1Name, localNSMName).
		ClientConnectionExists("id", "src", "dst", remoteNSMName, nse1Name, dataplane1Name).
		DataplaneExists(dataplane1Name).
		Verify(t)
}

func TestHealDstDown_LocalClientLocalEndpoint(t *testing.T) {
	RegisterTestingT(t)
	data := newHealTestData()

	nse1 := data.createEndpoint(nse1Name, localNSMName)

	nse2 := data.createEndpoint(nse2Name, localNSMName)
	data.model.AddOrUpdateEndpoint(&model.Endpoint{
		Endpoint: nse2,
	})

	xcon := data.createCrossConnection(false, false, "src", "dst")
	request := data.createRequest(false)
	clientConnection := data.createClientConnection(xcon, nse1, localNSMName, request)
	editor, _ := data.model.AddClientConnection("id", model.ClientConnectionHealing, clientConnection)

	data.serviceRegistry.discoveryClient.response = data.createFindNetworkServiceResponse(nse2)

	healed := data.healProcessor.healDstDown("", editor.ClientConnection)
	Expect(healed).To(BeTrue())

	test_utils.NewModelVerifier(data.model).
		EndpointNotExists(nse1Name).
		EndpointExists(nse2Name, localNSMName).
		ClientConnectionExists("id", "src", "dst", localNSMName, nse2Name, dataplane1Name).
		DataplaneExists(dataplane1Name).
		Verify(t)
}

func TestHealDstDown_LocalClientLocalEndpoint_NoNSEFound(t *testing.T) {
	RegisterTestingT(t)
	data := newHealTestData()

	nse1 := data.createEndpoint(nse1Name, localNSMName)
	data.model.AddOrUpdateEndpoint(&model.Endpoint{
		Endpoint: nse1,
	})

	xcon := data.createCrossConnection(false, false, "src", "dst")
	request := data.createRequest(false)
	clientConnection := data.createClientConnection(xcon, nse1, localNSMName, request)
	editor, _ := data.model.AddClientConnection("id", model.ClientConnectionHealing, clientConnection)

	healed := data.healProcessor.healDstDown("", editor.ClientConnection)
	Expect(healed).To(BeFalse())

	test_utils.NewModelVerifier(data.model).
		EndpointExists(nse1Name, localNSMName).
		ClientConnectionNotExists("id").
		DataplaneExists(dataplane1Name).
		Verify(t)
}

func TestHealDstDown_LocalClientLocalEndpoint_RequestFailed(t *testing.T) {
	RegisterTestingT(t)
	data := newHealTestData()

	nse1 := data.createEndpoint(nse1Name, localNSMName)
	data.model.AddOrUpdateEndpoint(&model.Endpoint{
		Endpoint: nse1,
	})

	xcon := data.createCrossConnection(false, false, "src", "dst")
	request := data.createRequest(false)
	clientConnection := data.createClientConnection(xcon, nse1, localNSMName, request)
	editor, _ := data.model.AddClientConnection("id", model.ClientConnectionHealing, clientConnection)

	data.connectionManager.requestError = fmt.Errorf("request error")

	healed := data.healProcessor.healDstDown("", editor.ClientConnection)
	Expect(healed).To(BeFalse())

	test_utils.NewModelVerifier(data.model).
		EndpointExists(nse1Name, localNSMName).
		DataplaneExists(dataplane1Name).
		Verify(t)
}

func TestHealDstDown_LocalClientRemoteEndpoint(t *testing.T) {
	RegisterTestingT(t)
	data := newHealTestData()

	nse1 := data.createEndpoint(nse1Name, remoteNSMName)

	nse2 := data.createEndpoint(nse2Name, remoteNSMName)
	data.nseManager.nses = append(data.nseManager.nses, nse2)

	xcon := data.createCrossConnection(false, true, "src", "dst")
	request := data.createRequest(true)
	clientConnection := data.createClientConnection(xcon, nse1, remoteNSMName, request)
	editor, _ := data.model.AddClientConnection("id", model.ClientConnectionHealing, clientConnection)

	data.serviceRegistry.discoveryClient.response = data.createFindNetworkServiceResponse(nse2)
	data.connectionManager.nse = nse2

	healed := data.healProcessor.healDstDown("", editor.ClientConnection)
	Expect(healed).To(BeTrue())

	Expect(data.nseManager.nseClients[nse1Name].cleanedUp).To(BeTrue())
	Expect(data.nseManager.nseClients[nse2Name]).To(BeNil())

	test_utils.NewModelVerifier(data.model).
		EndpointNotExists(nse1Name).
		EndpointNotExists(nse2Name).
		ClientConnectionExists("id", "src", "dst", remoteNSMName, nse2Name, dataplane1Name).
		DataplaneExists(dataplane1Name).
		Verify(t)
}

func TestHealDstDown_LocalClientRemoteEndpoint_NoNSEFound(t *testing.T) {
	RegisterTestingT(t)
	data := newHealTestData()

	nse1 := data.createEndpoint(nse1Name, remoteNSMName)

	xcon := data.createCrossConnection(false, true, "src", "dst")
	request := data.createRequest(true)
	clientConnection := data.createClientConnection(xcon, nse1, remoteNSMName, request)
	editor, _ := data.model.AddClientConnection("id", model.ClientConnectionHealing, clientConnection)

	healed := data.healProcessor.healDstDown("", editor.ClientConnection)
	Expect(healed).To(BeFalse())

	Expect(data.nseManager.nseClients[nse1Name].cleanedUp).To(BeTrue())

	test_utils.NewModelVerifier(data.model).
		EndpointNotExists(nse1Name).
		ClientConnectionNotExists("id").
		DataplaneExists(dataplane1Name).
		Verify(t)
}

type discoveryClientStub struct {
	response *registry.FindNetworkServiceResponse
	error    error
}

func (stub *discoveryClientStub) FindNetworkService(ctx net_context.Context, in *registry.FindNetworkServiceRequest, opts ...grpc.CallOption) (*registry.FindNetworkServiceResponse, error) {
	Expect(in.GetNetworkServiceName()).To(Equal(networkServiceName))
	return stub.response, stub.error
}

type serviceRegistryStub struct {
	discoveryClient *discoveryClientStub
	error           error

	serviceregistry.ServiceRegistry
}

func (stub *serviceRegistryStub) DiscoveryClient() (registry.NetworkServiceDiscoveryClient, error) {
	return stub.discoveryClient, stub.error
}

func (stub *serviceRegistryStub) WaitForDataplaneAvailable(model model.Model, timeout time.Duration) error {
	return nsmd.NewServiceRegistry().WaitForDataplaneAvailable(model, timeout)
}

type connectionManagerStub struct {
	model model.Model

	requestError error
	nse          *registry.NSERegistration

	closeError error
}

func (stub *connectionManagerStub) request(ctx context.Context, request networkservice.Request, existingCC *model.ClientConnection) (connection.Connection, error) {
	if stub.requestError != nil {
		return nil, stub.requestError
	}

	editor, err := stub.model.ChangeClientConnectionState(existingCC.GetID(), model.ClientConnectionRequesting)
	if err != nil {
		return nil, err
	}

	if existingCC.GetConnectionDestination().GetId() == "-" {
		if stub.nse != nil {
			existingCC.Endpoint = stub.nse
		} else {
			_ = stub.model.DeleteClientConnection(existingCC.GetID())
			return nil, fmt.Errorf("no NSE available")
		}
	}

	editor.DataplaneState = model.DataplaneStateReady

	if err = stub.model.CommitClientConnectionChanges(editor); err != nil {
		return nil, err
	}

	return request.GetRequestConnection().Clone(), nil
}

func (stub *connectionManagerStub) Close(ctx context.Context, clientConnection nsm.ClientConnection) error {
	if stub.closeError != nil {
		return stub.closeError
	}

	cc := clientConnection.(*model.ClientConnection)

	if _, err := stub.model.ChangeClientConnectionState(clientConnection.GetID(), model.ClientConnectionClosing); err != nil {
		return err
	}

	_ = stub.model.DeleteEndpoint(cc.Endpoint.GetNetworkserviceEndpoint().GetEndpointName())
	_ = stub.model.DeleteDataplane(cc.DataplaneRegisteredName)
	_ = stub.model.DeleteClientConnection(cc.GetID())

	return nil
}

type nseClientStub struct {
	cleanedUp bool

	nsm.NetworkServiceClient
}

func (stub *nseClientStub) Cleanup() error {
	stub.cleanedUp = true
	return nil
}

type nseManagerStub struct {
	model model.Model

	clientError error
	nseClients  map[string]*nseClientStub

	nses []*registry.NSERegistration
}

func (stub *nseManagerStub) getEndpoint(ctx context.Context, requestConnection connection.Connection, ignoreEndpoints map[string]*registry.NSERegistration) (*registry.NSERegistration, error) {
	panic("implement me")
}

func (stub *nseManagerStub) createNSEClient(ctx context.Context, endpoint *registry.NSERegistration) (nsm.NetworkServiceClient, error) {
	if stub.clientError != nil {
		return nil, stub.clientError
	}

	nseClient := &nseClientStub{
		cleanedUp: false,
	}
	stub.nseClients[endpoint.GetNetworkserviceEndpoint().GetEndpointName()] = nseClient

	return nseClient, nil
}

func (stub *nseManagerStub) isLocalEndpoint(endpoint *registry.NSERegistration) bool {
	return stub.model.GetNsm().GetName() == endpoint.GetNetworkserviceEndpoint().GetNetworkServiceManagerName()
}

func (stub *nseManagerStub) checkUpdateNSE(ctx context.Context, reg *registry.NSERegistration) bool {
	for _, nse := range stub.nses {
		if nse.GetNetworkserviceEndpoint().GetEndpointName() == reg.GetNetworkserviceEndpoint().GetEndpointName() {
			return true
		}
	}

	return false
}

func (data *healTestData) createEndpoint(nse, nsm string) *registry.NSERegistration {
	return &registry.NSERegistration{
		NetworkService: &registry.NetworkService{
			Name: networkServiceName,
		},
		NetworkServiceManager: &registry.NetworkServiceManager{
			Name: nsm,
		},
		NetworkserviceEndpoint: &registry.NetworkServiceEndpoint{
			NetworkServiceName:        networkServiceName,
			NetworkServiceManagerName: nsm,
			EndpointName:              nse,
		},
	}
}

func (data *healTestData) createCrossConnection(isRemoteSrc, isRemoteDst bool, srcID, dstID string) *crossconnect.CrossConnect {
	xcon := &crossconnect.CrossConnect{}

	if isRemoteSrc {
		xcon.SetSourceConnection(&remote_connection.Connection{Id: srcID})
	} else {
		xcon.SetSourceConnection(&local_connection.Connection{Id: srcID})
	}

	if isRemoteDst {
		xcon.SetDestinationConnection(&remote_connection.Connection{Id: dstID})
	} else {
		xcon.SetDestinationConnection(&local_connection.Connection{Id: dstID})
	}

	return xcon
}

func (data *healTestData) createRequest(isRemote bool) networkservice.Request {
	if isRemote {
		return &remote_networkservice.NetworkServiceRequest{
			Connection: &remote_connection.Connection{
				NetworkService: networkServiceName,
			},
		}
	}

	return &local_networkservice.NetworkServiceRequest{
		Connection: &local_connection.Connection{
			NetworkService: networkServiceName,
		},
	}
}

func (data *healTestData) createClientConnection(xcon *crossconnect.CrossConnect, nse *registry.NSERegistration, nsm string, request networkservice.Request) *model.ClientConnection {
	return &model.ClientConnection{
		Xcon: xcon,
		RemoteNsm: &registry.NetworkServiceManager{
			Name: nsm,
		},
		Endpoint:                nse,
		DataplaneRegisteredName: dataplane1Name,
		Request:                 request,
		DataplaneState:          model.DataplaneStateReady,
	}
}

func (data *healTestData) createFindNetworkServiceResponse(nses ...*registry.NSERegistration) *registry.FindNetworkServiceResponse {
	response := &registry.FindNetworkServiceResponse{
		NetworkService: &registry.NetworkService{
			Name: networkServiceName,
		},
		NetworkServiceManagers:  map[string]*registry.NetworkServiceManager{},
		NetworkServiceEndpoints: []*registry.NetworkServiceEndpoint{},
	}

	for _, nse := range nses {
		nsm := nse.GetNetworkServiceManager().GetName()
		response.NetworkServiceManagers[nsm] = &registry.NetworkServiceManager{
			Name: nsm,
		}
		response.NetworkServiceEndpoints = append(response.NetworkServiceEndpoints, nse.NetworkserviceEndpoint)
	}

	return response
}
