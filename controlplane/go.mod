module github.com/networkservicemesh/networkservicemesh/controlplane

require (
	github.com/gogo/protobuf v1.2.2-0.20190723190241-65acae22fc9d
	github.com/golang/protobuf v1.3.2
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/networkservicemesh/networkservicemesh v0.1.0
	github.com/networkservicemesh/networkservicemesh/controlplane/api v0.0.0-00010101000000-000000000000
	github.com/networkservicemesh/networkservicemesh/dataplane/api v0.0.0-00010101000000-000000000000
	github.com/networkservicemesh/networkservicemesh/sdk v0.0.0-00010101000000-000000000000
	github.com/networkservicemesh/networkservicemesh/side-cars v0.0.0-00010101000000-000000000000
	github.com/onsi/gomega v1.5.1-0.20190520121345-efe19c39ca10
	github.com/opentracing/opentracing-go v1.1.0
	github.com/sirupsen/logrus v1.4.2
	github.com/uber/jaeger-lib v2.1.1+incompatible // indirect
	golang.org/x/net v0.0.0-20190812203447-cdfb69ac37fc
	golang.org/x/sys v0.0.0-20190618155005-516e3c20635f
	google.golang.org/grpc v1.23.0
)

replace (
	github.com/networkservicemesh/networkservicemesh => ../
	github.com/networkservicemesh/networkservicemesh/controlplane => ./
	github.com/networkservicemesh/networkservicemesh/controlplane/api => ./api
	github.com/networkservicemesh/networkservicemesh/dataplane => ../dataplane
	github.com/networkservicemesh/networkservicemesh/dataplane/api => ../dataplane/api
	github.com/networkservicemesh/networkservicemesh/k8s/api => ../k8s/api
	github.com/networkservicemesh/networkservicemesh/sdk => ../sdk
	github.com/networkservicemesh/networkservicemesh/side-cars => ../side-cars
)
