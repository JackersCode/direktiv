// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pkg/ingress/get-namespaces.proto

package ingress

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type GetNamespacesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset *int32 `protobuf:"varint,1,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
	Limit  *int32 `protobuf:"varint,2,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
}

func (x *GetNamespacesRequest) Reset() {
	*x = GetNamespacesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_ingress_get_namespaces_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNamespacesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNamespacesRequest) ProtoMessage() {}

func (x *GetNamespacesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_ingress_get_namespaces_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNamespacesRequest.ProtoReflect.Descriptor instead.
func (*GetNamespacesRequest) Descriptor() ([]byte, []int) {
	return file_pkg_ingress_get_namespaces_proto_rawDescGZIP(), []int{0}
}

func (x *GetNamespacesRequest) GetOffset() int32 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

func (x *GetNamespacesRequest) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

type GetNamespacesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespaces []*GetNamespacesResponse_Namespace `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	Offset     *int32                             `protobuf:"varint,2,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
	Limit      *int32                             `protobuf:"varint,3,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
}

func (x *GetNamespacesResponse) Reset() {
	*x = GetNamespacesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_ingress_get_namespaces_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNamespacesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNamespacesResponse) ProtoMessage() {}

func (x *GetNamespacesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_ingress_get_namespaces_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNamespacesResponse.ProtoReflect.Descriptor instead.
func (*GetNamespacesResponse) Descriptor() ([]byte, []int) {
	return file_pkg_ingress_get_namespaces_proto_rawDescGZIP(), []int{1}
}

func (x *GetNamespacesResponse) GetNamespaces() []*GetNamespacesResponse_Namespace {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

func (x *GetNamespacesResponse) GetOffset() int32 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

func (x *GetNamespacesResponse) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

type GetNamespacesResponse_Namespace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      *string              `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=createdAt,proto3,oneof" json:"createdAt,omitempty"`
}

func (x *GetNamespacesResponse_Namespace) Reset() {
	*x = GetNamespacesResponse_Namespace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_ingress_get_namespaces_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNamespacesResponse_Namespace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNamespacesResponse_Namespace) ProtoMessage() {}

func (x *GetNamespacesResponse_Namespace) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_ingress_get_namespaces_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNamespacesResponse_Namespace.ProtoReflect.Descriptor instead.
func (*GetNamespacesResponse_Namespace) Descriptor() ([]byte, []int) {
	return file_pkg_ingress_get_namespaces_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetNamespacesResponse_Namespace) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *GetNamespacesResponse_Namespace) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_pkg_ingress_get_namespaces_proto protoreflect.FileDescriptor

var file_pkg_ingress_get_namespaces_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x67, 0x65,
	0x74, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x22, 0xaa, 0x02, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0a, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x1a, 0x7a, 0x0a,
	0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x48, 0x01, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88,
	0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x29,
	0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x72,
	0x74, 0x65, 0x69, 0x6c, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_ingress_get_namespaces_proto_rawDescOnce sync.Once
	file_pkg_ingress_get_namespaces_proto_rawDescData = file_pkg_ingress_get_namespaces_proto_rawDesc
)

func file_pkg_ingress_get_namespaces_proto_rawDescGZIP() []byte {
	file_pkg_ingress_get_namespaces_proto_rawDescOnce.Do(func() {
		file_pkg_ingress_get_namespaces_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_ingress_get_namespaces_proto_rawDescData)
	})
	return file_pkg_ingress_get_namespaces_proto_rawDescData
}

var file_pkg_ingress_get_namespaces_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_ingress_get_namespaces_proto_goTypes = []interface{}{
	(*GetNamespacesRequest)(nil),            // 0: ingress.GetNamespacesRequest
	(*GetNamespacesResponse)(nil),           // 1: ingress.GetNamespacesResponse
	(*GetNamespacesResponse_Namespace)(nil), // 2: ingress.GetNamespacesResponse.Namespace
	(*timestamp.Timestamp)(nil),             // 3: google.protobuf.Timestamp
}
var file_pkg_ingress_get_namespaces_proto_depIdxs = []int32{
	2, // 0: ingress.GetNamespacesResponse.namespaces:type_name -> ingress.GetNamespacesResponse.Namespace
	3, // 1: ingress.GetNamespacesResponse.Namespace.createdAt:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_ingress_get_namespaces_proto_init() }
func file_pkg_ingress_get_namespaces_proto_init() {
	if File_pkg_ingress_get_namespaces_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_ingress_get_namespaces_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNamespacesRequest); i {
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
		file_pkg_ingress_get_namespaces_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNamespacesResponse); i {
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
		file_pkg_ingress_get_namespaces_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNamespacesResponse_Namespace); i {
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
	file_pkg_ingress_get_namespaces_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_pkg_ingress_get_namespaces_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_pkg_ingress_get_namespaces_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_ingress_get_namespaces_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_ingress_get_namespaces_proto_goTypes,
		DependencyIndexes: file_pkg_ingress_get_namespaces_proto_depIdxs,
		MessageInfos:      file_pkg_ingress_get_namespaces_proto_msgTypes,
	}.Build()
	File_pkg_ingress_get_namespaces_proto = out.File
	file_pkg_ingress_get_namespaces_proto_rawDesc = nil
	file_pkg_ingress_get_namespaces_proto_goTypes = nil
	file_pkg_ingress_get_namespaces_proto_depIdxs = nil
}
