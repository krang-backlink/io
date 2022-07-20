// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: backlink/linksnatcher.proto

package backlink

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

type Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url  string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Mode string `protobuf:"bytes,2,opt,name=mode,proto3" json:"mode,omitempty"`
}

func (x *Req) Reset() {
	*x = Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backlink_linksnatcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req) ProtoMessage() {}

func (x *Req) ProtoReflect() protoreflect.Message {
	mi := &file_backlink_linksnatcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req.ProtoReflect.Descriptor instead.
func (*Req) Descriptor() ([]byte, []int) {
	return file_backlink_linksnatcher_proto_rawDescGZIP(), []int{0}
}

func (x *Req) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Req) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Backlinks []*Response_Links `protobuf:"bytes,1,rep,name=Backlinks,proto3" json:"Backlinks,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backlink_linksnatcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_backlink_linksnatcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_backlink_linksnatcher_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetBacklinks() []*Response_Links {
	if x != nil {
		return x.Backlinks
	}
	return nil
}

type Response_Links struct {
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

func (x *Response_Links) Reset() {
	*x = Response_Links{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backlink_linksnatcher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response_Links) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response_Links) ProtoMessage() {}

func (x *Response_Links) ProtoReflect() protoreflect.Message {
	mi := &file_backlink_linksnatcher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response_Links.ProtoReflect.Descriptor instead.
func (*Response_Links) Descriptor() ([]byte, []int) {
	return file_backlink_linksnatcher_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Response_Links) GetPageTitle() string {
	if x != nil {
		return x.PageTitle
	}
	return ""
}

func (x *Response_Links) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Response_Links) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Response_Links) GetAnchorAlt() string {
	if x != nil {
		return x.AnchorAlt
	}
	return ""
}

func (x *Response_Links) GetIsNofollow() bool {
	if x != nil {
		return x.IsNofollow
	}
	return false
}

func (x *Response_Links) GetDr() float32 {
	if x != nil {
		return x.Dr
	}
	return 0
}

var File_backlink_linksnatcher_proto protoreflect.FileDescriptor

var file_backlink_linksnatcher_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x62, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x73,
	0x6e, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6b,
	0x72, 0x61, 0x6e, 0x67, 0x22, 0x2b, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64,
	0x65, 0x22, 0xdb, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33,
	0x0a, 0x09, 0x42, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x6b, 0x72, 0x61, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x09, 0x42, 0x61, 0x63, 0x6b, 0x6c, 0x69,
	0x6e, 0x6b, 0x73, 0x1a, 0x99, 0x01, 0x0a, 0x05, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x50, 0x61, 0x67, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x50, 0x61, 0x67, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x55,
	0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x41, 0x6c, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x41, 0x6c, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x49, 0x73, 0x4e, 0x6f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x49, 0x73, 0x4e, 0x6f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12,
	0x0e, 0x0a, 0x02, 0x44, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x02, 0x44, 0x72, 0x32,
	0x41, 0x0a, 0x10, 0x42, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x6c, 0x69,
	0x6e, 0x6b, 0x73, 0x12, 0x0a, 0x2e, 0x6b, 0x72, 0x61, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x71, 0x1a,
	0x0f, 0x2e, 0x6b, 0x72, 0x61, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x72, 0x61, 0x6e, 0x67, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x6e, 0x6b, 0x2f,
	0x69, 0x6f, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x6e, 0x6b, 0x3b, 0x62, 0x61, 0x63, 0x6b,
	0x6c, 0x69, 0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backlink_linksnatcher_proto_rawDescOnce sync.Once
	file_backlink_linksnatcher_proto_rawDescData = file_backlink_linksnatcher_proto_rawDesc
)

func file_backlink_linksnatcher_proto_rawDescGZIP() []byte {
	file_backlink_linksnatcher_proto_rawDescOnce.Do(func() {
		file_backlink_linksnatcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_backlink_linksnatcher_proto_rawDescData)
	})
	return file_backlink_linksnatcher_proto_rawDescData
}

var file_backlink_linksnatcher_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_backlink_linksnatcher_proto_goTypes = []interface{}{
	(*Req)(nil),            // 0: krang.Req
	(*Response)(nil),       // 1: krang.Response
	(*Response_Links)(nil), // 2: krang.Response.Links
}
var file_backlink_linksnatcher_proto_depIdxs = []int32{
	2, // 0: krang.Response.Backlinks:type_name -> krang.Response.Links
	0, // 1: krang.BacklinksService.GetBacklinks:input_type -> krang.Req
	1, // 2: krang.BacklinksService.GetBacklinks:output_type -> krang.Response
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_backlink_linksnatcher_proto_init() }
func file_backlink_linksnatcher_proto_init() {
	if File_backlink_linksnatcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_backlink_linksnatcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req); i {
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
		file_backlink_linksnatcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_backlink_linksnatcher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response_Links); i {
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
			RawDescriptor: file_backlink_linksnatcher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_backlink_linksnatcher_proto_goTypes,
		DependencyIndexes: file_backlink_linksnatcher_proto_depIdxs,
		MessageInfos:      file_backlink_linksnatcher_proto_msgTypes,
	}.Build()
	File_backlink_linksnatcher_proto = out.File
	file_backlink_linksnatcher_proto_rawDesc = nil
	file_backlink_linksnatcher_proto_goTypes = nil
	file_backlink_linksnatcher_proto_depIdxs = nil
}
