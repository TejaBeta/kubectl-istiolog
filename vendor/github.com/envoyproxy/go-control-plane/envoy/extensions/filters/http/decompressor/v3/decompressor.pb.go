// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: envoy/extensions/filters/http/decompressor/v3/decompressor.proto

package decompressorv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Decompressor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A decompressor library to use for both request and response decompression. Currently only
	// :ref:`envoy.compression.gzip.compressor<envoy_v3_api_msg_extensions.compression.gzip.decompressor.v3.Gzip>`
	// is included in Envoy.
	// [#extension-category: envoy.compression.decompressor]
	DecompressorLibrary *v3.TypedExtensionConfig `protobuf:"bytes,1,opt,name=decompressor_library,json=decompressorLibrary,proto3" json:"decompressor_library,omitempty"`
	// Configuration for request decompression. Decompression is enabled by default if left empty.
	RequestDirectionConfig *Decompressor_RequestDirectionConfig `protobuf:"bytes,2,opt,name=request_direction_config,json=requestDirectionConfig,proto3" json:"request_direction_config,omitempty"`
	// Configuration for response decompression. Decompression is enabled by default if left empty.
	ResponseDirectionConfig *Decompressor_ResponseDirectionConfig `protobuf:"bytes,3,opt,name=response_direction_config,json=responseDirectionConfig,proto3" json:"response_direction_config,omitempty"`
}

func (x *Decompressor) Reset() {
	*x = Decompressor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Decompressor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Decompressor) ProtoMessage() {}

func (x *Decompressor) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Decompressor.ProtoReflect.Descriptor instead.
func (*Decompressor) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_rawDescGZIP(), []int{0}
}

func (x *Decompressor) GetDecompressorLibrary() *v3.TypedExtensionConfig {
	if x != nil {
		return x.DecompressorLibrary
	}
	return nil
}

func (x *Decompressor) GetRequestDirectionConfig() *Decompressor_RequestDirectionConfig {
	if x != nil {
		return x.RequestDirectionConfig
	}
	return nil
}

func (x *Decompressor) GetResponseDirectionConfig() *Decompressor_ResponseDirectionConfig {
	if x != nil {
		return x.ResponseDirectionConfig
	}
	return nil
}

// Common configuration for filter behavior on both the request and response direction.
type Decompressor_CommonDirectionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Runtime flag that controls whether the filter is enabled for decompression or not. If set to false, the
	// filter will operate as a pass-through filter. If the message is unspecified, the filter will be enabled.
	Enabled *v3.RuntimeFeatureFlag `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// If set to true, will decompress response even if a ``no-transform`` cache control header is set.
	IgnoreNoTransformHeader bool `protobuf:"varint,2,opt,name=ignore_no_transform_header,json=ignoreNoTransformHeader,proto3" json:"ignore_no_transform_header,omitempty"`
}

func (x *Decompressor_CommonDirectionConfig) Reset() {
	*x = Decompressor_CommonDirectionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Decompressor_CommonDirectionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Decompressor_CommonDirectionConfig) ProtoMessage() {}

func (x *Decompressor_CommonDirectionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Decompressor_CommonDirectionConfig.ProtoReflect.Descriptor instead.
func (*Decompressor_CommonDirectionConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Decompressor_CommonDirectionConfig) GetEnabled() *v3.RuntimeFeatureFlag {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *Decompressor_CommonDirectionConfig) GetIgnoreNoTransformHeader() bool {
	if x != nil {
		return x.IgnoreNoTransformHeader
	}
	return false
}

// Configuration for filter behavior on the request direction.
type Decompressor_RequestDirectionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonConfig *Decompressor_CommonDirectionConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	// If set to true, and response decompression is enabled, the filter modifies the Accept-Encoding
	// request header by appending the decompressor_library's encoding. Defaults to true.
	AdvertiseAcceptEncoding *wrappers.BoolValue `protobuf:"bytes,2,opt,name=advertise_accept_encoding,json=advertiseAcceptEncoding,proto3" json:"advertise_accept_encoding,omitempty"`
}

func (x *Decompressor_RequestDirectionConfig) Reset() {
	*x = Decompressor_RequestDirectionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Decompressor_RequestDirectionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Decompressor_RequestDirectionConfig) ProtoMessage() {}

func (x *Decompressor_RequestDirectionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Decompressor_RequestDirectionConfig.ProtoReflect.Descriptor instead.
func (*Decompressor_RequestDirectionConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Decompressor_RequestDirectionConfig) GetCommonConfig() *Decompressor_CommonDirectionConfig {
	if x != nil {
		return x.CommonConfig
	}
	return nil
}

func (x *Decompressor_RequestDirectionConfig) GetAdvertiseAcceptEncoding() *wrappers.BoolValue {
	if x != nil {
		return x.AdvertiseAcceptEncoding
	}
	return nil
}

// Configuration for filter behavior on the response direction.
type Decompressor_ResponseDirectionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonConfig *Decompressor_CommonDirectionConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
}

func (x *Decompressor_ResponseDirectionConfig) Reset() {
	*x = Decompressor_ResponseDirectionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Decompressor_ResponseDirectionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Decompressor_ResponseDirectionConfig) ProtoMessage() {}

func (x *Decompressor_ResponseDirectionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Decompressor_ResponseDirectionConfig.ProtoReflect.Descriptor instead.
func (*Decompressor_ResponseDirectionConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Decompressor_ResponseDirectionConfig) GetCommonConfig() *Decompressor_CommonDirectionConfig {
	if x != nil {
		return x.CommonConfig
	}
	return nil
}

var File_envoy_extensions_filters_http_decompressor_v3_decompressor_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_rawDesc = []byte{
	0x0a, 0x40, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2f, 0x76, 0x33, 0x2f,
	0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x2d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x76,
	0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x24, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb2, 0x07, 0x0a, 0x0c, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f,
	0x72, 0x12, 0x67, 0x0a, 0x14, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f,
	0x72, 0x5f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x13, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x6f, 0x72, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x8c, 0x01, 0x0a, 0x18, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x52, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x64, 0x65,
	0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65,
	0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x16, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x8f, 0x01, 0x0a, 0x19, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x53, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x64, 0x65,
	0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65,
	0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x17, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x98, 0x01, 0x0a, 0x15,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x42, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67,
	0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x3b, 0x0a, 0x1a, 0x69, 0x67, 0x6e,
	0x6f, 0x72, 0x65, 0x5f, 0x6e, 0x6f, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x69,
	0x67, 0x6e, 0x6f, 0x72, 0x65, 0x4e, 0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a, 0xe8, 0x01, 0x0a, 0x16, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x76, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x51, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0c, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x56, 0x0a, 0x19, 0x61, 0x64, 0x76,
	0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x5f, 0x65, 0x6e,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x17, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74,
	0x69, 0x73, 0x65, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e,
	0x67, 0x1a, 0x91, 0x01, 0x0a, 0x17, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x76, 0x0a,
	0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x51, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f,
	0x72, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f,
	0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0xbf, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02,
	0x0a, 0x3b, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x64, 0x65,
	0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x33, 0x42, 0x11, 0x44,
	0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x63, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x2f, 0x76, 0x33, 0x3b, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_rawDescData = file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_rawDesc
)

func file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_rawDescData
}

var file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_goTypes = []interface{}{
	(*Decompressor)(nil),                         // 0: envoy.extensions.filters.http.decompressor.v3.Decompressor
	(*Decompressor_CommonDirectionConfig)(nil),   // 1: envoy.extensions.filters.http.decompressor.v3.Decompressor.CommonDirectionConfig
	(*Decompressor_RequestDirectionConfig)(nil),  // 2: envoy.extensions.filters.http.decompressor.v3.Decompressor.RequestDirectionConfig
	(*Decompressor_ResponseDirectionConfig)(nil), // 3: envoy.extensions.filters.http.decompressor.v3.Decompressor.ResponseDirectionConfig
	(*v3.TypedExtensionConfig)(nil),              // 4: envoy.config.core.v3.TypedExtensionConfig
	(*v3.RuntimeFeatureFlag)(nil),                // 5: envoy.config.core.v3.RuntimeFeatureFlag
	(*wrappers.BoolValue)(nil),                   // 6: google.protobuf.BoolValue
}
var file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_depIdxs = []int32{
	4, // 0: envoy.extensions.filters.http.decompressor.v3.Decompressor.decompressor_library:type_name -> envoy.config.core.v3.TypedExtensionConfig
	2, // 1: envoy.extensions.filters.http.decompressor.v3.Decompressor.request_direction_config:type_name -> envoy.extensions.filters.http.decompressor.v3.Decompressor.RequestDirectionConfig
	3, // 2: envoy.extensions.filters.http.decompressor.v3.Decompressor.response_direction_config:type_name -> envoy.extensions.filters.http.decompressor.v3.Decompressor.ResponseDirectionConfig
	5, // 3: envoy.extensions.filters.http.decompressor.v3.Decompressor.CommonDirectionConfig.enabled:type_name -> envoy.config.core.v3.RuntimeFeatureFlag
	1, // 4: envoy.extensions.filters.http.decompressor.v3.Decompressor.RequestDirectionConfig.common_config:type_name -> envoy.extensions.filters.http.decompressor.v3.Decompressor.CommonDirectionConfig
	6, // 5: envoy.extensions.filters.http.decompressor.v3.Decompressor.RequestDirectionConfig.advertise_accept_encoding:type_name -> google.protobuf.BoolValue
	1, // 6: envoy.extensions.filters.http.decompressor.v3.Decompressor.ResponseDirectionConfig.common_config:type_name -> envoy.extensions.filters.http.decompressor.v3.Decompressor.CommonDirectionConfig
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_init() }
func file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_init() {
	if File_envoy_extensions_filters_http_decompressor_v3_decompressor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Decompressor); i {
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
		file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Decompressor_CommonDirectionConfig); i {
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
		file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Decompressor_RequestDirectionConfig); i {
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
		file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Decompressor_ResponseDirectionConfig); i {
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
			RawDescriptor: file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_decompressor_v3_decompressor_proto = out.File
	file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_rawDesc = nil
	file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_goTypes = nil
	file_envoy_extensions_filters_http_decompressor_v3_decompressor_proto_depIdxs = nil
}
