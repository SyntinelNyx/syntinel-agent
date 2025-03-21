// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.5
// source: internal/proto/hardwareinfo.proto

package proto

import (
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

type HardwareInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JsonData string `protobuf:"bytes,1,opt,name=json_data,json=jsonData,proto3" json:"json_data,omitempty"` // The JSON string containing hardware info
}

func (x *HardwareInfoRequest) Reset() {
	*x = HardwareInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_hardwareinfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HardwareInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HardwareInfoRequest) ProtoMessage() {}

func (x *HardwareInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_hardwareinfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HardwareInfoRequest.ProtoReflect.Descriptor instead.
func (*HardwareInfoRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_hardwareinfo_proto_rawDescGZIP(), []int{0}
}

func (x *HardwareInfoRequest) GetJsonData() string {
	if x != nil {
		return x.JsonData
	}
	return ""
}

type HardwareResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // Response message from the server
}

func (x *HardwareResponse) Reset() {
	*x = HardwareResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_hardwareinfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HardwareResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HardwareResponse) ProtoMessage() {}

func (x *HardwareResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_hardwareinfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HardwareResponse.ProtoReflect.Descriptor instead.
func (*HardwareResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_hardwareinfo_proto_rawDescGZIP(), []int{1}
}

func (x *HardwareResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_internal_proto_hardwareinfo_proto protoreflect.FileDescriptor

var file_internal_proto_hardwareinfo_proto_rawDesc = []byte{
	0x0a, 0x21, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x32, 0x0a, 0x13, 0x48, 0x61, 0x72,
	0x64, 0x77, 0x61, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x2c, 0x0a,
	0x10, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x55, 0x0a, 0x0c, 0x48,
	0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x45, 0x0a, 0x10, 0x53,
	0x65, 0x6e, 0x64, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x53, 0x79, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6c, 0x4e, 0x79, 0x78, 0x2f, 0x73, 0x79, 0x6e,
	0x74, 0x69, 0x6e, 0x65, 0x6c, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_internal_proto_hardwareinfo_proto_rawDescOnce sync.Once
	file_internal_proto_hardwareinfo_proto_rawDescData = file_internal_proto_hardwareinfo_proto_rawDesc
)

func file_internal_proto_hardwareinfo_proto_rawDescGZIP() []byte {
	file_internal_proto_hardwareinfo_proto_rawDescOnce.Do(func() {
		file_internal_proto_hardwareinfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_hardwareinfo_proto_rawDescData)
	})
	return file_internal_proto_hardwareinfo_proto_rawDescData
}

var file_internal_proto_hardwareinfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_proto_hardwareinfo_proto_goTypes = []any{
	(*HardwareInfoRequest)(nil), // 0: grpc.HardwareInfoRequest
	(*HardwareResponse)(nil),    // 1: grpc.HardwareResponse
}
var file_internal_proto_hardwareinfo_proto_depIdxs = []int32{
	0, // 0: grpc.HardwareInfo.SendHardwareInfo:input_type -> grpc.HardwareInfoRequest
	1, // 1: grpc.HardwareInfo.SendHardwareInfo:output_type -> grpc.HardwareResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_proto_hardwareinfo_proto_init() }
func file_internal_proto_hardwareinfo_proto_init() {
	if File_internal_proto_hardwareinfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_hardwareinfo_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*HardwareInfoRequest); i {
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
		file_internal_proto_hardwareinfo_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*HardwareResponse); i {
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
			RawDescriptor: file_internal_proto_hardwareinfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_hardwareinfo_proto_goTypes,
		DependencyIndexes: file_internal_proto_hardwareinfo_proto_depIdxs,
		MessageInfos:      file_internal_proto_hardwareinfo_proto_msgTypes,
	}.Build()
	File_internal_proto_hardwareinfo_proto = out.File
	file_internal_proto_hardwareinfo_proto_rawDesc = nil
	file_internal_proto_hardwareinfo_proto_goTypes = nil
	file_internal_proto_hardwareinfo_proto_depIdxs = nil
}
