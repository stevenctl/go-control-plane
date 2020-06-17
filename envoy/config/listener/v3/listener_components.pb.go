// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/listener/v3/listener_components.proto

package envoy_config_listener_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v31 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	v32 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type FilterChainMatch_ConnectionSourceType int32

const (
	FilterChainMatch_ANY                 FilterChainMatch_ConnectionSourceType = 0
	FilterChainMatch_SAME_IP_OR_LOOPBACK FilterChainMatch_ConnectionSourceType = 1
	FilterChainMatch_EXTERNAL            FilterChainMatch_ConnectionSourceType = 2
)

var FilterChainMatch_ConnectionSourceType_name = map[int32]string{
	0: "ANY",
	1: "SAME_IP_OR_LOOPBACK",
	2: "EXTERNAL",
}

var FilterChainMatch_ConnectionSourceType_value = map[string]int32{
	"ANY":                 0,
	"SAME_IP_OR_LOOPBACK": 1,
	"EXTERNAL":            2,
}

func (x FilterChainMatch_ConnectionSourceType) String() string {
	return proto.EnumName(FilterChainMatch_ConnectionSourceType_name, int32(x))
}

func (FilterChainMatch_ConnectionSourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_87f255d2eccc91b5, []int{1, 0}
}

type Filter struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*Filter_TypedConfig
	//	*Filter_HiddenEnvoyDeprecatedConfig
	ConfigType           isFilter_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f255d2eccc91b5, []int{0}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isFilter_ConfigType interface {
	isFilter_ConfigType()
}

type Filter_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,4,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

type Filter_HiddenEnvoyDeprecatedConfig struct {
	HiddenEnvoyDeprecatedConfig *_struct.Struct `protobuf:"bytes,2,opt,name=hidden_envoy_deprecated_config,json=hiddenEnvoyDeprecatedConfig,proto3,oneof"`
}

func (*Filter_TypedConfig) isFilter_ConfigType() {}

func (*Filter_HiddenEnvoyDeprecatedConfig) isFilter_ConfigType() {}

func (m *Filter) GetConfigType() isFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *Filter) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*Filter_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// Deprecated: Do not use.
func (m *Filter) GetHiddenEnvoyDeprecatedConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*Filter_HiddenEnvoyDeprecatedConfig); ok {
		return x.HiddenEnvoyDeprecatedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Filter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Filter_TypedConfig)(nil),
		(*Filter_HiddenEnvoyDeprecatedConfig)(nil),
	}
}

type FilterChainMatch struct {
	DestinationPort      *wrappers.UInt32Value                 `protobuf:"bytes,8,opt,name=destination_port,json=destinationPort,proto3" json:"destination_port,omitempty"`
	PrefixRanges         []*v3.CidrRange                       `protobuf:"bytes,3,rep,name=prefix_ranges,json=prefixRanges,proto3" json:"prefix_ranges,omitempty"`
	AddressSuffix        string                                `protobuf:"bytes,4,opt,name=address_suffix,json=addressSuffix,proto3" json:"address_suffix,omitempty"`
	SuffixLen            *wrappers.UInt32Value                 `protobuf:"bytes,5,opt,name=suffix_len,json=suffixLen,proto3" json:"suffix_len,omitempty"`
	SourceType           FilterChainMatch_ConnectionSourceType `protobuf:"varint,12,opt,name=source_type,json=sourceType,proto3,enum=envoy.config.listener.v3.FilterChainMatch_ConnectionSourceType" json:"source_type,omitempty"`
	SourcePrefixRanges   []*v3.CidrRange                       `protobuf:"bytes,6,rep,name=source_prefix_ranges,json=sourcePrefixRanges,proto3" json:"source_prefix_ranges,omitempty"`
	SourcePorts          []uint32                              `protobuf:"varint,7,rep,packed,name=source_ports,json=sourcePorts,proto3" json:"source_ports,omitempty"`
	ServerNames          []string                              `protobuf:"bytes,11,rep,name=server_names,json=serverNames,proto3" json:"server_names,omitempty"`
	TransportProtocol    string                                `protobuf:"bytes,9,opt,name=transport_protocol,json=transportProtocol,proto3" json:"transport_protocol,omitempty"`
	ApplicationProtocols []string                              `protobuf:"bytes,10,rep,name=application_protocols,json=applicationProtocols,proto3" json:"application_protocols,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *FilterChainMatch) Reset()         { *m = FilterChainMatch{} }
func (m *FilterChainMatch) String() string { return proto.CompactTextString(m) }
func (*FilterChainMatch) ProtoMessage()    {}
func (*FilterChainMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f255d2eccc91b5, []int{1}
}

func (m *FilterChainMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterChainMatch.Unmarshal(m, b)
}
func (m *FilterChainMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterChainMatch.Marshal(b, m, deterministic)
}
func (m *FilterChainMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterChainMatch.Merge(m, src)
}
func (m *FilterChainMatch) XXX_Size() int {
	return xxx_messageInfo_FilterChainMatch.Size(m)
}
func (m *FilterChainMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterChainMatch.DiscardUnknown(m)
}

var xxx_messageInfo_FilterChainMatch proto.InternalMessageInfo

func (m *FilterChainMatch) GetDestinationPort() *wrappers.UInt32Value {
	if m != nil {
		return m.DestinationPort
	}
	return nil
}

func (m *FilterChainMatch) GetPrefixRanges() []*v3.CidrRange {
	if m != nil {
		return m.PrefixRanges
	}
	return nil
}

func (m *FilterChainMatch) GetAddressSuffix() string {
	if m != nil {
		return m.AddressSuffix
	}
	return ""
}

func (m *FilterChainMatch) GetSuffixLen() *wrappers.UInt32Value {
	if m != nil {
		return m.SuffixLen
	}
	return nil
}

func (m *FilterChainMatch) GetSourceType() FilterChainMatch_ConnectionSourceType {
	if m != nil {
		return m.SourceType
	}
	return FilterChainMatch_ANY
}

func (m *FilterChainMatch) GetSourcePrefixRanges() []*v3.CidrRange {
	if m != nil {
		return m.SourcePrefixRanges
	}
	return nil
}

func (m *FilterChainMatch) GetSourcePorts() []uint32 {
	if m != nil {
		return m.SourcePorts
	}
	return nil
}

func (m *FilterChainMatch) GetServerNames() []string {
	if m != nil {
		return m.ServerNames
	}
	return nil
}

func (m *FilterChainMatch) GetTransportProtocol() string {
	if m != nil {
		return m.TransportProtocol
	}
	return ""
}

func (m *FilterChainMatch) GetApplicationProtocols() []string {
	if m != nil {
		return m.ApplicationProtocols
	}
	return nil
}

type FilterChain struct {
	FilterChainMatch                *FilterChainMatch         `protobuf:"bytes,1,opt,name=filter_chain_match,json=filterChainMatch,proto3" json:"filter_chain_match,omitempty"`
	Filters                         []*Filter                 `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
	UseProxyProto                   *wrappers.BoolValue       `protobuf:"bytes,4,opt,name=use_proxy_proto,json=useProxyProto,proto3" json:"use_proxy_proto,omitempty"`
	Metadata                        *v3.Metadata              `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	TransportSocket                 *v3.TransportSocket       `protobuf:"bytes,6,opt,name=transport_socket,json=transportSocket,proto3" json:"transport_socket,omitempty"`
	Name                            string                    `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	HiddenEnvoyDeprecatedTlsContext *v31.DownstreamTlsContext `protobuf:"bytes,2,opt,name=hidden_envoy_deprecated_tls_context,json=hiddenEnvoyDeprecatedTlsContext,proto3" json:"hidden_envoy_deprecated_tls_context,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral            struct{}                  `json:"-"`
	XXX_unrecognized                []byte                    `json:"-"`
	XXX_sizecache                   int32                     `json:"-"`
}

func (m *FilterChain) Reset()         { *m = FilterChain{} }
func (m *FilterChain) String() string { return proto.CompactTextString(m) }
func (*FilterChain) ProtoMessage()    {}
func (*FilterChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f255d2eccc91b5, []int{2}
}

func (m *FilterChain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterChain.Unmarshal(m, b)
}
func (m *FilterChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterChain.Marshal(b, m, deterministic)
}
func (m *FilterChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterChain.Merge(m, src)
}
func (m *FilterChain) XXX_Size() int {
	return xxx_messageInfo_FilterChain.Size(m)
}
func (m *FilterChain) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterChain.DiscardUnknown(m)
}

var xxx_messageInfo_FilterChain proto.InternalMessageInfo

func (m *FilterChain) GetFilterChainMatch() *FilterChainMatch {
	if m != nil {
		return m.FilterChainMatch
	}
	return nil
}

func (m *FilterChain) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *FilterChain) GetUseProxyProto() *wrappers.BoolValue {
	if m != nil {
		return m.UseProxyProto
	}
	return nil
}

func (m *FilterChain) GetMetadata() *v3.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *FilterChain) GetTransportSocket() *v3.TransportSocket {
	if m != nil {
		return m.TransportSocket
	}
	return nil
}

func (m *FilterChain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Deprecated: Do not use.
func (m *FilterChain) GetHiddenEnvoyDeprecatedTlsContext() *v31.DownstreamTlsContext {
	if m != nil {
		return m.HiddenEnvoyDeprecatedTlsContext
	}
	return nil
}

type ListenerFilterChainMatchPredicate struct {
	// Types that are valid to be assigned to Rule:
	//	*ListenerFilterChainMatchPredicate_OrMatch
	//	*ListenerFilterChainMatchPredicate_AndMatch
	//	*ListenerFilterChainMatchPredicate_NotMatch
	//	*ListenerFilterChainMatchPredicate_AnyMatch
	//	*ListenerFilterChainMatchPredicate_DestinationPortRange
	Rule                 isListenerFilterChainMatchPredicate_Rule `protobuf_oneof:"rule"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *ListenerFilterChainMatchPredicate) Reset()         { *m = ListenerFilterChainMatchPredicate{} }
func (m *ListenerFilterChainMatchPredicate) String() string { return proto.CompactTextString(m) }
func (*ListenerFilterChainMatchPredicate) ProtoMessage()    {}
func (*ListenerFilterChainMatchPredicate) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f255d2eccc91b5, []int{3}
}

func (m *ListenerFilterChainMatchPredicate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerFilterChainMatchPredicate.Unmarshal(m, b)
}
func (m *ListenerFilterChainMatchPredicate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerFilterChainMatchPredicate.Marshal(b, m, deterministic)
}
func (m *ListenerFilterChainMatchPredicate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerFilterChainMatchPredicate.Merge(m, src)
}
func (m *ListenerFilterChainMatchPredicate) XXX_Size() int {
	return xxx_messageInfo_ListenerFilterChainMatchPredicate.Size(m)
}
func (m *ListenerFilterChainMatchPredicate) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerFilterChainMatchPredicate.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerFilterChainMatchPredicate proto.InternalMessageInfo

type isListenerFilterChainMatchPredicate_Rule interface {
	isListenerFilterChainMatchPredicate_Rule()
}

type ListenerFilterChainMatchPredicate_OrMatch struct {
	OrMatch *ListenerFilterChainMatchPredicate_MatchSet `protobuf:"bytes,1,opt,name=or_match,json=orMatch,proto3,oneof"`
}

type ListenerFilterChainMatchPredicate_AndMatch struct {
	AndMatch *ListenerFilterChainMatchPredicate_MatchSet `protobuf:"bytes,2,opt,name=and_match,json=andMatch,proto3,oneof"`
}

type ListenerFilterChainMatchPredicate_NotMatch struct {
	NotMatch *ListenerFilterChainMatchPredicate `protobuf:"bytes,3,opt,name=not_match,json=notMatch,proto3,oneof"`
}

type ListenerFilterChainMatchPredicate_AnyMatch struct {
	AnyMatch bool `protobuf:"varint,4,opt,name=any_match,json=anyMatch,proto3,oneof"`
}

type ListenerFilterChainMatchPredicate_DestinationPortRange struct {
	DestinationPortRange *v32.Int32Range `protobuf:"bytes,5,opt,name=destination_port_range,json=destinationPortRange,proto3,oneof"`
}

func (*ListenerFilterChainMatchPredicate_OrMatch) isListenerFilterChainMatchPredicate_Rule() {}

func (*ListenerFilterChainMatchPredicate_AndMatch) isListenerFilterChainMatchPredicate_Rule() {}

func (*ListenerFilterChainMatchPredicate_NotMatch) isListenerFilterChainMatchPredicate_Rule() {}

func (*ListenerFilterChainMatchPredicate_AnyMatch) isListenerFilterChainMatchPredicate_Rule() {}

func (*ListenerFilterChainMatchPredicate_DestinationPortRange) isListenerFilterChainMatchPredicate_Rule() {
}

func (m *ListenerFilterChainMatchPredicate) GetRule() isListenerFilterChainMatchPredicate_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *ListenerFilterChainMatchPredicate) GetOrMatch() *ListenerFilterChainMatchPredicate_MatchSet {
	if x, ok := m.GetRule().(*ListenerFilterChainMatchPredicate_OrMatch); ok {
		return x.OrMatch
	}
	return nil
}

func (m *ListenerFilterChainMatchPredicate) GetAndMatch() *ListenerFilterChainMatchPredicate_MatchSet {
	if x, ok := m.GetRule().(*ListenerFilterChainMatchPredicate_AndMatch); ok {
		return x.AndMatch
	}
	return nil
}

func (m *ListenerFilterChainMatchPredicate) GetNotMatch() *ListenerFilterChainMatchPredicate {
	if x, ok := m.GetRule().(*ListenerFilterChainMatchPredicate_NotMatch); ok {
		return x.NotMatch
	}
	return nil
}

func (m *ListenerFilterChainMatchPredicate) GetAnyMatch() bool {
	if x, ok := m.GetRule().(*ListenerFilterChainMatchPredicate_AnyMatch); ok {
		return x.AnyMatch
	}
	return false
}

func (m *ListenerFilterChainMatchPredicate) GetDestinationPortRange() *v32.Int32Range {
	if x, ok := m.GetRule().(*ListenerFilterChainMatchPredicate_DestinationPortRange); ok {
		return x.DestinationPortRange
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ListenerFilterChainMatchPredicate) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ListenerFilterChainMatchPredicate_OrMatch)(nil),
		(*ListenerFilterChainMatchPredicate_AndMatch)(nil),
		(*ListenerFilterChainMatchPredicate_NotMatch)(nil),
		(*ListenerFilterChainMatchPredicate_AnyMatch)(nil),
		(*ListenerFilterChainMatchPredicate_DestinationPortRange)(nil),
	}
}

type ListenerFilterChainMatchPredicate_MatchSet struct {
	Rules                []*ListenerFilterChainMatchPredicate `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *ListenerFilterChainMatchPredicate_MatchSet) Reset() {
	*m = ListenerFilterChainMatchPredicate_MatchSet{}
}
func (m *ListenerFilterChainMatchPredicate_MatchSet) String() string {
	return proto.CompactTextString(m)
}
func (*ListenerFilterChainMatchPredicate_MatchSet) ProtoMessage() {}
func (*ListenerFilterChainMatchPredicate_MatchSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f255d2eccc91b5, []int{3, 0}
}

func (m *ListenerFilterChainMatchPredicate_MatchSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerFilterChainMatchPredicate_MatchSet.Unmarshal(m, b)
}
func (m *ListenerFilterChainMatchPredicate_MatchSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerFilterChainMatchPredicate_MatchSet.Marshal(b, m, deterministic)
}
func (m *ListenerFilterChainMatchPredicate_MatchSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerFilterChainMatchPredicate_MatchSet.Merge(m, src)
}
func (m *ListenerFilterChainMatchPredicate_MatchSet) XXX_Size() int {
	return xxx_messageInfo_ListenerFilterChainMatchPredicate_MatchSet.Size(m)
}
func (m *ListenerFilterChainMatchPredicate_MatchSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerFilterChainMatchPredicate_MatchSet.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerFilterChainMatchPredicate_MatchSet proto.InternalMessageInfo

func (m *ListenerFilterChainMatchPredicate_MatchSet) GetRules() []*ListenerFilterChainMatchPredicate {
	if m != nil {
		return m.Rules
	}
	return nil
}

type ListenerFilter struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*ListenerFilter_TypedConfig
	//	*ListenerFilter_HiddenEnvoyDeprecatedConfig
	ConfigType           isListenerFilter_ConfigType        `protobuf_oneof:"config_type"`
	FilterDisabled       *ListenerFilterChainMatchPredicate `protobuf:"bytes,4,opt,name=filter_disabled,json=filterDisabled,proto3" json:"filter_disabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *ListenerFilter) Reset()         { *m = ListenerFilter{} }
func (m *ListenerFilter) String() string { return proto.CompactTextString(m) }
func (*ListenerFilter) ProtoMessage()    {}
func (*ListenerFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_87f255d2eccc91b5, []int{4}
}

func (m *ListenerFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerFilter.Unmarshal(m, b)
}
func (m *ListenerFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerFilter.Marshal(b, m, deterministic)
}
func (m *ListenerFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerFilter.Merge(m, src)
}
func (m *ListenerFilter) XXX_Size() int {
	return xxx_messageInfo_ListenerFilter.Size(m)
}
func (m *ListenerFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerFilter proto.InternalMessageInfo

func (m *ListenerFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isListenerFilter_ConfigType interface {
	isListenerFilter_ConfigType()
}

type ListenerFilter_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

type ListenerFilter_HiddenEnvoyDeprecatedConfig struct {
	HiddenEnvoyDeprecatedConfig *_struct.Struct `protobuf:"bytes,2,opt,name=hidden_envoy_deprecated_config,json=hiddenEnvoyDeprecatedConfig,proto3,oneof"`
}

func (*ListenerFilter_TypedConfig) isListenerFilter_ConfigType() {}

func (*ListenerFilter_HiddenEnvoyDeprecatedConfig) isListenerFilter_ConfigType() {}

func (m *ListenerFilter) GetConfigType() isListenerFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *ListenerFilter) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*ListenerFilter_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// Deprecated: Do not use.
func (m *ListenerFilter) GetHiddenEnvoyDeprecatedConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*ListenerFilter_HiddenEnvoyDeprecatedConfig); ok {
		return x.HiddenEnvoyDeprecatedConfig
	}
	return nil
}

func (m *ListenerFilter) GetFilterDisabled() *ListenerFilterChainMatchPredicate {
	if m != nil {
		return m.FilterDisabled
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ListenerFilter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ListenerFilter_TypedConfig)(nil),
		(*ListenerFilter_HiddenEnvoyDeprecatedConfig)(nil),
	}
}

func init() {
	proto.RegisterEnum("envoy.config.listener.v3.FilterChainMatch_ConnectionSourceType", FilterChainMatch_ConnectionSourceType_name, FilterChainMatch_ConnectionSourceType_value)
	proto.RegisterType((*Filter)(nil), "envoy.config.listener.v3.Filter")
	proto.RegisterType((*FilterChainMatch)(nil), "envoy.config.listener.v3.FilterChainMatch")
	proto.RegisterType((*FilterChain)(nil), "envoy.config.listener.v3.FilterChain")
	proto.RegisterType((*ListenerFilterChainMatchPredicate)(nil), "envoy.config.listener.v3.ListenerFilterChainMatchPredicate")
	proto.RegisterType((*ListenerFilterChainMatchPredicate_MatchSet)(nil), "envoy.config.listener.v3.ListenerFilterChainMatchPredicate.MatchSet")
	proto.RegisterType((*ListenerFilter)(nil), "envoy.config.listener.v3.ListenerFilter")
}

func init() {
	proto.RegisterFile("envoy/config/listener/v3/listener_components.proto", fileDescriptor_87f255d2eccc91b5)
}

var fileDescriptor_87f255d2eccc91b5 = []byte{
	// 1236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0x4d, 0x6f, 0x1b, 0x45,
	0x18, 0xce, 0xda, 0x8e, 0xb3, 0x1e, 0xe7, 0xc3, 0x0c, 0x81, 0x6c, 0xd3, 0x92, 0x3a, 0x2e, 0x8d,
	0xac, 0xa2, 0xae, 0x25, 0xfb, 0x80, 0xea, 0x4a, 0x40, 0xd6, 0x49, 0x49, 0xdb, 0xb4, 0x75, 0x37,
	0x01, 0x15, 0x38, 0xac, 0x26, 0xbb, 0xe3, 0x74, 0xcb, 0x66, 0x66, 0x35, 0x33, 0x76, 0xe3, 0x1b,
	0xe2, 0x84, 0xb8, 0xd1, 0x23, 0xbf, 0x82, 0x33, 0x47, 0x24, 0x04, 0x57, 0xfe, 0x08, 0x67, 0x94,
	0x03, 0x45, 0xf3, 0xb1, 0x71, 0xe2, 0x24, 0x24, 0x10, 0x0e, 0xdc, 0x3c, 0xef, 0xc7, 0xf3, 0xbe,
	0xf3, 0xce, 0x33, 0xcf, 0xac, 0x41, 0x13, 0x93, 0x01, 0x1d, 0x36, 0x42, 0x4a, 0x7a, 0xf1, 0x6e,
	0x23, 0x89, 0xb9, 0xc0, 0x04, 0xb3, 0xc6, 0xa0, 0x75, 0xf8, 0x3b, 0x08, 0xe9, 0x5e, 0x4a, 0x09,
	0x26, 0x82, 0xbb, 0x29, 0xa3, 0x82, 0x42, 0x47, 0xe5, 0xb8, 0x3a, 0xc7, 0xcd, 0xe2, 0xdc, 0x41,
	0x6b, 0xb1, 0x76, 0x0c, 0x2d, 0xa4, 0x0c, 0x4b, 0x24, 0x14, 0x45, 0x0c, 0x73, 0x93, 0xbd, 0x78,
	0xfd, 0xd4, 0x98, 0x1d, 0xc4, 0xb1, 0x09, 0x68, 0xe9, 0x00, 0xbc, 0x2f, 0x30, 0xe1, 0x31, 0x25,
	0xbc, 0x21, 0x18, 0x22, 0x3c, 0xa5, 0x4c, 0x04, 0x9c, 0x86, 0x5f, 0x62, 0xc1, 0x1b, 0x22, 0xe1,
	0x32, 0x4b, 0x24, 0x19, 0xea, 0x15, 0x9d, 0x24, 0x86, 0xa9, 0x82, 0x63, 0x88, 0xec, 0x66, 0x78,
	0x57, 0x76, 0x29, 0xdd, 0x4d, 0x70, 0x43, 0xad, 0x76, 0xfa, 0xbd, 0x06, 0x22, 0x43, 0xe3, 0xba,
	0x36, 0xee, 0xe2, 0x82, 0xf5, 0x43, 0x61, 0xbc, 0x4b, 0xe3, 0xde, 0x97, 0x0c, 0xa5, 0x29, 0x66,
	0x59, 0xcd, 0x77, 0xfa, 0x51, 0x8a, 0x1a, 0x88, 0x10, 0x2a, 0x90, 0x50, 0x8d, 0x72, 0x81, 0x44,
	0x3f, 0x73, 0x2f, 0x9f, 0x70, 0x0f, 0x30, 0x93, 0x1b, 0x8a, 0xc9, 0xae, 0x09, 0x59, 0x18, 0xa0,
	0x24, 0x8e, 0x90, 0xc0, 0x8d, 0xec, 0x87, 0x76, 0xd4, 0xfe, 0xb4, 0x40, 0xf1, 0x5e, 0x9c, 0x08,
	0xcc, 0xe0, 0x55, 0x50, 0x20, 0x68, 0x0f, 0x3b, 0x56, 0xd5, 0xaa, 0x97, 0xbc, 0xa9, 0x03, 0xaf,
	0xc0, 0x72, 0x55, 0xcb, 0x57, 0x46, 0x78, 0x07, 0x4c, 0xcb, 0x2d, 0x47, 0x81, 0x1e, 0xa7, 0x53,
	0xa8, 0x5a, 0xf5, 0x72, 0x73, 0xde, 0xd5, 0x9d, 0xbb, 0x59, 0xe7, 0xee, 0x2a, 0x19, 0x6e, 0x4c,
	0xf8, 0x65, 0x15, 0xdb, 0x51, 0xa1, 0x70, 0x07, 0x2c, 0x3d, 0x8f, 0xa3, 0x08, 0x93, 0x40, 0x8d,
	0x2e, 0x88, 0x70, 0xca, 0x70, 0x88, 0xc4, 0x08, 0x2c, 0xa7, 0xc0, 0x16, 0x4e, 0x80, 0x6d, 0xa9,
	0x21, 0x79, 0x39, 0xc7, 0xda, 0x98, 0xf0, 0xaf, 0x6a, 0x90, 0x75, 0x89, 0xb1, 0x76, 0x08, 0xa1,
	0x6b, 0xb4, 0x6f, 0x7c, 0xff, 0xf3, 0x37, 0x4b, 0x4b, 0xe0, 0x9a, 0x26, 0x0c, 0x4a, 0x63, 0x77,
	0xd0, 0x1c, 0x11, 0x46, 0x6f, 0xd0, 0x9b, 0x01, 0x65, 0x5d, 0x30, 0x90, 0xed, 0x3d, 0x28, 0xd8,
	0xf9, 0x4a, 0xa1, 0xf6, 0x6d, 0x11, 0x54, 0xb4, 0xbf, 0xf3, 0x1c, 0xc5, 0xe4, 0x11, 0x12, 0xe1,
	0x73, 0xb8, 0x0d, 0x2a, 0x11, 0xe6, 0x22, 0x26, 0x6a, 0x9e, 0x81, 0x64, 0x84, 0x63, 0xab, 0x26,
	0xaf, 0x9d, 0x68, 0xf2, 0x93, 0xfb, 0x44, 0xb4, 0x9a, 0x9f, 0xa2, 0xa4, 0x8f, 0xbd, 0xf2, 0x81,
	0x67, 0xdf, 0x2a, 0x3a, 0xaf, 0x5f, 0xe7, 0xeb, 0x96, 0x3f, 0x77, 0x04, 0xa2, 0x4b, 0x99, 0x80,
	0x6b, 0x60, 0x26, 0x65, 0xb8, 0x17, 0xef, 0x07, 0x8a, 0x35, 0xdc, 0xc9, 0x57, 0xf3, 0xf5, 0x72,
	0xf3, 0xba, 0x7b, 0x8c, 0xe6, 0x92, 0xa8, 0xee, 0xa0, 0xe5, 0x76, 0xe2, 0x88, 0xf9, 0x32, 0xce,
	0x9f, 0xd6, 0x59, 0x6a, 0xc1, 0xe1, 0x4d, 0x30, 0x6b, 0x78, 0x1e, 0xf0, 0x7e, 0xaf, 0x17, 0xef,
	0xab, 0xb3, 0x28, 0xf9, 0x33, 0xc6, 0xba, 0xa5, 0x8c, 0xf0, 0x2e, 0x00, 0xda, 0x1d, 0x24, 0x98,
	0x38, 0x93, 0xe7, 0x37, 0xef, 0x97, 0x74, 0xfc, 0x26, 0x26, 0xf0, 0x05, 0x28, 0x73, 0xda, 0x67,
	0x21, 0x56, 0x93, 0x72, 0xa6, 0xab, 0x56, 0x7d, 0xb6, 0xf9, 0xa1, 0x7b, 0xd6, 0x75, 0x74, 0xc7,
	0x07, 0xe8, 0x76, 0x28, 0x21, 0x38, 0x94, 0x3b, 0xdf, 0x52, 0x38, 0xdb, 0xc3, 0x14, 0x7b, 0xf6,
	0x81, 0x37, 0xf9, 0xb5, 0x95, 0xab, 0x58, 0x3e, 0xe0, 0x87, 0x56, 0xf8, 0x14, 0xcc, 0x9b, 0x5a,
	0xc7, 0x87, 0x53, 0xbc, 0xd8, 0x70, 0xa0, 0x4e, 0xee, 0x1e, 0x1d, 0x51, 0x0b, 0x4c, 0x67, 0x90,
	0x94, 0x09, 0xee, 0x4c, 0x55, 0xf3, 0xf5, 0x19, 0xaf, 0x72, 0xe0, 0xcd, 0xbc, 0xb2, 0x40, 0x6d,
	0x74, 0x42, 0x66, 0x93, 0xf2, 0x70, 0x38, 0x5c, 0x06, 0xd3, 0x1c, 0xb3, 0x01, 0x66, 0x81, 0x24,
	0x3c, 0x77, 0xca, 0xd5, 0x7c, 0xbd, 0xe4, 0x97, 0xb5, 0xed, 0xb1, 0x34, 0xc1, 0xdb, 0x00, 0x8e,
	0x14, 0x42, 0xcd, 0x30, 0xa4, 0x89, 0x53, 0x52, 0xe3, 0x7f, 0xe3, 0xd0, 0xd3, 0x35, 0x0e, 0xd8,
	0x02, 0x6f, 0xa1, 0x34, 0x4d, 0xe2, 0xd0, 0xb0, 0xc8, 0xd8, 0xb9, 0x03, 0x14, 0xf4, 0xfc, 0x11,
	0x67, 0x96, 0xc3, 0x6b, 0xf7, 0xc0, 0xfc, 0x69, 0xc3, 0x83, 0x53, 0x20, 0xbf, 0xfa, 0xf8, 0xb3,
	0xca, 0x04, 0x5c, 0x00, 0x6f, 0x6e, 0xad, 0x3e, 0x5a, 0x0f, 0xee, 0x77, 0x83, 0x27, 0x7e, 0xb0,
	0xf9, 0xe4, 0x49, 0xd7, 0x5b, 0xed, 0x3c, 0xac, 0x58, 0x70, 0x1a, 0xd8, 0xeb, 0xcf, 0xb6, 0xd7,
	0xfd, 0xc7, 0xab, 0x9b, 0x95, 0x5c, 0xfb, 0xb6, 0xbc, 0x11, 0x75, 0xb0, 0xf2, 0x77, 0x37, 0x62,
	0x74, 0x60, 0x0f, 0x0a, 0xb6, 0x55, 0xc9, 0xd5, 0x7e, 0x29, 0x80, 0xf2, 0x11, 0x17, 0x7c, 0x06,
	0x60, 0x4f, 0x2d, 0x83, 0x50, 0xae, 0x83, 0x3d, 0x19, 0xab, 0x04, 0xa2, 0xdc, 0xbc, 0x75, 0x71,
	0x3a, 0xf8, 0x95, 0xde, 0xf8, 0x0d, 0x6b, 0x83, 0x29, 0x6d, 0xcb, 0x6e, 0x41, 0xf5, 0x3c, 0x38,
	0x3f, 0x4b, 0x80, 0x1e, 0x98, 0xeb, 0x73, 0x49, 0x17, 0xba, 0x3f, 0xd4, 0x53, 0x35, 0x72, 0xb4,
	0x78, 0x82, 0xdf, 0x1e, 0xa5, 0x89, 0x66, 0xf7, 0x4c, 0x9f, 0xe3, 0xae, 0xcc, 0x50, 0xa3, 0x86,
	0x6d, 0x60, 0xef, 0x61, 0x81, 0x22, 0x24, 0x90, 0xb9, 0x1c, 0x4b, 0xa7, 0x33, 0xed, 0x91, 0x89,
	0xf2, 0x0f, 0xe3, 0x61, 0x17, 0x54, 0xc6, 0x1f, 0x0a, 0xa7, 0xa8, 0x30, 0x6e, 0x9e, 0x8e, 0xb1,
	0x9d, 0x45, 0x6f, 0xa9, 0x60, 0x7f, 0x4e, 0x1c, 0x37, 0x40, 0x68, 0xa4, 0x77, 0x4a, 0x51, 0x49,
	0x2b, 0xee, 0x77, 0x16, 0xb8, 0x71, 0x96, 0x6e, 0x8a, 0x84, 0x4b, 0xed, 0x14, 0x78, 0x5f, 0x18,
	0xf1, 0xcc, 0x2e, 0xe7, 0xe8, 0x31, 0x73, 0x4f, 0x3c, 0x66, 0xae, 0x7c, 0xc5, 0x06, 0x2d, 0x77,
	0x8d, 0xbe, 0x24, 0x5c, 0x30, 0x8c, 0xf6, 0xb6, 0x13, 0xde, 0xd1, 0x30, 0x52, 0x64, 0xfd, 0xeb,
	0xa7, 0x4a, 0xec, 0x28, 0xa8, 0x5d, 0x97, 0xa4, 0xba, 0x01, 0x96, 0xcf, 0x25, 0x55, 0xed, 0xa7,
	0x49, 0xb0, 0xbc, 0x69, 0x1c, 0xe3, 0x74, 0xe8, 0x32, 0x1c, 0x49, 0xde, 0x63, 0x88, 0x80, 0x4d,
	0xd9, 0x31, 0x56, 0xad, 0x9d, 0x4d, 0x83, 0x73, 0xe1, 0x5c, 0xb5, 0xdc, 0xc2, 0x62, 0x63, 0xc2,
	0x9f, 0xa2, 0x4c, 0x13, 0x2d, 0x04, 0x25, 0x44, 0x22, 0x53, 0x23, 0xf7, 0x9f, 0xd6, 0xb0, 0x11,
	0x89, 0x74, 0x91, 0xcf, 0x41, 0x89, 0x50, 0x61, 0x8a, 0xe4, 0x55, 0x91, 0xbb, 0x97, 0x28, 0x22,
	0xb1, 0x09, 0x15, 0x1a, 0x7b, 0x45, 0x6e, 0x60, 0x68, 0xb0, 0x25, 0xcf, 0x6d, 0xf5, 0x36, 0xbf,
	0xc8, 0xd9, 0x96, 0xee, 0x61, 0xa8, 0xe3, 0x9e, 0x82, 0xb7, 0xc7, 0xdf, 0x2c, 0x2d, 0xa5, 0x86,
	0xdf, 0x57, 0x4c, 0x43, 0x52, 0xd1, 0x65, 0x17, 0x4a, 0xf9, 0x95, 0x60, 0x6e, 0x4c, 0xf8, 0xf3,
	0x63, 0x6f, 0x95, 0xb2, 0x2f, 0xfe, 0x60, 0x01, 0x3b, 0xdb, 0x2f, 0xfc, 0x02, 0x4c, 0xb2, 0x7e,
	0x82, 0xb9, 0x63, 0xa9, 0xfb, 0x7a, 0x99, 0xfd, 0xa9, 0x97, 0xe0, 0x95, 0x95, 0xb3, 0x73, 0xbe,
	0xc6, 0x6c, 0x7f, 0x2c, 0x89, 0xe5, 0x81, 0x8f, 0x4e, 0x27, 0xd6, 0xc5, 0x4f, 0xa5, 0xfd, 0x81,
	0x04, 0xba, 0x03, 0xde, 0xff, 0x97, 0x40, 0x5e, 0x19, 0x14, 0x64, 0x47, 0x30, 0xff, 0x87, 0x67,
	0xd5, 0x7e, 0xcf, 0x81, 0xd9, 0xe3, 0x29, 0xff, 0xec, 0x23, 0x29, 0xff, 0xbf, 0xfa, 0x48, 0x82,
	0x11, 0x98, 0x33, 0x6a, 0x1e, 0xc5, 0x1c, 0xed, 0x24, 0x38, 0x32, 0xba, 0x79, 0x99, 0xb3, 0xf4,
	0x67, 0x35, 0xe6, 0x9a, 0x81, 0x6c, 0xbf, 0x27, 0x4f, 0x60, 0x05, 0xbc, 0x7b, 0x91, 0x13, 0x18,
	0xfb, 0x24, 0xf3, 0x1e, 0xfe, 0xf8, 0xd5, 0xaf, 0xbf, 0x15, 0x73, 0x95, 0x1c, 0x58, 0x89, 0xa9,
	0x6e, 0x4a, 0xa9, 0xfc, 0x99, 0xfd, 0x79, 0x0b, 0x19, 0x60, 0xe7, 0xf0, 0xcf, 0x83, 0x12, 0xf8,
	0xae, 0xb5, 0x53, 0x54, 0x23, 0x6a, 0xfd, 0x15, 0x00, 0x00, 0xff, 0xff, 0xb4, 0xd7, 0xb8, 0xc8,
	0x7a, 0x0c, 0x00, 0x00,
}