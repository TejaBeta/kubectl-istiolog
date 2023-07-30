// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: contrib/envoy/extensions/filters/network/postgres_proxy/v3alpha/postgres_proxy.proto

package v3alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
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

// Upstream SSL operational modes.
type PostgresProxy_SSLMode int32

const (
	// Do not encrypt upstream connection to the server.
	PostgresProxy_DISABLE PostgresProxy_SSLMode = 0
	// Establish upstream SSL connection to the server. If the server does not
	// accept the request for SSL connection, the session is terminated.
	PostgresProxy_REQUIRE PostgresProxy_SSLMode = 1
)

// Enum value maps for PostgresProxy_SSLMode.
var (
	PostgresProxy_SSLMode_name = map[int32]string{
		0: "DISABLE",
		1: "REQUIRE",
	}
	PostgresProxy_SSLMode_value = map[string]int32{
		"DISABLE": 0,
		"REQUIRE": 1,
	}
)

func (x PostgresProxy_SSLMode) Enum() *PostgresProxy_SSLMode {
	p := new(PostgresProxy_SSLMode)
	*p = x
	return p
}

func (x PostgresProxy_SSLMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PostgresProxy_SSLMode) Descriptor() protoreflect.EnumDescriptor {
	return file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_enumTypes[0].Descriptor()
}

func (PostgresProxy_SSLMode) Type() protoreflect.EnumType {
	return &file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_enumTypes[0]
}

func (x PostgresProxy_SSLMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PostgresProxy_SSLMode.Descriptor instead.
func (PostgresProxy_SSLMode) EnumDescriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_rawDescGZIP(), []int{0, 0}
}

type PostgresProxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The human readable prefix to use when emitting :ref:`statistics
	// <config_network_filters_postgres_proxy_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// Controls whether SQL statements received in Frontend Query messages
	// are parsed. Parsing is required to produce Postgres proxy filter
	// metadata. Defaults to true.
	EnableSqlParsing *wrappers.BoolValue `protobuf:"bytes,2,opt,name=enable_sql_parsing,json=enableSqlParsing,proto3" json:"enable_sql_parsing,omitempty"`
	// Controls whether to terminate SSL session initiated by a client.
	// If the value is false, the Postgres proxy filter will not try to
	// terminate SSL session, but will pass all the packets to the upstream server.
	// If the value is true, the Postgres proxy filter will try to terminate SSL
	// session. In order to do that, the filter chain must use :ref:`starttls transport socket
	// <envoy_v3_api_msg_extensions.transport_sockets.starttls.v3.StartTlsConfig>`.
	// If the filter does not manage to terminate the SSL session, it will close the connection from the client.
	// Refer to official documentation for details
	// `SSL Session Encryption Message Flow <https://www.postgresql.org/docs/current/protocol-flow.html#id-1.10.5.7.11>`_.
	TerminateSsl bool `protobuf:"varint,3,opt,name=terminate_ssl,json=terminateSsl,proto3" json:"terminate_ssl,omitempty"`
	// Controls whether to establish upstream SSL connection to the server.
	// Envoy will try to establish upstream SSL connection to the server only when
	// Postgres filter is able to read Postgres payload in clear-text. It happens when
	// a client established a clear-text connection to Envoy or when a client established
	// SSL connection to Envoy and Postgres filter is configured to terminate SSL.
	// In order for upstream encryption to work, the corresponding cluster must be configured to use
	// :ref:`starttls transport socket <envoy_v3_api_msg_extensions.transport_sockets.starttls.v3.UpstreamStartTlsConfig>`.
	// Defaults to ``SSL_DISABLE``.
	UpstreamSsl PostgresProxy_SSLMode `protobuf:"varint,4,opt,name=upstream_ssl,json=upstreamSsl,proto3,enum=envoy.extensions.filters.network.postgres_proxy.v3alpha.PostgresProxy_SSLMode" json:"upstream_ssl,omitempty"`
}

func (x *PostgresProxy) Reset() {
	*x = PostgresProxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostgresProxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostgresProxy) ProtoMessage() {}

func (x *PostgresProxy) ProtoReflect() protoreflect.Message {
	mi := &file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostgresProxy.ProtoReflect.Descriptor instead.
func (*PostgresProxy) Descriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *PostgresProxy) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

func (x *PostgresProxy) GetEnableSqlParsing() *wrappers.BoolValue {
	if x != nil {
		return x.EnableSqlParsing
	}
	return nil
}

func (x *PostgresProxy) GetTerminateSsl() bool {
	if x != nil {
		return x.TerminateSsl
	}
	return false
}

func (x *PostgresProxy) GetUpstreamSsl() PostgresProxy_SSLMode {
	if x != nil {
		return x.UpstreamSsl
	}
	return PostgresProxy_DISABLE
}

var File_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto protoreflect.FileDescriptor

var file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_rawDesc = []byte{
	0x0a, 0x54, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x67,
	0x72, 0x65, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x37, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65,
	0x73, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x02, 0x0a, 0x0d, 0x50, 0x6f, 0x73, 0x74,
	0x67, 0x72, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x28, 0x0a, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x12, 0x48, 0x0a, 0x12, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x71,
	0x6c, 0x5f, 0x70, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x53, 0x71, 0x6c, 0x50, 0x61, 0x72, 0x73, 0x69, 0x6e, 0x67, 0x12, 0x23, 0x0a,
	0x0d, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x73, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x53,
	0x73, 0x6c, 0x12, 0x71, 0x0a, 0x0c, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x73,
	0x73, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x67, 0x72, 0x65, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x53, 0x53, 0x4c, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x53, 0x73, 0x6c, 0x22, 0x23, 0x0a, 0x07, 0x53, 0x53, 0x4c, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x10, 0x01, 0x42, 0xc7, 0x01, 0xba, 0x80, 0xc8,
	0xd1, 0x06, 0x04, 0x08, 0x01, 0x10, 0x02, 0x0a, 0x45, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x12,
	0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x70, 0x6f,
	0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x33, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_rawDescOnce sync.Once
	file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_rawDescData = file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_rawDesc
)

func file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_rawDescGZIP() []byte {
	file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_rawDescOnce.Do(func() {
		file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_rawDescData)
	})
	return file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_rawDescData
}

var file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_goTypes = []interface{}{
	(PostgresProxy_SSLMode)(0), // 0: envoy.extensions.filters.network.postgres_proxy.v3alpha.PostgresProxy.SSLMode
	(*PostgresProxy)(nil),      // 1: envoy.extensions.filters.network.postgres_proxy.v3alpha.PostgresProxy
	(*wrappers.BoolValue)(nil), // 2: google.protobuf.BoolValue
}
var file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.filters.network.postgres_proxy.v3alpha.PostgresProxy.enable_sql_parsing:type_name -> google.protobuf.BoolValue
	0, // 1: envoy.extensions.filters.network.postgres_proxy.v3alpha.PostgresProxy.upstream_ssl:type_name -> envoy.extensions.filters.network.postgres_proxy.v3alpha.PostgresProxy.SSLMode
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_init()
}
func file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_init() {
	if File_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostgresProxy); i {
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
			RawDescriptor: file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_goTypes,
		DependencyIndexes: file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_depIdxs,
		EnumInfos:         file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_enumTypes,
		MessageInfos:      file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_msgTypes,
	}.Build()
	File_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto = out.File
	file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_rawDesc = nil
	file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_goTypes = nil
	file_contrib_envoy_extensions_filters_network_postgres_proxy_v3alpha_postgres_proxy_proto_depIdxs = nil
}
