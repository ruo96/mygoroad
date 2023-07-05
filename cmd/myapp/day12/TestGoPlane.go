package day12

import (
	"fmt"
	"github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"time"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	ext_http "github.com/envoyproxy/go-control-plane/envoy/extensions/upstreams/http/v3"
	router "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/router/v3"
)

func makeHTTP2Listener(config string, peer string) {
	name := fmt.Sprintf("%s-%s", "mpc", "1100")
	routerConfig, _ := anypb.New(&router.Router{})
	cache.NewSnapshot("123", map[resource.Type][]types.Resource{
		resource.ClusterType: &cluster.Cluster{
			Name:                          name,
			ConnectTimeout:                durationpb.New(5 * time.Second),
			ClusterDiscoveryType:          &cluster.Cluster_Type{Type: cluster.Cluster_STRICT_DNS},
			LbPolicy:                      cluster.Cluster_ROUND_ROBIN,
			DnsLookupFamily:               cluster.Cluster_V4_ONLY,
			TypedExtensionProtocolOptions: makeTypedExtensionProtocolOptions(),
			LoadAssignment:                makeEndpoint(name, peer, peer),
			TransportSocket:               makeTransportSocketForUpstream(config),
		}
	})

}

func makeEndpoint(clusterName string, host string, port uint32) *endpoint.ClusterLoadAssignment {
	return &endpoint.ClusterLoadAssignment{
		ClusterName: clusterName,
		Endpoints: []*endpoint.LocalityLbEndpoints{{
			LbEndpoints: []*endpoint.LbEndpoint{{
				HostIdentifier: &endpoint.LbEndpoint_Endpoint{
					Endpoint: &endpoint.Endpoint{
						Address: &core.Address{
							Address: &core.Address_SocketAddress{
								SocketAddress: &core.SocketAddress{
									Protocol: core.SocketAddress_TCP,
									Address:  host,
									PortSpecifier: &core.SocketAddress_PortValue{
										PortValue: port,
									},
								},
							},
						},
					},
				},
			}},
		}},
	}
}

func makeTypedExtensionProtocolOptions() map[string]*anypb.Any {
	httpProtocolOptions, err := anypb.New(&ext_http.HttpProtocolOptions{
		UpstreamProtocolOptions: &ext_http.HttpProtocolOptions_ExplicitHttpConfig_{
			ExplicitHttpConfig: &ext_http.HttpProtocolOptions_ExplicitHttpConfig{
				ProtocolConfig: &ext_http.HttpProtocolOptions_ExplicitHttpConfig_Http2ProtocolOptions{
					Http2ProtocolOptions: &core.Http2ProtocolOptions{},
				},
			},
		},
	})
	if httpProtocolOptions == nil && err != nil {
		panic(err)
	}

	return map[string]*anypb.Any{
		"envoy.extensions.upstreams.http.v3.HttpProtocolOptions": httpProtocolOptions,
	}
}

func makeTransportSocketForUpstream(config *XDSServerConfig) *core.TransportSocket {
	tlsc := &auth.UpstreamTlsContext{
		CommonTlsContext: &auth.CommonTlsContext{
			TlsCertificateSdsSecretConfigs: []*auth.SdsSecretConfig{{
				Name: config.Server.TlsCert,
			}},
			ValidationContextType: &auth.CommonTlsContext_ValidationContextSdsSecretConfig{
				ValidationContextSdsSecretConfig: &auth.SdsSecretConfig{
					Name: config.Server.TlsValidationContext,
				},
			},
		},
	}

	mt, _ := anypb.New(tlsc)
	return &core.TransportSocket{
		Name: wellknown.TransportSocketTLS,
		ConfigType: &core.TransportSocket_TypedConfig{
			TypedConfig: mt,
		},
	}
}
