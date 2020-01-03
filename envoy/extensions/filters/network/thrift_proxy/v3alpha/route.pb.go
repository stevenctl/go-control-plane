// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/network/thrift_proxy/v3alpha/route.proto

package envoy_extensions_filters_network_thrift_proxy_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3alpha1 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3alpha"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/config/route/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RouteConfiguration struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Routes               []*Route `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_35b85368996c9dcc, []int{0}
}

func (m *RouteConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfiguration.Unmarshal(m, b)
}
func (m *RouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfiguration.Marshal(b, m, deterministic)
}
func (m *RouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfiguration.Merge(m, src)
}
func (m *RouteConfiguration) XXX_Size() int {
	return xxx_messageInfo_RouteConfiguration.Size(m)
}
func (m *RouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfiguration proto.InternalMessageInfo

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type Route struct {
	Match                *RouteMatch  `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	Route                *RouteAction `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_35b85368996c9dcc, []int{1}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetMatch() *RouteMatch {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *Route) GetRoute() *RouteAction {
	if m != nil {
		return m.Route
	}
	return nil
}

type RouteMatch struct {
	// Types that are valid to be assigned to MatchSpecifier:
	//	*RouteMatch_MethodName
	//	*RouteMatch_ServiceName
	MatchSpecifier       isRouteMatch_MatchSpecifier `protobuf_oneof:"match_specifier"`
	Invert               bool                        `protobuf:"varint,3,opt,name=invert,proto3" json:"invert,omitempty"`
	Headers              []*v3alpha.HeaderMatcher    `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *RouteMatch) Reset()         { *m = RouteMatch{} }
func (m *RouteMatch) String() string { return proto.CompactTextString(m) }
func (*RouteMatch) ProtoMessage()    {}
func (*RouteMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_35b85368996c9dcc, []int{2}
}

func (m *RouteMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteMatch.Unmarshal(m, b)
}
func (m *RouteMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteMatch.Marshal(b, m, deterministic)
}
func (m *RouteMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteMatch.Merge(m, src)
}
func (m *RouteMatch) XXX_Size() int {
	return xxx_messageInfo_RouteMatch.Size(m)
}
func (m *RouteMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteMatch.DiscardUnknown(m)
}

var xxx_messageInfo_RouteMatch proto.InternalMessageInfo

type isRouteMatch_MatchSpecifier interface {
	isRouteMatch_MatchSpecifier()
}

type RouteMatch_MethodName struct {
	MethodName string `protobuf:"bytes,1,opt,name=method_name,json=methodName,proto3,oneof"`
}

type RouteMatch_ServiceName struct {
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3,oneof"`
}

func (*RouteMatch_MethodName) isRouteMatch_MatchSpecifier() {}

func (*RouteMatch_ServiceName) isRouteMatch_MatchSpecifier() {}

func (m *RouteMatch) GetMatchSpecifier() isRouteMatch_MatchSpecifier {
	if m != nil {
		return m.MatchSpecifier
	}
	return nil
}

func (m *RouteMatch) GetMethodName() string {
	if x, ok := m.GetMatchSpecifier().(*RouteMatch_MethodName); ok {
		return x.MethodName
	}
	return ""
}

func (m *RouteMatch) GetServiceName() string {
	if x, ok := m.GetMatchSpecifier().(*RouteMatch_ServiceName); ok {
		return x.ServiceName
	}
	return ""
}

func (m *RouteMatch) GetInvert() bool {
	if m != nil {
		return m.Invert
	}
	return false
}

func (m *RouteMatch) GetHeaders() []*v3alpha.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RouteMatch) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RouteMatch_MethodName)(nil),
		(*RouteMatch_ServiceName)(nil),
	}
}

type RouteAction struct {
	// Types that are valid to be assigned to ClusterSpecifier:
	//	*RouteAction_Cluster
	//	*RouteAction_WeightedClusters
	//	*RouteAction_ClusterHeader
	ClusterSpecifier     isRouteAction_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
	MetadataMatch        *v3alpha1.Metadata             `protobuf:"bytes,3,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	RateLimits           []*v3alpha.RateLimit           `protobuf:"bytes,4,rep,name=rate_limits,json=rateLimits,proto3" json:"rate_limits,omitempty"`
	StripServiceName     bool                           `protobuf:"varint,5,opt,name=strip_service_name,json=stripServiceName,proto3" json:"strip_service_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *RouteAction) Reset()         { *m = RouteAction{} }
func (m *RouteAction) String() string { return proto.CompactTextString(m) }
func (*RouteAction) ProtoMessage()    {}
func (*RouteAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_35b85368996c9dcc, []int{3}
}

func (m *RouteAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteAction.Unmarshal(m, b)
}
func (m *RouteAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteAction.Marshal(b, m, deterministic)
}
func (m *RouteAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteAction.Merge(m, src)
}
func (m *RouteAction) XXX_Size() int {
	return xxx_messageInfo_RouteAction.Size(m)
}
func (m *RouteAction) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteAction.DiscardUnknown(m)
}

var xxx_messageInfo_RouteAction proto.InternalMessageInfo

type isRouteAction_ClusterSpecifier interface {
	isRouteAction_ClusterSpecifier()
}

type RouteAction_Cluster struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3,oneof"`
}

type RouteAction_WeightedClusters struct {
	WeightedClusters *WeightedCluster `protobuf:"bytes,2,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

type RouteAction_ClusterHeader struct {
	ClusterHeader string `protobuf:"bytes,6,opt,name=cluster_header,json=clusterHeader,proto3,oneof"`
}

func (*RouteAction_Cluster) isRouteAction_ClusterSpecifier() {}

func (*RouteAction_WeightedClusters) isRouteAction_ClusterSpecifier() {}

func (*RouteAction_ClusterHeader) isRouteAction_ClusterSpecifier() {}

func (m *RouteAction) GetClusterSpecifier() isRouteAction_ClusterSpecifier {
	if m != nil {
		return m.ClusterSpecifier
	}
	return nil
}

func (m *RouteAction) GetCluster() string {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *RouteAction) GetWeightedClusters() *WeightedCluster {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

func (m *RouteAction) GetClusterHeader() string {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_ClusterHeader); ok {
		return x.ClusterHeader
	}
	return ""
}

func (m *RouteAction) GetMetadataMatch() *v3alpha1.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func (m *RouteAction) GetRateLimits() []*v3alpha.RateLimit {
	if m != nil {
		return m.RateLimits
	}
	return nil
}

func (m *RouteAction) GetStripServiceName() bool {
	if m != nil {
		return m.StripServiceName
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RouteAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RouteAction_Cluster)(nil),
		(*RouteAction_WeightedClusters)(nil),
		(*RouteAction_ClusterHeader)(nil),
	}
}

type WeightedCluster struct {
	Clusters             []*WeightedCluster_ClusterWeight `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *WeightedCluster) Reset()         { *m = WeightedCluster{} }
func (m *WeightedCluster) String() string { return proto.CompactTextString(m) }
func (*WeightedCluster) ProtoMessage()    {}
func (*WeightedCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_35b85368996c9dcc, []int{4}
}

func (m *WeightedCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeightedCluster.Unmarshal(m, b)
}
func (m *WeightedCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeightedCluster.Marshal(b, m, deterministic)
}
func (m *WeightedCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedCluster.Merge(m, src)
}
func (m *WeightedCluster) XXX_Size() int {
	return xxx_messageInfo_WeightedCluster.Size(m)
}
func (m *WeightedCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedCluster.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedCluster proto.InternalMessageInfo

func (m *WeightedCluster) GetClusters() []*WeightedCluster_ClusterWeight {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type WeightedCluster_ClusterWeight struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Weight               *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=weight,proto3" json:"weight,omitempty"`
	MetadataMatch        *v3alpha1.Metadata    `protobuf:"bytes,3,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *WeightedCluster_ClusterWeight) Reset()         { *m = WeightedCluster_ClusterWeight{} }
func (m *WeightedCluster_ClusterWeight) String() string { return proto.CompactTextString(m) }
func (*WeightedCluster_ClusterWeight) ProtoMessage()    {}
func (*WeightedCluster_ClusterWeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_35b85368996c9dcc, []int{4, 0}
}

func (m *WeightedCluster_ClusterWeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeightedCluster_ClusterWeight.Unmarshal(m, b)
}
func (m *WeightedCluster_ClusterWeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeightedCluster_ClusterWeight.Marshal(b, m, deterministic)
}
func (m *WeightedCluster_ClusterWeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedCluster_ClusterWeight.Merge(m, src)
}
func (m *WeightedCluster_ClusterWeight) XXX_Size() int {
	return xxx_messageInfo_WeightedCluster_ClusterWeight.Size(m)
}
func (m *WeightedCluster_ClusterWeight) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedCluster_ClusterWeight.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedCluster_ClusterWeight proto.InternalMessageInfo

func (m *WeightedCluster_ClusterWeight) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WeightedCluster_ClusterWeight) GetWeight() *wrappers.UInt32Value {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (m *WeightedCluster_ClusterWeight) GetMetadataMatch() *v3alpha1.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.extensions.filters.network.thrift_proxy.v3alpha.RouteConfiguration")
	proto.RegisterType((*Route)(nil), "envoy.extensions.filters.network.thrift_proxy.v3alpha.Route")
	proto.RegisterType((*RouteMatch)(nil), "envoy.extensions.filters.network.thrift_proxy.v3alpha.RouteMatch")
	proto.RegisterType((*RouteAction)(nil), "envoy.extensions.filters.network.thrift_proxy.v3alpha.RouteAction")
	proto.RegisterType((*WeightedCluster)(nil), "envoy.extensions.filters.network.thrift_proxy.v3alpha.WeightedCluster")
	proto.RegisterType((*WeightedCluster_ClusterWeight)(nil), "envoy.extensions.filters.network.thrift_proxy.v3alpha.WeightedCluster.ClusterWeight")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/thrift_proxy/v3alpha/route.proto", fileDescriptor_35b85368996c9dcc)
}

var fileDescriptor_35b85368996c9dcc = []byte{
	// 773 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x4e, 0x13, 0x41,
	0x14, 0x66, 0xb7, 0xb4, 0xd4, 0x53, 0xf9, 0x9b, 0x0b, 0x6c, 0xaa, 0x21, 0x05, 0xd4, 0x54, 0x63,
	0x76, 0xb5, 0xc4, 0x18, 0x1b, 0xc0, 0xb0, 0x25, 0xa4, 0xfe, 0xa0, 0xb8, 0x20, 0x5e, 0x6e, 0xa6,
	0xed, 0xb4, 0x9d, 0xd8, 0xee, 0x6e, 0x66, 0xa7, 0x2d, 0xf8, 0x04, 0xc6, 0x0b, 0x4d, 0xbc, 0xf4,
	0x7d, 0xbc, 0xf2, 0x49, 0x7c, 0x03, 0x43, 0x62, 0x62, 0xe6, 0x67, 0xfb, 0x83, 0x4a, 0x62, 0xe1,
	0xaa, 0xd3, 0x33, 0x67, 0xbe, 0xf3, 0x9d, 0xef, 0x3b, 0x3b, 0x03, 0xdb, 0xc4, 0xef, 0x05, 0x27,
	0x36, 0x39, 0xe6, 0xc4, 0x8f, 0x68, 0xe0, 0x47, 0x76, 0x83, 0xb6, 0x39, 0x61, 0x91, 0xed, 0x13,
	0xde, 0x0f, 0xd8, 0x3b, 0x9b, 0xb7, 0x18, 0x6d, 0x70, 0x2f, 0x64, 0xc1, 0xf1, 0x89, 0xdd, 0x5b,
	0xc7, 0xed, 0xb0, 0x85, 0x6d, 0x16, 0x74, 0x39, 0xb1, 0x42, 0x16, 0xf0, 0x00, 0x3d, 0x94, 0x10,
	0xd6, 0x10, 0xc2, 0xd2, 0x10, 0x96, 0x86, 0xb0, 0x46, 0x21, 0x2c, 0x0d, 0x91, 0xbb, 0xa9, 0x2a,
	0xd7, 0x02, 0xbf, 0x41, 0x9b, 0x76, 0x2d, 0x60, 0x64, 0x80, 0x5e, 0xc5, 0x91, 0x06, 0xcf, 0xdd,
	0x1e, 0xcb, 0x92, 0x65, 0xff, 0x46, 0x22, 0xb7, 0xdc, 0x0c, 0x82, 0x66, 0x9b, 0xd8, 0xf2, 0x5f,
	0xb5, 0xdb, 0xb0, 0xfb, 0x0c, 0x87, 0xa1, 0x20, 0xa1, 0xf6, 0x57, 0xba, 0xf5, 0x10, 0xdb, 0xd8,
	0xf7, 0x03, 0x8e, 0xb9, 0xec, 0xb3, 0x47, 0x98, 0x60, 0x4b, 0xfd, 0xa6, 0x4e, 0xb9, 0xd6, 0xc3,
	0x6d, 0x5a, 0xc7, 0xa2, 0x80, 0x5e, 0xa8, 0x8d, 0xd5, 0xef, 0x06, 0x20, 0x57, 0xd4, 0x2a, 0x4b,
	0x16, 0x5d, 0x26, 0x11, 0x10, 0x82, 0x69, 0x1f, 0x77, 0x48, 0xd6, 0xc8, 0x1b, 0x85, 0x2b, 0xae,
	0x5c, 0xa3, 0x43, 0x48, 0x49, 0x56, 0x51, 0xd6, 0xcc, 0x27, 0x0a, 0x99, 0xe2, 0x86, 0x35, 0x91,
	0x38, 0x96, 0x2c, 0xe7, 0x6a, 0xac, 0xd2, 0xf3, 0xaf, 0xdf, 0x3e, 0x2c, 0xef, 0xc2, 0x8e, 0xc2,
	0x52, 0x5a, 0x68, 0x9c, 0x7f, 0xc0, 0x14, 0x25, 0xcc, 0x03, 0xeb, 0x4f, 0xda, 0xab, 0x9f, 0x4c,
	0x48, 0xca, 0x30, 0xc2, 0x90, 0xec, 0x60, 0x5e, 0x6b, 0xc9, 0x0e, 0x32, 0xc5, 0xed, 0x8b, 0x70,
	0xdd, 0x13, 0x40, 0x4e, 0xfa, 0xd4, 0x49, 0x7e, 0x34, 0xcc, 0x05, 0xc3, 0x55, 0xc8, 0xa8, 0x0a,
	0x49, 0xd9, 0x43, 0xd6, 0x94, 0x25, 0x9c, 0x8b, 0x94, 0xd8, 0xae, 0x09, 0xfe, 0xa3, 0x35, 0x24,
	0x74, 0x69, 0x4b, 0xa8, 0xf3, 0x18, 0x1e, 0x4d, 0xa8, 0x8e, 0x10, 0x04, 0x86, 0x3d, 0xa0, 0x15,
	0xc8, 0x74, 0x08, 0x6f, 0x05, 0x75, 0x6f, 0xe8, 0x6e, 0x65, 0xca, 0x05, 0x15, 0x7c, 0x29, 0x5c,
	0x5e, 0x83, 0xab, 0x11, 0x61, 0x3d, 0x5a, 0x23, 0x2a, 0xc7, 0xd4, 0x39, 0x19, 0x1d, 0x95, 0x49,
	0x4b, 0x90, 0xa2, 0x7e, 0x8f, 0x30, 0x9e, 0x4d, 0xe4, 0x8d, 0x42, 0xda, 0xd5, 0xff, 0x50, 0x19,
	0x66, 0x5a, 0x04, 0xd7, 0x09, 0x8b, 0xb2, 0xd3, 0x72, 0x46, 0xee, 0x58, 0x63, 0xcc, 0xd5, 0x54,
	0xc7, 0x9d, 0x57, 0x64, 0xaa, 0x64, 0x46, 0x98, 0x1b, 0x9f, 0x2c, 0x95, 0x45, 0xcf, 0x5b, 0xb0,
	0x31, 0x61, 0xcf, 0xca, 0xad, 0x25, 0x98, 0x97, 0x2e, 0x79, 0x51, 0x48, 0x6a, 0xb4, 0x41, 0x09,
	0x43, 0x89, 0x9f, 0x8e, 0xb1, 0xfa, 0x2b, 0x01, 0x99, 0x11, 0xc5, 0xd1, 0x1a, 0xcc, 0xd4, 0xda,
	0xdd, 0x88, 0x13, 0xa6, 0xd4, 0x70, 0x66, 0x4e, 0x9d, 0x69, 0x66, 0xe6, 0x8d, 0xca, 0x94, 0x1b,
	0xef, 0xa0, 0x2e, 0x2c, 0xf6, 0x09, 0x6d, 0xb6, 0x38, 0xa9, 0x7b, 0x3a, 0x16, 0x69, 0xd7, 0x77,
	0x27, 0x74, 0xfd, 0xad, 0xc6, 0x2b, 0x2b, 0xb8, 0xca, 0x94, 0xbb, 0xd0, 0x1f, 0x0f, 0x45, 0xe8,
	0x3e, 0xcc, 0xe9, 0x6a, 0x9e, 0xd2, 0x26, 0x9b, 0x3a, 0x4b, 0x71, 0x56, 0x27, 0x28, 0x29, 0xd1,
	0x33, 0x98, 0xeb, 0x10, 0x8e, 0xeb, 0x98, 0x63, 0x4f, 0x8d, 0x7f, 0x42, 0xb2, 0x5c, 0x1b, 0xb7,
	0x41, 0x5c, 0x48, 0x03, 0x26, 0x7b, 0xfa, 0x80, 0x3b, 0x1b, 0x1f, 0x55, 0xb3, 0xb2, 0x0b, 0x19,
	0x86, 0x39, 0xf1, 0xda, 0xb4, 0x43, 0x79, 0xec, 0xe7, 0xad, 0xf3, 0xfc, 0x74, 0x31, 0x27, 0x2f,
	0x44, 0xb6, 0x0b, 0x2c, 0x5e, 0x46, 0xe8, 0x1e, 0xa0, 0x88, 0x33, 0x1a, 0x7a, 0x63, 0x63, 0x95,
	0x94, 0x73, 0xb3, 0x20, 0x77, 0x0e, 0x86, 0x93, 0x55, 0xda, 0x11, 0xe6, 0x3f, 0x81, 0xcd, 0x09,
	0xcd, 0xd7, 0xdf, 0x51, 0x16, 0x16, 0x63, 0xe5, 0xce, 0xf8, 0xff, 0x23, 0x01, 0xf3, 0x67, 0xb4,
	0x47, 0xef, 0x21, 0x3d, 0x70, 0xd5, 0x90, 0x6d, 0x1e, 0x5e, 0x8e, 0xab, 0x96, 0xfe, 0x55, 0x61,
	0xf9, 0x75, 0x7f, 0x31, 0xcc, 0xb4, 0xe1, 0x0e, 0xea, 0xe5, 0x3e, 0x9b, 0x30, 0x3b, 0x96, 0x85,
	0xae, 0x8f, 0x5e, 0xbd, 0x03, 0xaf, 0xf5, 0x1d, 0xbc, 0x09, 0x29, 0x35, 0x26, 0x7a, 0xfc, 0x6e,
	0x58, 0xea, 0x6d, 0xb0, 0xe2, 0xb7, 0xc1, 0x7a, 0xf3, 0xd4, 0xe7, 0xeb, 0xc5, 0x23, 0xdc, 0xee,
	0x12, 0x79, 0xf8, 0xae, 0x59, 0x30, 0x5c, 0x7d, 0xe8, 0x32, 0xe7, 0xa3, 0x74, 0x24, 0x9c, 0x7a,
	0x0d, 0xaf, 0xfe, 0xdf, 0xa9, 0x73, 0x55, 0x2a, 0x55, 0x04, 0x6e, 0x59, 0x3f, 0xde, 0x17, 0xc1,
	0x75, 0x0e, 0xa0, 0x4c, 0x03, 0xd5, 0x99, 0xca, 0x9c, 0xc8, 0x54, 0x47, 0x5d, 0xa0, 0xfb, 0x42,
	0xe0, 0x7d, 0xa3, 0x9a, 0x92, 0x4a, 0xaf, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x46, 0xde, 0x49,
	0xf7, 0x5d, 0x08, 0x00, 0x00,
}