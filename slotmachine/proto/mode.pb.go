// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: proto/mode.proto

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

type ModeUsageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AreaCode string `protobuf:"bytes,1,opt,name=area_code,json=areaCode,proto3" json:"area_code,omitempty"`
}

func (x *ModeUsageRequest) Reset() {
	*x = ModeUsageRequest{}
	mi := &file_proto_mode_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ModeUsageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModeUsageRequest) ProtoMessage() {}

func (x *ModeUsageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mode_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModeUsageRequest.ProtoReflect.Descriptor instead.
func (*ModeUsageRequest) Descriptor() ([]byte, []int) {
	return file_proto_mode_proto_rawDescGZIP(), []int{0}
}

func (x *ModeUsageRequest) GetAreaCode() string {
	if x != nil {
		return x.AreaCode
	}
	return ""
}

type ModeUsageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Modes []*ModeInfo `protobuf:"bytes,1,rep,name=modes,proto3" json:"modes,omitempty"`
}

func (x *ModeUsageResponse) Reset() {
	*x = ModeUsageResponse{}
	mi := &file_proto_mode_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ModeUsageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModeUsageResponse) ProtoMessage() {}

func (x *ModeUsageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mode_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModeUsageResponse.ProtoReflect.Descriptor instead.
func (*ModeUsageResponse) Descriptor() ([]byte, []int) {
	return file_proto_mode_proto_rawDescGZIP(), []int{1}
}

func (x *ModeUsageResponse) GetModes() []*ModeInfo {
	if x != nil {
		return x.Modes
	}
	return nil
}

type ModeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode  string `protobuf:"bytes,1,opt,name=mode,proto3" json:"mode,omitempty"`
	Usage int32  `protobuf:"varint,2,opt,name=usage,proto3" json:"usage,omitempty"`
}

func (x *ModeInfo) Reset() {
	*x = ModeInfo{}
	mi := &file_proto_mode_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ModeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModeInfo) ProtoMessage() {}

func (x *ModeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mode_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModeInfo.ProtoReflect.Descriptor instead.
func (*ModeInfo) Descriptor() ([]byte, []int) {
	return file_proto_mode_proto_rawDescGZIP(), []int{2}
}

func (x *ModeInfo) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *ModeInfo) GetUsage() int32 {
	if x != nil {
		return x.Usage
	}
	return 0
}

var File_proto_mode_proto protoreflect.FileDescriptor

var file_proto_mode_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x10, 0x4d, 0x6f, 0x64,
	0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x3a, 0x0a, 0x11, 0x4d, 0x6f,
	0x64, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x25, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x05, 0x6d, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x34, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x42, 0x13, 0x5a, 0x11,
	0x73, 0x6c, 0x6f, 0x74, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_mode_proto_rawDescOnce sync.Once
	file_proto_mode_proto_rawDescData = file_proto_mode_proto_rawDesc
)

func file_proto_mode_proto_rawDescGZIP() []byte {
	file_proto_mode_proto_rawDescOnce.Do(func() {
		file_proto_mode_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_mode_proto_rawDescData)
	})
	return file_proto_mode_proto_rawDescData
}

var file_proto_mode_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_mode_proto_goTypes = []any{
	(*ModeUsageRequest)(nil),  // 0: proto.ModeUsageRequest
	(*ModeUsageResponse)(nil), // 1: proto.ModeUsageResponse
	(*ModeInfo)(nil),          // 2: proto.ModeInfo
}
var file_proto_mode_proto_depIdxs = []int32{
	2, // 0: proto.ModeUsageResponse.modes:type_name -> proto.ModeInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_mode_proto_init() }
func file_proto_mode_proto_init() {
	if File_proto_mode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_mode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_mode_proto_goTypes,
		DependencyIndexes: file_proto_mode_proto_depIdxs,
		MessageInfos:      file_proto_mode_proto_msgTypes,
	}.Build()
	File_proto_mode_proto = out.File
	file_proto_mode_proto_rawDesc = nil
	file_proto_mode_proto_goTypes = nil
	file_proto_mode_proto_depIdxs = nil
}
