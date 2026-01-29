// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

const (
	// RDSHttpProxy is the special name for HTTP PROXY route
	RDSHttpProxy = "http_proxy"

	// VirtualOutboundListenerName is the name for traffic capture listener
	VirtualOutboundListenerName = "virtualOutbound"

	// VirtualOutboundCatchAllTCPFilterChainName is the name of the catch all tcp filter chain
	VirtualOutboundCatchAllTCPFilterChainName = "virtualOutbound-catchall-tcp"

	// VirtualOutboundHTTPRouteConfigName is the RDS route config name for the virtual outbound listener's HCM.
	VirtualOutboundHTTPRouteConfigName = "virtualOutbound-http"

	// VirtualOutboundHTTPFilterChainName is the name of the http filter chain on the virtual outbound listener.
	VirtualOutboundHTTPFilterChainName = "virtualOutbound-http"

	// VirtualOutboundBlackholeFilterChainName is the name of the filter chain to blackhole undesired traffic
	VirtualOutboundBlackholeFilterChainName = "virtualOutbound-blackhole"
	// VirtualInboundBlackholeFilterChainName is the name of the filter chain to blackhole undesired traffic
	VirtualInboundBlackholeFilterChainName = "virtualInbound-blackhole"

	// VirtualInboundListenerName is the name for traffic capture listener
	VirtualInboundListenerName = "virtualInbound"

	// VirtualInboundCatchAllHTTPFilterChainName is the name of the catch all http filter chain
	VirtualInboundCatchAllHTTPFilterChainName = "virtualInbound-catchall-http"
)

const (

	// HBoneInboundListenPort is the port on which incoming HBone traffic will be captured.
	HBoneInboundListenPort = 15008
)
