// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: envoy/extensions/filters/http/cache/v3/cache.proto

package cachev3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v31 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
	any1 "github.com/golang/protobuf/ptypes/any"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// [#extension: envoy.filters.http.cache]
// [#next-free-field: 6]
type CacheConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Config specific to the cache storage implementation. Required unless ``disabled``
	// is true.
	// [#extension-category: envoy.http.cache]
	TypedConfig *any1.Any `protobuf:"bytes,1,opt,name=typed_config,json=typedConfig,proto3" json:"typed_config,omitempty"`
	// When true, the cache filter is a no-op filter.
	//
	// Possible use-cases for this include:
	// - Turning a filter on and off with :ref:`ECDS <envoy_v3_api_file_envoy/service/extension/v3/config_discovery.proto>`.
	// [#comment: once route-specific overrides are implemented, they are the more likely use-case.]
	Disabled *wrappers.BoolValue `protobuf:"bytes,5,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// List of matching rules that defines allowed ``Vary`` headers.
	//
	// The ``vary`` response header holds a list of header names that affect the
	// contents of a response, as described by
	// https://httpwg.org/specs/rfc7234.html#caching.negotiated.responses.
	//
	// During insertion, ``allowed_vary_headers`` acts as a allowlist: if a
	// response's ``vary`` header mentions any header names that aren't matched by any rules in
	// ``allowed_vary_headers``, that response will not be cached.
	//
	// During lookup, ``allowed_vary_headers`` controls what request headers will be
	// sent to the cache storage implementation.
	AllowedVaryHeaders []*v3.StringMatcher `protobuf:"bytes,2,rep,name=allowed_vary_headers,json=allowedVaryHeaders,proto3" json:"allowed_vary_headers,omitempty"`
	// [#not-implemented-hide:]
	// <TODO(toddmgreer) implement key customization>
	//
	// Modifies cache key creation by restricting which parts of the URL are included.
	KeyCreatorParams *CacheConfig_KeyCreatorParams `protobuf:"bytes,3,opt,name=key_creator_params,json=keyCreatorParams,proto3" json:"key_creator_params,omitempty"`
	// [#not-implemented-hide:]
	// <TODO(toddmgreer) implement size limit>
	//
	// Max body size the cache filter will insert into a cache. 0 means unlimited (though the cache
	// storage implementation may have its own limit beyond which it will reject insertions).
	MaxBodyBytes uint32 `protobuf:"varint,4,opt,name=max_body_bytes,json=maxBodyBytes,proto3" json:"max_body_bytes,omitempty"`
}

func (x *CacheConfig) Reset() {
	*x = CacheConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_cache_v3_cache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheConfig) ProtoMessage() {}

func (x *CacheConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_cache_v3_cache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheConfig.ProtoReflect.Descriptor instead.
func (*CacheConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_cache_v3_cache_proto_rawDescGZIP(), []int{0}
}

func (x *CacheConfig) GetTypedConfig() *any1.Any {
	if x != nil {
		return x.TypedConfig
	}
	return nil
}

func (x *CacheConfig) GetDisabled() *wrappers.BoolValue {
	if x != nil {
		return x.Disabled
	}
	return nil
}

func (x *CacheConfig) GetAllowedVaryHeaders() []*v3.StringMatcher {
	if x != nil {
		return x.AllowedVaryHeaders
	}
	return nil
}

func (x *CacheConfig) GetKeyCreatorParams() *CacheConfig_KeyCreatorParams {
	if x != nil {
		return x.KeyCreatorParams
	}
	return nil
}

func (x *CacheConfig) GetMaxBodyBytes() uint32 {
	if x != nil {
		return x.MaxBodyBytes
	}
	return 0
}

// [#not-implemented-hide:]
// Modifies cache key creation by restricting which parts of the URL are included.
type CacheConfig_KeyCreatorParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If true, exclude the URL scheme from the cache key. Set to true if your origins always
	// produce the same response for http and https requests.
	ExcludeScheme bool `protobuf:"varint,1,opt,name=exclude_scheme,json=excludeScheme,proto3" json:"exclude_scheme,omitempty"`
	// If true, exclude the host from the cache key. Set to true if your origins' responses don't
	// ever depend on host.
	ExcludeHost bool `protobuf:"varint,2,opt,name=exclude_host,json=excludeHost,proto3" json:"exclude_host,omitempty"`
	// If ``query_parameters_included`` is nonempty, only query parameters matched
	// by one or more of its matchers are included in the cache key. Any other
	// query params will not affect cache lookup.
	QueryParametersIncluded []*v31.QueryParameterMatcher `protobuf:"bytes,3,rep,name=query_parameters_included,json=queryParametersIncluded,proto3" json:"query_parameters_included,omitempty"`
	// If ``query_parameters_excluded`` is nonempty, query parameters matched by one
	// or more of its matchers are excluded from the cache key (even if also
	// matched by ``query_parameters_included``), and will not affect cache lookup.
	QueryParametersExcluded []*v31.QueryParameterMatcher `protobuf:"bytes,4,rep,name=query_parameters_excluded,json=queryParametersExcluded,proto3" json:"query_parameters_excluded,omitempty"`
}

func (x *CacheConfig_KeyCreatorParams) Reset() {
	*x = CacheConfig_KeyCreatorParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_cache_v3_cache_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CacheConfig_KeyCreatorParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheConfig_KeyCreatorParams) ProtoMessage() {}

func (x *CacheConfig_KeyCreatorParams) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_cache_v3_cache_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CacheConfig_KeyCreatorParams.ProtoReflect.Descriptor instead.
func (*CacheConfig_KeyCreatorParams) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_cache_v3_cache_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CacheConfig_KeyCreatorParams) GetExcludeScheme() bool {
	if x != nil {
		return x.ExcludeScheme
	}
	return false
}

func (x *CacheConfig_KeyCreatorParams) GetExcludeHost() bool {
	if x != nil {
		return x.ExcludeHost
	}
	return false
}

func (x *CacheConfig_KeyCreatorParams) GetQueryParametersIncluded() []*v31.QueryParameterMatcher {
	if x != nil {
		return x.QueryParametersIncluded
	}
	return nil
}

func (x *CacheConfig_KeyCreatorParams) GetQueryParametersExcluded() []*v31.QueryParameterMatcher {
	if x != nil {
		return x.QueryParametersExcluded
	}
	return nil
}

var File_envoy_extensions_filters_http_cache_v3_cache_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_cache_v3_cache_proto_rawDesc = []byte{
	0x0a, 0x32, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x63, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x2c, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x2f, 0x76, 0x33, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76,
	0x33, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x06, 0x0a, 0x0b,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x37, 0x0a, 0x0c, 0x74,
	0x79, 0x70, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x64, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x36, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x56, 0x0a, 0x14,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x72, 0x79, 0x5f, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x76, 0x33, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x52, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x56, 0x61, 0x72, 0x79, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x72, 0x0a, 0x12, 0x6b, 0x65, 0x79, 0x5f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x44, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4b, 0x65, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x10, 0x6b, 0x65, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x6f, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f,
	0x62, 0x6f, 0x64, 0x79, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x6d, 0x61, 0x78, 0x42, 0x6f, 0x64, 0x79, 0x42, 0x79, 0x74, 0x65, 0x73, 0x1a, 0xfc,
	0x02, 0x0a, 0x10, 0x4b, 0x65, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x78, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x68, 0x0a,
	0x19, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x17,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x49,
	0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x12, 0x68, 0x0a, 0x19, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x65, 0x78, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e,
	0x76, 0x33, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x17, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x64, 0x3a, 0x4a, 0x9a, 0xc5, 0x88, 0x1e, 0x45, 0x0a, 0x43, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4b, 0x65, 0x79,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x3a, 0x39, 0x9a,
	0xc5, 0x88, 0x1e, 0x34, 0x0a, 0x32, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63,
	0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0xa3, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06,
	0x02, 0x10, 0x02, 0x0a, 0x34, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x76, 0x33, 0x42, 0x0a, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67,
	0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x2f, 0x76, 0x33, 0x3b, 0x63, 0x61, 0x63, 0x68, 0x65, 0x76, 0x33, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_cache_v3_cache_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_cache_v3_cache_proto_rawDescData = file_envoy_extensions_filters_http_cache_v3_cache_proto_rawDesc
)

func file_envoy_extensions_filters_http_cache_v3_cache_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_cache_v3_cache_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_cache_v3_cache_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_cache_v3_cache_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_cache_v3_cache_proto_rawDescData
}

var file_envoy_extensions_filters_http_cache_v3_cache_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_filters_http_cache_v3_cache_proto_goTypes = []interface{}{
	(*CacheConfig)(nil),                  // 0: envoy.extensions.filters.http.cache.v3.CacheConfig
	(*CacheConfig_KeyCreatorParams)(nil), // 1: envoy.extensions.filters.http.cache.v3.CacheConfig.KeyCreatorParams
	(*any1.Any)(nil),                     // 2: google.protobuf.Any
	(*wrappers.BoolValue)(nil),           // 3: google.protobuf.BoolValue
	(*v3.StringMatcher)(nil),             // 4: envoy.type.matcher.v3.StringMatcher
	(*v31.QueryParameterMatcher)(nil),    // 5: envoy.config.route.v3.QueryParameterMatcher
}
var file_envoy_extensions_filters_http_cache_v3_cache_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.filters.http.cache.v3.CacheConfig.typed_config:type_name -> google.protobuf.Any
	3, // 1: envoy.extensions.filters.http.cache.v3.CacheConfig.disabled:type_name -> google.protobuf.BoolValue
	4, // 2: envoy.extensions.filters.http.cache.v3.CacheConfig.allowed_vary_headers:type_name -> envoy.type.matcher.v3.StringMatcher
	1, // 3: envoy.extensions.filters.http.cache.v3.CacheConfig.key_creator_params:type_name -> envoy.extensions.filters.http.cache.v3.CacheConfig.KeyCreatorParams
	5, // 4: envoy.extensions.filters.http.cache.v3.CacheConfig.KeyCreatorParams.query_parameters_included:type_name -> envoy.config.route.v3.QueryParameterMatcher
	5, // 5: envoy.extensions.filters.http.cache.v3.CacheConfig.KeyCreatorParams.query_parameters_excluded:type_name -> envoy.config.route.v3.QueryParameterMatcher
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_cache_v3_cache_proto_init() }
func file_envoy_extensions_filters_http_cache_v3_cache_proto_init() {
	if File_envoy_extensions_filters_http_cache_v3_cache_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_cache_v3_cache_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_extensions_filters_http_cache_v3_cache_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CacheConfig_KeyCreatorParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_http_cache_v3_cache_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_cache_v3_cache_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_cache_v3_cache_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_cache_v3_cache_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_cache_v3_cache_proto = out.File
	file_envoy_extensions_filters_http_cache_v3_cache_proto_rawDesc = nil
	file_envoy_extensions_filters_http_cache_v3_cache_proto_goTypes = nil
	file_envoy_extensions_filters_http_cache_v3_cache_proto_depIdxs = nil
}
