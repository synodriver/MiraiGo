// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: oidb0x8a7.proto

package oidb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type D8A7ReqBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubCmd                    *uint32 `protobuf:"varint,1,opt,name=subCmd" json:"subCmd,omitempty"`
	LimitIntervalTypeForUin   *uint32 `protobuf:"varint,2,opt,name=limitIntervalTypeForUin" json:"limitIntervalTypeForUin,omitempty"`
	LimitIntervalTypeForGroup *uint32 `protobuf:"varint,3,opt,name=limitIntervalTypeForGroup" json:"limitIntervalTypeForGroup,omitempty"`
	Uin                       *uint64 `protobuf:"varint,4,opt,name=uin" json:"uin,omitempty"`
	GroupCode                 *uint64 `protobuf:"varint,5,opt,name=groupCode" json:"groupCode,omitempty"`
}

func (x *D8A7ReqBody) Reset() {
	*x = D8A7ReqBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0x8a7_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *D8A7ReqBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*D8A7ReqBody) ProtoMessage() {}

func (x *D8A7ReqBody) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0x8a7_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use D8A7ReqBody.ProtoReflect.Descriptor instead.
func (*D8A7ReqBody) Descriptor() ([]byte, []int) {
	return file_oidb0x8a7_proto_rawDescGZIP(), []int{0}
}

func (x *D8A7ReqBody) GetSubCmd() uint32 {
	if x != nil && x.SubCmd != nil {
		return *x.SubCmd
	}
	return 0
}

func (x *D8A7ReqBody) GetLimitIntervalTypeForUin() uint32 {
	if x != nil && x.LimitIntervalTypeForUin != nil {
		return *x.LimitIntervalTypeForUin
	}
	return 0
}

func (x *D8A7ReqBody) GetLimitIntervalTypeForGroup() uint32 {
	if x != nil && x.LimitIntervalTypeForGroup != nil {
		return *x.LimitIntervalTypeForGroup
	}
	return 0
}

func (x *D8A7ReqBody) GetUin() uint64 {
	if x != nil && x.Uin != nil {
		return *x.Uin
	}
	return 0
}

func (x *D8A7ReqBody) GetGroupCode() uint64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

type D8A7RspBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CanAtAll                 *bool   `protobuf:"varint,1,opt,name=canAtAll" json:"canAtAll,omitempty"`
	RemainAtAllCountForUin   *uint32 `protobuf:"varint,2,opt,name=remainAtAllCountForUin" json:"remainAtAllCountForUin,omitempty"`
	RemainAtAllCountForGroup *uint32 `protobuf:"varint,3,opt,name=remainAtAllCountForGroup" json:"remainAtAllCountForGroup,omitempty"`
	PromptMsg1               []byte  `protobuf:"bytes,4,opt,name=promptMsg1" json:"promptMsg1,omitempty"`
	PromptMsg2               []byte  `protobuf:"bytes,5,opt,name=promptMsg2" json:"promptMsg2,omitempty"`
}

func (x *D8A7RspBody) Reset() {
	*x = D8A7RspBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oidb0x8a7_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *D8A7RspBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*D8A7RspBody) ProtoMessage() {}

func (x *D8A7RspBody) ProtoReflect() protoreflect.Message {
	mi := &file_oidb0x8a7_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use D8A7RspBody.ProtoReflect.Descriptor instead.
func (*D8A7RspBody) Descriptor() ([]byte, []int) {
	return file_oidb0x8a7_proto_rawDescGZIP(), []int{1}
}

func (x *D8A7RspBody) GetCanAtAll() bool {
	if x != nil && x.CanAtAll != nil {
		return *x.CanAtAll
	}
	return false
}

func (x *D8A7RspBody) GetRemainAtAllCountForUin() uint32 {
	if x != nil && x.RemainAtAllCountForUin != nil {
		return *x.RemainAtAllCountForUin
	}
	return 0
}

func (x *D8A7RspBody) GetRemainAtAllCountForGroup() uint32 {
	if x != nil && x.RemainAtAllCountForGroup != nil {
		return *x.RemainAtAllCountForGroup
	}
	return 0
}

func (x *D8A7RspBody) GetPromptMsg1() []byte {
	if x != nil {
		return x.PromptMsg1
	}
	return nil
}

func (x *D8A7RspBody) GetPromptMsg2() []byte {
	if x != nil {
		return x.PromptMsg2
	}
	return nil
}

var File_oidb0x8a7_proto protoreflect.FileDescriptor

var file_oidb0x8a7_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6f, 0x69, 0x64, 0x62, 0x30, 0x78, 0x38, 0x61, 0x37, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x0b, 0x44, 0x38, 0x41, 0x37, 0x52, 0x65, 0x71, 0x42, 0x6f, 0x64,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x43, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x73, 0x75, 0x62, 0x43, 0x6d, 0x64, 0x12, 0x38, 0x0a, 0x17, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x46, 0x6f,
	0x72, 0x55, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x17, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x46, 0x6f, 0x72,
	0x55, 0x69, 0x6e, 0x12, 0x3c, 0x0a, 0x19, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x46, 0x6f, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x19, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x46, 0x6f, 0x72, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x75, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x64,
	0x65, 0x22, 0xdd, 0x01, 0x0a, 0x0b, 0x44, 0x38, 0x41, 0x37, 0x52, 0x73, 0x70, 0x42, 0x6f, 0x64,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6e, 0x41, 0x74, 0x41, 0x6c, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x61, 0x6e, 0x41, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x36, 0x0a,
	0x16, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x41, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x46, 0x6f, 0x72, 0x55, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x72,
	0x65, 0x6d, 0x61, 0x69, 0x6e, 0x41, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x46,
	0x6f, 0x72, 0x55, 0x69, 0x6e, 0x12, 0x3a, 0x0a, 0x18, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x41,
	0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x72, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x18, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x41,
	0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x72, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x4d, 0x73, 0x67, 0x31, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x4d, 0x73, 0x67,
	0x31, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x4d, 0x73, 0x67, 0x32, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x4d, 0x73, 0x67,
	0x32, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x6f, 0x69, 0x64, 0x62,
}

var (
	file_oidb0x8a7_proto_rawDescOnce sync.Once
	file_oidb0x8a7_proto_rawDescData = file_oidb0x8a7_proto_rawDesc
)

func file_oidb0x8a7_proto_rawDescGZIP() []byte {
	file_oidb0x8a7_proto_rawDescOnce.Do(func() {
		file_oidb0x8a7_proto_rawDescData = protoimpl.X.CompressGZIP(file_oidb0x8a7_proto_rawDescData)
	})
	return file_oidb0x8a7_proto_rawDescData
}

var file_oidb0x8a7_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_oidb0x8a7_proto_goTypes = []interface{}{
	(*D8A7ReqBody)(nil), // 0: D8A7ReqBody
	(*D8A7RspBody)(nil), // 1: D8A7RspBody
}
var file_oidb0x8a7_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_oidb0x8a7_proto_init() }
func file_oidb0x8a7_proto_init() {
	if File_oidb0x8a7_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oidb0x8a7_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*D8A7ReqBody); i {
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
		file_oidb0x8a7_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*D8A7RspBody); i {
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
			RawDescriptor: file_oidb0x8a7_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_oidb0x8a7_proto_goTypes,
		DependencyIndexes: file_oidb0x8a7_proto_depIdxs,
		MessageInfos:      file_oidb0x8a7_proto_msgTypes,
	}.Build()
	File_oidb0x8a7_proto = out.File
	file_oidb0x8a7_proto_rawDesc = nil
	file_oidb0x8a7_proto_goTypes = nil
	file_oidb0x8a7_proto_depIdxs = nil
}
