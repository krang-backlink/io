// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: backlink.proto

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

type BacklinksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url  string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Mode string `protobuf:"bytes,2,opt,name=mode,proto3" json:"mode,omitempty"`
}

func (x *BacklinksRequest) Reset() {
	*x = BacklinksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backlink_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BacklinksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BacklinksRequest) ProtoMessage() {}

func (x *BacklinksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backlink_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BacklinksRequest.ProtoReflect.Descriptor instead.
func (*BacklinksRequest) Descriptor() ([]byte, []int) {
	return file_backlink_proto_rawDescGZIP(), []int{0}
}

func (x *BacklinksRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *BacklinksRequest) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

type BacklinksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Backlinks []*Links `protobuf:"bytes,1,rep,name=Backlinks,proto3" json:"Backlinks,omitempty"`
}

func (x *BacklinksResponse) Reset() {
	*x = BacklinksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backlink_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BacklinksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BacklinksResponse) ProtoMessage() {}

func (x *BacklinksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backlink_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BacklinksResponse.ProtoReflect.Descriptor instead.
func (*BacklinksResponse) Descriptor() ([]byte, []int) {
	return file_backlink_proto_rawDescGZIP(), []int{1}
}

func (x *BacklinksResponse) GetBacklinks() []*Links {
	if x != nil {
		return x.Backlinks
	}
	return nil
}

type Links struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageTitle  string  `protobuf:"bytes,1,opt,name=PageTitle,proto3" json:"PageTitle,omitempty"`
	Url        string  `protobuf:"bytes,2,opt,name=Url,proto3" json:"Url,omitempty"`
	Type       string  `protobuf:"bytes,3,opt,name=Type,proto3" json:"Type,omitempty"`
	AnchorAlt  string  `protobuf:"bytes,4,opt,name=AnchorAlt,proto3" json:"AnchorAlt,omitempty"`
	IsNofollow bool    `protobuf:"varint,5,opt,name=IsNofollow,proto3" json:"IsNofollow,omitempty"`
	Dr         float32 `protobuf:"fixed32,6,opt,name=Dr,proto3" json:"Dr,omitempty"`
}

func (x *Links) Reset() {
	*x = Links{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backlink_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Links) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Links) ProtoMessage() {}

func (x *Links) ProtoReflect() protoreflect.Message {
	mi := &file_backlink_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Links.ProtoReflect.Descriptor instead.
func (*Links) Descriptor() ([]byte, []int) {
	return file_backlink_proto_rawDescGZIP(), []int{2}
}

func (x *Links) GetPageTitle() string {
	if x != nil {
		return x.PageTitle
	}
	return ""
}

func (x *Links) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Links) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Links) GetAnchorAlt() string {
	if x != nil {
		return x.AnchorAlt
	}
	return ""
}

func (x *Links) GetIsNofollow() bool {
	if x != nil {
		return x.IsNofollow
	}
	return false
}

func (x *Links) GetDr() float32 {
	if x != nil {
		return x.Dr
	}
	return 0
}

var File_backlink_proto protoreflect.FileDescriptor

var file_backlink_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x6b, 0x72, 0x61, 0x6e, 0x67, 0x22, 0x38, 0x0a, 0x10, 0x42, 0x61, 0x63, 0x6b, 0x6c,
	0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64,
	0x65, 0x22, 0x3f, 0x0a, 0x11, 0x42, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x09, 0x42, 0x61, 0x63, 0x6b, 0x6c, 0x69,
	0x6e, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6b, 0x72, 0x61, 0x6e,
	0x67, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x09, 0x42, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x6e,
	0x6b, 0x73, 0x22, 0x99, 0x01, 0x0a, 0x05, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x50, 0x61, 0x67, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x50, 0x61, 0x67, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x41, 0x6c, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x41, 0x6c, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x49, 0x73, 0x4e, 0x6f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x49, 0x73, 0x4e, 0x6f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x0e,
	0x0a, 0x02, 0x44, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x02, 0x44, 0x72, 0x32, 0x58,
	0x0a, 0x11, 0x42, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x6e, 0x6b, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x43, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x6c, 0x69,
	0x6e, 0x6b, 0x73, 0x12, 0x17, 0x2e, 0x6b, 0x72, 0x61, 0x6e, 0x67, 0x2e, 0x42, 0x61, 0x63, 0x6b,
	0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6b,
	0x72, 0x61, 0x6e, 0x67, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x72, 0x61, 0x6e, 0x67, 0x2d, 0x62, 0x61, 0x63,
	0x6b, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backlink_proto_rawDescOnce sync.Once
	file_backlink_proto_rawDescData = file_backlink_proto_rawDesc
)

func file_backlink_proto_rawDescGZIP() []byte {
	file_backlink_proto_rawDescOnce.Do(func() {
		file_backlink_proto_rawDescData = protoimpl.X.CompressGZIP(file_backlink_proto_rawDescData)
	})
	return file_backlink_proto_rawDescData
}

var file_backlink_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_backlink_proto_goTypes = []interface{}{
	(*BacklinksRequest)(nil),  // 0: krang.BacklinksRequest
	(*BacklinksResponse)(nil), // 1: krang.BacklinksResponse
	(*Links)(nil),             // 2: krang.Links
}
var file_backlink_proto_depIdxs = []int32{
	2, // 0: krang.BacklinksResponse.Backlinks:type_name -> krang.Links
	0, // 1: krang.BacklinkExtractor.GetBacklinks:input_type -> krang.BacklinksRequest
	1, // 2: krang.BacklinkExtractor.GetBacklinks:output_type -> krang.BacklinksResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_backlink_proto_init() }
func file_backlink_proto_init() {
	if File_backlink_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_backlink_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BacklinksRequest); i {
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
		file_backlink_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BacklinksResponse); i {
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
		file_backlink_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Links); i {
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
			RawDescriptor: file_backlink_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_backlink_proto_goTypes,
		DependencyIndexes: file_backlink_proto_depIdxs,
		MessageInfos:      file_backlink_proto_msgTypes,
	}.Build()
	File_backlink_proto = out.File
	file_backlink_proto_rawDesc = nil
	file_backlink_proto_goTypes = nil
	file_backlink_proto_depIdxs = nil
}