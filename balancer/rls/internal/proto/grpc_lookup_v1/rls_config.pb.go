// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/rls/grpc_lookup_v1/rls_config.proto

package grpc_lookup_v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// Specify how to process a request when not already in the cache.
type RouteLookupConfig_RequestProcessingStrategy int32

const (
	RouteLookupConfig_STRATEGY_UNSPECIFIED RouteLookupConfig_RequestProcessingStrategy = 0
	// Query the RLS and process the request using target returned by the
	// lookup. The target will then be cached and used for processing
	// subsequent requests for the same key. Any errors during lookup service
	// processing will fall back to default target for request processing.
	RouteLookupConfig_SYNC_LOOKUP_DEFAULT_TARGET_ON_ERROR RouteLookupConfig_RequestProcessingStrategy = 1
	// Query the RLS and process the request using target returned by the
	// lookup. The target will then be cached and used for processing
	// subsequent requests for the same key. Any errors during lookup service
	// processing will return an error back to the client.  Services with
	// strict regional routing requirements should use this strategy.
	RouteLookupConfig_SYNC_LOOKUP_CLIENT_SEES_ERROR RouteLookupConfig_RequestProcessingStrategy = 2
	// Query the RLS asynchronously but respond with the default target.  The
	// target in the lookup response will then be cached and used for
	// subsequent requests.  Services with strict latency requirements (but not
	// strict regional routing requirements) should use this strategy.
	RouteLookupConfig_ASYNC_LOOKUP_DEFAULT_TARGET_ON_MISS RouteLookupConfig_RequestProcessingStrategy = 3
)

var RouteLookupConfig_RequestProcessingStrategy_name = map[int32]string{
	0: "STRATEGY_UNSPECIFIED",
	1: "SYNC_LOOKUP_DEFAULT_TARGET_ON_ERROR",
	2: "SYNC_LOOKUP_CLIENT_SEES_ERROR",
	3: "ASYNC_LOOKUP_DEFAULT_TARGET_ON_MISS",
}

var RouteLookupConfig_RequestProcessingStrategy_value = map[string]int32{
	"STRATEGY_UNSPECIFIED":                0,
	"SYNC_LOOKUP_DEFAULT_TARGET_ON_ERROR": 1,
	"SYNC_LOOKUP_CLIENT_SEES_ERROR":       2,
	"ASYNC_LOOKUP_DEFAULT_TARGET_ON_MISS": 3,
}

func (x RouteLookupConfig_RequestProcessingStrategy) String() string {
	return proto.EnumName(RouteLookupConfig_RequestProcessingStrategy_name, int32(x))
}

func (RouteLookupConfig_RequestProcessingStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f013e3228551a7a8, []int{3, 0}
}

// Extract a key based on a given name (e.g. header name or query parameter
// name).  The name must match one of the names listed in the "name" field.  If
// the "required_match" field is true, one of the specified names must be
// present for the keybuilder to match.
type NameMatcher struct {
	// The name that will be used in the RLS key_map to refer to this value.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Ordered list of names (headers or query parameter names) that can supply
	// this value; the first one with a non-empty value is used.
	Names []string `protobuf:"bytes,2,rep,name=names,proto3" json:"names,omitempty"`
	// If true, make this extraction required; the key builder will not match
	// if no value is found.
	RequiredMatch        bool     `protobuf:"varint,3,opt,name=required_match,json=requiredMatch,proto3" json:"required_match,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameMatcher) Reset()         { *m = NameMatcher{} }
func (m *NameMatcher) String() string { return proto.CompactTextString(m) }
func (*NameMatcher) ProtoMessage()    {}
func (*NameMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_f013e3228551a7a8, []int{0}
}

func (m *NameMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameMatcher.Unmarshal(m, b)
}
func (m *NameMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameMatcher.Marshal(b, m, deterministic)
}
func (m *NameMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameMatcher.Merge(m, src)
}
func (m *NameMatcher) XXX_Size() int {
	return xxx_messageInfo_NameMatcher.Size(m)
}
func (m *NameMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_NameMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_NameMatcher proto.InternalMessageInfo

func (m *NameMatcher) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *NameMatcher) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *NameMatcher) GetRequiredMatch() bool {
	if m != nil {
		return m.RequiredMatch
	}
	return false
}

// A GrpcKeyBuilder applies to a given gRPC service, name, and headers.
type GrpcKeyBuilder struct {
	Names []*GrpcKeyBuilder_Name `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	// Extract keys from all listed headers.
	// For gRPC, it is an error to specify "required_match" on the NameMatcher
	// protos.
	Headers              []*NameMatcher `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GrpcKeyBuilder) Reset()         { *m = GrpcKeyBuilder{} }
func (m *GrpcKeyBuilder) String() string { return proto.CompactTextString(m) }
func (*GrpcKeyBuilder) ProtoMessage()    {}
func (*GrpcKeyBuilder) Descriptor() ([]byte, []int) {
	return fileDescriptor_f013e3228551a7a8, []int{1}
}

func (m *GrpcKeyBuilder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcKeyBuilder.Unmarshal(m, b)
}
func (m *GrpcKeyBuilder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcKeyBuilder.Marshal(b, m, deterministic)
}
func (m *GrpcKeyBuilder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcKeyBuilder.Merge(m, src)
}
func (m *GrpcKeyBuilder) XXX_Size() int {
	return xxx_messageInfo_GrpcKeyBuilder.Size(m)
}
func (m *GrpcKeyBuilder) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcKeyBuilder.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcKeyBuilder proto.InternalMessageInfo

func (m *GrpcKeyBuilder) GetNames() []*GrpcKeyBuilder_Name {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *GrpcKeyBuilder) GetHeaders() []*NameMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

// To match, one of the given Name fields must match; the service and method
// fields are specified as fixed strings.  The service name is required and
// includes the proto package name.  The method name may be omitted, in
// which case any method on the given service is matched.
type GrpcKeyBuilder_Name struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Method               string   `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcKeyBuilder_Name) Reset()         { *m = GrpcKeyBuilder_Name{} }
func (m *GrpcKeyBuilder_Name) String() string { return proto.CompactTextString(m) }
func (*GrpcKeyBuilder_Name) ProtoMessage()    {}
func (*GrpcKeyBuilder_Name) Descriptor() ([]byte, []int) {
	return fileDescriptor_f013e3228551a7a8, []int{1, 0}
}

func (m *GrpcKeyBuilder_Name) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcKeyBuilder_Name.Unmarshal(m, b)
}
func (m *GrpcKeyBuilder_Name) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcKeyBuilder_Name.Marshal(b, m, deterministic)
}
func (m *GrpcKeyBuilder_Name) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcKeyBuilder_Name.Merge(m, src)
}
func (m *GrpcKeyBuilder_Name) XXX_Size() int {
	return xxx_messageInfo_GrpcKeyBuilder_Name.Size(m)
}
func (m *GrpcKeyBuilder_Name) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcKeyBuilder_Name.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcKeyBuilder_Name proto.InternalMessageInfo

func (m *GrpcKeyBuilder_Name) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *GrpcKeyBuilder_Name) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

// An HttpKeyBuilder applies to a given HTTP URL and headers.
//
// Path and host patterns use the matching syntax from gRPC transcoding to
// extract named key/value pairs from the path and host components of the URL:
// https://github.com/googleapis/googleapis/blob/master/google/api/http.proto
//
// It is invalid to specify the same key name in multiple places in a pattern.
//
// For a service where the project id can be expressed either as a subdomain or
// in the path, separate HttpKeyBuilders must be used:
//     host_pattern: 'example.com' path_pattern: '/{id}/{object}/**'
//     host_pattern: '{id}.example.com' path_pattern: '/{object}/**'
// If the host is exactly 'example.com', the first path segment will be used as
// the id and the second segment as the object. If the host has a subdomain, the
// subdomain will be used as the id and the first segment as the object. If
// neither pattern matches, no keys will be extracted.
type HttpKeyBuilder struct {
	// host_pattern is an ordered list of host template patterns for the desired
	// value.  If any host_pattern values are specified, then at least one must
	// match, and the last one wins and sets any specified variables.  A host
	// consists of labels separated by dots. Each label is matched against the
	// label in the pattern as follows:
	//   - "*": Matches any single label.
	//   - "**": Matches zero or more labels (first or last part of host only).
	//   - "{<name>=...}": One or more label capture, where "..." can be any
	//      template that does not include a capture.
	//   - "{<name>}": A single label capture. Identical to {<name>=*}.
	//
	// Examples:
	//   - "example.com": Only applies to the exact host example.com.
	//   - "*.example.com": Matches subdomains of example.com.
	//   - "**.example.com": matches example.com, and all levels of subdomains.
	//   - "{project}.example.com": Extracts the third level subdomain.
	//   - "{project=**}.example.com": Extracts the third level+ subdomains.
	//   - "{project=**}": Extracts the entire host.
	HostPatterns []string `protobuf:"bytes,1,rep,name=host_patterns,json=hostPatterns,proto3" json:"host_patterns,omitempty"`
	// path_pattern is an ordered list of path template patterns for the desired
	// value.  If any path_pattern values are specified, then at least one must
	// match, and the last one wins and sets any specified variables.  A path
	// consists of segments separated by slashes. Each segment is matched against
	// the segment in the pattern as follows:
	//   - "*": Matches any single segment.
	//   - "**": Matches zero or more segments (first or last part of path only).
	//   - "{<name>=...}": One or more segment capture, where "..." can be any
	//      template that does not include a capture.
	//   - "{<name>}": A single segment capture. Identical to {<name>=*}.
	// A custom method may also be specified by appending ":" and the custom
	// method name or "*" to indicate any custom method (including no custom
	// method).  For example, "/*/projects/{project_id}/**:*" extracts
	// `{project_id}` for any version, resource and custom method that includes
	// it.  By default, any custom method will be matched.
	//
	// Examples:
	//   - "/v1/{name=messages/*}": extracts a name like "messages/12345".
	//   - "/v1/messages/{message_id}": extracts a message_id like "12345".
	//   - "/v1/users/{user_id}/messages/{message_id}": extracts two key values.
	PathPatterns []string `protobuf:"bytes,2,rep,name=path_patterns,json=pathPatterns,proto3" json:"path_patterns,omitempty"`
	// List of query parameter names to try to match.
	// For example: ["parent", "name", "resource.name"]
	// We extract all the specified query_parameters (case-sensitively).  If any
	// are marked as "required_match" and are not present, this keybuilder fails
	// to match.  If a given parameter appears multiple times (?foo=a&foo=b) we
	// will report it as a comma-separated string (foo=a,b).
	QueryParameters []*NameMatcher `protobuf:"bytes,3,rep,name=query_parameters,json=queryParameters,proto3" json:"query_parameters,omitempty"`
	// List of headers to try to match.
	// We extract all the specified header values (case-insensitively).  If any
	// are marked as "required_match" and are not present, this keybuilder fails
	// to match.  If a given header appears multiple times in the request we will
	// report it as a comma-separated string, in standard HTTP fashion.
	Headers              []*NameMatcher `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *HttpKeyBuilder) Reset()         { *m = HttpKeyBuilder{} }
func (m *HttpKeyBuilder) String() string { return proto.CompactTextString(m) }
func (*HttpKeyBuilder) ProtoMessage()    {}
func (*HttpKeyBuilder) Descriptor() ([]byte, []int) {
	return fileDescriptor_f013e3228551a7a8, []int{2}
}

func (m *HttpKeyBuilder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpKeyBuilder.Unmarshal(m, b)
}
func (m *HttpKeyBuilder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpKeyBuilder.Marshal(b, m, deterministic)
}
func (m *HttpKeyBuilder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpKeyBuilder.Merge(m, src)
}
func (m *HttpKeyBuilder) XXX_Size() int {
	return xxx_messageInfo_HttpKeyBuilder.Size(m)
}
func (m *HttpKeyBuilder) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpKeyBuilder.DiscardUnknown(m)
}

var xxx_messageInfo_HttpKeyBuilder proto.InternalMessageInfo

func (m *HttpKeyBuilder) GetHostPatterns() []string {
	if m != nil {
		return m.HostPatterns
	}
	return nil
}

func (m *HttpKeyBuilder) GetPathPatterns() []string {
	if m != nil {
		return m.PathPatterns
	}
	return nil
}

func (m *HttpKeyBuilder) GetQueryParameters() []*NameMatcher {
	if m != nil {
		return m.QueryParameters
	}
	return nil
}

func (m *HttpKeyBuilder) GetHeaders() []*NameMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

type RouteLookupConfig struct {
	// Ordered specifications for constructing keys for HTTP requests.  Last
	// match wins.  If no HttpKeyBuilder matches, an empty key_map will be sent to
	// the lookup service; it should likely reply with a global default route
	// and raise an alert.
	HttpKeybuilders []*HttpKeyBuilder `protobuf:"bytes,1,rep,name=http_keybuilders,json=httpKeybuilders,proto3" json:"http_keybuilders,omitempty"`
	// Unordered specifications for constructing keys for gRPC requests.  All
	// GrpcKeyBuilders on this list must have unique "name" fields so that the
	// client is free to prebuild a hash map keyed by name.  If no GrpcKeyBuilder
	// matches, an empty key_map will be sent to the lookup service; it should
	// likely reply with a global default route and raise an alert.
	GrpcKeybuilders []*GrpcKeyBuilder `protobuf:"bytes,2,rep,name=grpc_keybuilders,json=grpcKeybuilders,proto3" json:"grpc_keybuilders,omitempty"`
	// The name of the lookup service as a gRPC URI.  Typically, this will be
	// a subdomain of the target, such as "lookup.datastore.googleapis.com".
	LookupService string `protobuf:"bytes,3,opt,name=lookup_service,json=lookupService,proto3" json:"lookup_service,omitempty"`
	// Configure a timeout value for lookup service requests.
	// Defaults to 10 seconds if not specified.
	LookupServiceTimeout *duration.Duration `protobuf:"bytes,4,opt,name=lookup_service_timeout,json=lookupServiceTimeout,proto3" json:"lookup_service_timeout,omitempty"`
	// How long are responses valid for (like HTTP Cache-Control).
	// If omitted or zero, the longest valid cache time is used.
	// This value is clamped to 5 minutes to avoid unflushable bad responses.
	MaxAge *duration.Duration `protobuf:"bytes,5,opt,name=max_age,json=maxAge,proto3" json:"max_age,omitempty"`
	// After a response has been in the client cache for this amount of time
	// and is re-requested, start an asynchronous RPC to re-validate it.
	// This value should be less than max_age by at least the length of a
	// typical RTT to the Route Lookup Service to fully mask the RTT latency.
	// If omitted, keys are only re-requested after they have expired.
	StaleAge *duration.Duration `protobuf:"bytes,6,opt,name=stale_age,json=staleAge,proto3" json:"stale_age,omitempty"`
	// Rough indicator of amount of memory to use for the client cache.  Some of
	// the data structure overhead is not accounted for, so actual memory consumed
	// will be somewhat greater than this value.  If this field is omitted or set
	// to zero, a client default will be used.  The value may be capped to a lower
	// amount based on client configuration.
	CacheSizeBytes int64 `protobuf:"varint,7,opt,name=cache_size_bytes,json=cacheSizeBytes,proto3" json:"cache_size_bytes,omitempty"`
	// This is a list of all the possible targets that can be returned by the
	// lookup service.  If a target not on this list is returned, it will be
	// treated the same as an RPC error from the RLS.
	ValidTargets []string `protobuf:"bytes,8,rep,name=valid_targets,json=validTargets,proto3" json:"valid_targets,omitempty"`
	// This value provides a default target to use if needed.  It will be used for
	// request processing strategy SYNC_LOOKUP_DEFAULT_TARGET_ON_ERROR if RLS
	// returns an error, or strategy ASYNC_LOOKUP_DEFAULT_TARGET_ON_MISS if RLS
	// returns an error or there is a cache miss in the client.  It will also be
	// used if there are no healthy backends for an RLS target.  Note that
	// requests can be routed only to a subdomain of the original target,
	// e.g. "us_east_1.cloudbigtable.googleapis.com".
	DefaultTarget             string                                      `protobuf:"bytes,9,opt,name=default_target,json=defaultTarget,proto3" json:"default_target,omitempty"`
	RequestProcessingStrategy RouteLookupConfig_RequestProcessingStrategy `protobuf:"varint,10,opt,name=request_processing_strategy,json=requestProcessingStrategy,proto3,enum=grpc.lookup.v1.RouteLookupConfig_RequestProcessingStrategy" json:"request_processing_strategy,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                                    `json:"-"`
	XXX_unrecognized          []byte                                      `json:"-"`
	XXX_sizecache             int32                                       `json:"-"`
}

func (m *RouteLookupConfig) Reset()         { *m = RouteLookupConfig{} }
func (m *RouteLookupConfig) String() string { return proto.CompactTextString(m) }
func (*RouteLookupConfig) ProtoMessage()    {}
func (*RouteLookupConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_f013e3228551a7a8, []int{3}
}

func (m *RouteLookupConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteLookupConfig.Unmarshal(m, b)
}
func (m *RouteLookupConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteLookupConfig.Marshal(b, m, deterministic)
}
func (m *RouteLookupConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteLookupConfig.Merge(m, src)
}
func (m *RouteLookupConfig) XXX_Size() int {
	return xxx_messageInfo_RouteLookupConfig.Size(m)
}
func (m *RouteLookupConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteLookupConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RouteLookupConfig proto.InternalMessageInfo

func (m *RouteLookupConfig) GetHttpKeybuilders() []*HttpKeyBuilder {
	if m != nil {
		return m.HttpKeybuilders
	}
	return nil
}

func (m *RouteLookupConfig) GetGrpcKeybuilders() []*GrpcKeyBuilder {
	if m != nil {
		return m.GrpcKeybuilders
	}
	return nil
}

func (m *RouteLookupConfig) GetLookupService() string {
	if m != nil {
		return m.LookupService
	}
	return ""
}

func (m *RouteLookupConfig) GetLookupServiceTimeout() *duration.Duration {
	if m != nil {
		return m.LookupServiceTimeout
	}
	return nil
}

func (m *RouteLookupConfig) GetMaxAge() *duration.Duration {
	if m != nil {
		return m.MaxAge
	}
	return nil
}

func (m *RouteLookupConfig) GetStaleAge() *duration.Duration {
	if m != nil {
		return m.StaleAge
	}
	return nil
}

func (m *RouteLookupConfig) GetCacheSizeBytes() int64 {
	if m != nil {
		return m.CacheSizeBytes
	}
	return 0
}

func (m *RouteLookupConfig) GetValidTargets() []string {
	if m != nil {
		return m.ValidTargets
	}
	return nil
}

func (m *RouteLookupConfig) GetDefaultTarget() string {
	if m != nil {
		return m.DefaultTarget
	}
	return ""
}

func (m *RouteLookupConfig) GetRequestProcessingStrategy() RouteLookupConfig_RequestProcessingStrategy {
	if m != nil {
		return m.RequestProcessingStrategy
	}
	return RouteLookupConfig_STRATEGY_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("grpc.s.lookup.v1.RouteLookupConfig_RequestProcessingStrategy", RouteLookupConfig_RequestProcessingStrategy_name, RouteLookupConfig_RequestProcessingStrategy_value)
	proto.RegisterType((*NameMatcher)(nil), "grpc.lookup.v1.NameMatcher")
	proto.RegisterType((*GrpcKeyBuilder)(nil), "grpc.lookup.v1.GrpcKeyBuilder")
	proto.RegisterType((*GrpcKeyBuilder_Name)(nil), "grpc.lookup.v1.GrpcKeyBuilder.Name")
	proto.RegisterType((*HttpKeyBuilder)(nil), "grpc.lookup.v1.HttpKeyBuilder")
	proto.RegisterType((*RouteLookupConfig)(nil), "grpc.lookup.v1.RouteLookupConfig")
}

func init() {
	proto.RegisterFile("grpc/rls/grpc_lookup_v1/rls_config.proto", fileDescriptor_f013e3228551a7a8)
}

var fileDescriptor_f013e3228551a7a8 = []byte{
	// 742 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xdb, 0x36,
	0x18, 0x9d, 0xa2, 0xd4, 0x89, 0x99, 0x46, 0x71, 0x85, 0xa0, 0x50, 0x5a, 0xac, 0xf0, 0x1c, 0x14,
	0xd3, 0xc5, 0x20, 0xa3, 0x1e, 0x36, 0x6c, 0xd8, 0x95, 0xed, 0x28, 0x99, 0x51, 0xd7, 0x36, 0x28,
	0xe5, 0xa2, 0xc3, 0x00, 0x82, 0x96, 0xbf, 0x48, 0x42, 0x24, 0x53, 0xa5, 0x28, 0xa3, 0xee, 0xde,
	0x68, 0xc0, 0xde, 0x60, 0x2f, 0xb2, 0xb7, 0x19, 0x28, 0x4a, 0x9e, 0xed, 0x2d, 0x4b, 0xef, 0xf4,
	0x1d, 0x9e, 0x73, 0xc4, 0xef, 0x8f, 0xc8, 0x0e, 0x79, 0x16, 0x74, 0x79, 0x92, 0x77, 0xe5, 0x07,
	0x49, 0x18, 0xbb, 0x2f, 0x32, 0xb2, 0x7a, 0x23, 0x21, 0x12, 0xb0, 0xe5, 0x5d, 0x1c, 0x3a, 0x19,
	0x67, 0x82, 0x99, 0x86, 0x24, 0x38, 0x8a, 0xe0, 0xac, 0xde, 0xbc, 0x78, 0x15, 0x32, 0x16, 0x26,
	0xd0, 0x2d, 0x4f, 0xe7, 0xc5, 0x5d, 0x77, 0x51, 0x70, 0x2a, 0x62, 0xb6, 0x54, 0xfc, 0xce, 0xaf,
	0xe8, 0x64, 0x42, 0x53, 0x78, 0x47, 0x45, 0x10, 0x01, 0x37, 0x5b, 0x48, 0xbf, 0x87, 0xb5, 0xa5,
	0xb5, 0x35, 0xbb, 0x89, 0xe5, 0xa7, 0x79, 0x8e, 0x9e, 0x2c, 0x69, 0x0a, 0xb9, 0x75, 0xd0, 0xd6,
	0xed, 0x26, 0x56, 0x81, 0xf9, 0x1a, 0x19, 0x1c, 0x3e, 0x14, 0x31, 0x87, 0x05, 0x49, 0xa5, 0xd6,
	0xd2, 0xdb, 0x9a, 0x7d, 0x8c, 0x4f, 0x6b, 0xb4, 0x34, 0xec, 0xfc, 0xa9, 0x21, 0xe3, 0x86, 0x67,
	0xc1, 0x5b, 0x58, 0x0f, 0x8a, 0x38, 0x59, 0x00, 0x37, 0x7f, 0xac, 0xfd, 0xb4, 0xb6, 0x6e, 0x9f,
	0xf4, 0x2e, 0x9d, 0xdd, 0x0b, 0x3b, 0xbb, 0x74, 0x47, 0x5e, 0xae, 0xfe, 0xe9, 0x77, 0xe8, 0x28,
	0x02, 0xba, 0x00, 0xae, 0x2e, 0x73, 0xd2, 0x7b, 0xb9, 0x2f, 0xde, 0x4a, 0x05, 0xd7, 0xdc, 0x17,
	0x3f, 0xa0, 0x43, 0x89, 0x9b, 0x16, 0x3a, 0xca, 0x81, 0xaf, 0xe2, 0x00, 0xaa, 0xfc, 0xea, 0xd0,
	0x7c, 0x8e, 0x1a, 0x29, 0x88, 0x88, 0x2d, 0xac, 0x83, 0xf2, 0xa0, 0x8a, 0x3a, 0x7f, 0x69, 0xc8,
	0xf8, 0x59, 0x88, 0x6c, 0xeb, 0xfa, 0x97, 0xe8, 0x34, 0x62, 0xb9, 0x20, 0x19, 0x15, 0x02, 0xf8,
	0x52, 0xa5, 0xd1, 0xc4, 0x4f, 0x25, 0x38, 0xab, 0x30, 0x49, 0xca, 0xa8, 0x88, 0xfe, 0x21, 0xa9,
	0xda, 0x3d, 0x95, 0xe0, 0x86, 0x74, 0x8d, 0x5a, 0x1f, 0x0a, 0xe0, 0x6b, 0x92, 0x51, 0x4e, 0x53,
	0x10, 0x32, 0x2d, 0xfd, 0xf1, 0xb4, 0xce, 0x4a, 0xd1, 0x6c, 0xa3, 0xd9, 0xae, 0xca, 0xe1, 0xe7,
	0x57, 0xa5, 0xf3, 0x47, 0x03, 0x3d, 0xc3, 0xac, 0x10, 0x30, 0x2e, 0x79, 0xc3, 0x72, 0x88, 0xcc,
	0x11, 0x6a, 0x45, 0x42, 0x64, 0xe4, 0x1e, 0xd6, 0x73, 0x95, 0x71, 0xdd, 0xa8, 0x57, 0xfb, 0xae,
	0xbb, 0x85, 0xc1, 0x67, 0x91, 0x8a, 0x6b, 0x99, 0xb4, 0x2a, 0x87, 0x75, 0xdb, 0xea, 0xe0, 0xbf,
	0xad, 0x76, 0x7b, 0x8e, 0xcf, 0x42, 0x15, 0x6f, 0xac, 0x5e, 0x23, 0xa3, 0x1a, 0xf9, 0xba, 0x81,
	0x7a, 0xd9, 0xa7, 0x53, 0x85, 0x7a, 0x55, 0x1b, 0xa7, 0xe8, 0xf9, 0x2e, 0x8d, 0x88, 0x38, 0x05,
	0x56, 0x08, 0xeb, 0xb0, 0xad, 0xd9, 0x27, 0xbd, 0x0b, 0x47, 0x2d, 0x83, 0x53, 0x2f, 0x83, 0x73,
	0x55, 0x2d, 0x03, 0x3e, 0xdf, 0x71, 0xf2, 0x95, 0xcc, 0xec, 0xa1, 0xa3, 0x94, 0x7e, 0x24, 0x34,
	0x04, 0xeb, 0xc9, 0x63, 0x0e, 0x8d, 0x94, 0x7e, 0xec, 0x87, 0x60, 0x7e, 0x8f, 0x9a, 0xb9, 0xa0,
	0x09, 0x94, 0xaa, 0xc6, 0x63, 0xaa, 0xe3, 0x92, 0x2b, 0x75, 0x36, 0x6a, 0x05, 0x34, 0x88, 0x80,
	0xe4, 0xf1, 0x27, 0x20, 0xf3, 0xb5, 0x80, 0xdc, 0x3a, 0x6a, 0x6b, 0xb6, 0x8e, 0x8d, 0x12, 0xf7,
	0xe2, 0x4f, 0x30, 0x90, 0xa8, 0x9c, 0xae, 0x15, 0x4d, 0xe2, 0x05, 0x11, 0x94, 0x87, 0x20, 0x72,
	0xeb, 0x58, 0x4d, 0x57, 0x09, 0xfa, 0x0a, 0x93, 0x25, 0x5b, 0xc0, 0x1d, 0x2d, 0x12, 0x51, 0xd1,
	0xac, 0xa6, 0x2a, 0x59, 0x85, 0x2a, 0x9e, 0xf9, 0x1b, 0x7a, 0x29, 0x37, 0x16, 0xe4, 0x44, 0x73,
	0x16, 0x40, 0x9e, 0xc7, 0xcb, 0x90, 0xe4, 0x82, 0x53, 0x01, 0xe1, 0xda, 0x42, 0x6d, 0xcd, 0x36,
	0x7a, 0x3f, 0xed, 0xf7, 0xeb, 0x5f, 0x73, 0xe3, 0x60, 0x65, 0x32, 0xdb, 0x78, 0x78, 0x95, 0x05,
	0xbe, 0xe0, 0x0f, 0x1d, 0x75, 0x7e, 0xd7, 0xd0, 0xc5, 0x83, 0x42, 0xd3, 0x42, 0xe7, 0x9e, 0x8f,
	0xfb, 0xbe, 0x7b, 0xf3, 0x9e, 0xdc, 0x4e, 0xbc, 0x99, 0x3b, 0x1c, 0x5d, 0x8f, 0xdc, 0xab, 0xd6,
	0x17, 0xe6, 0xd7, 0xe8, 0xd2, 0x7b, 0x3f, 0x19, 0x92, 0xf1, 0x74, 0xfa, 0xf6, 0x76, 0x46, 0xae,
	0xdc, 0xeb, 0xfe, 0xed, 0xd8, 0x27, 0x7e, 0x1f, 0xdf, 0xb8, 0x3e, 0x99, 0x4e, 0x88, 0x8b, 0xf1,
	0x14, 0xb7, 0x34, 0xf3, 0x2b, 0xf4, 0xe5, 0x36, 0x71, 0x38, 0x1e, 0xb9, 0x13, 0x9f, 0x78, 0xae,
	0xeb, 0x55, 0x94, 0x03, 0xe9, 0xd5, 0xff, 0x7f, 0xb3, 0x77, 0x23, 0xcf, 0x6b, 0xe9, 0x03, 0x0f,
	0x3d, 0x8b, 0xd9, 0x5e, 0x21, 0x06, 0x06, 0x4e, 0x72, 0x55, 0x81, 0x99, 0x6c, 0xed, 0x4c, 0xfb,
	0xe5, 0x9b, 0xaa, 0xd5, 0x21, 0x4b, 0xe8, 0x32, 0x74, 0x18, 0x0f, 0xcb, 0x27, 0xbb, 0xab, 0x34,
	0x7b, 0xcf, 0xf7, 0xbc, 0x51, 0x4e, 0xc4, 0xb7, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x47, 0xa7,
	0x94, 0xbe, 0xe0, 0x05, 0x00, 0x00,
}
