// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: worker.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DispatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId     int64   `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	GroupSlug  string  `protobuf:"bytes,2,opt,name=group_slug,json=groupSlug,proto3" json:"group_slug,omitempty"`
	SearchTerm string  `protobuf:"bytes,3,opt,name=search_term,json=searchTerm,proto3" json:"search_term,omitempty"`
	Pages      []*Page `protobuf:"bytes,4,rep,name=pages,proto3" json:"pages,omitempty"`
}

func (x *DispatchRequest) Reset() {
	*x = DispatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchRequest) ProtoMessage() {}

func (x *DispatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchRequest.ProtoReflect.Descriptor instead.
func (*DispatchRequest) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{0}
}

func (x *DispatchRequest) GetTaskId() int64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *DispatchRequest) GetGroupSlug() string {
	if x != nil {
		return x.GroupSlug
	}
	return ""
}

func (x *DispatchRequest) GetSearchTerm() string {
	if x != nil {
		return x.SearchTerm
	}
	return ""
}

func (x *DispatchRequest) GetPages() []*Page {
	if x != nil {
		return x.Pages
	}
	return nil
}

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url         string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	SearchIterm string `protobuf:"bytes,2,opt,name=search_iterm,json=searchIterm,proto3" json:"search_iterm,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{1}
}

func (x *Page) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Page) GetSearchIterm() string {
	if x != nil {
		return x.SearchIterm
	}
	return ""
}

type CompletePageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ScrapeId   string `protobuf:"bytes,2,opt,name=scrape_id,json=scrapeId,proto3" json:"scrape_id,omitempty"`
	Url        string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	GroupSlug  string `protobuf:"bytes,4,opt,name=group_slug,json=groupSlug,proto3" json:"group_slug,omitempty"`
	TaskId     int64  `protobuf:"varint,5,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	SearchTerm string `protobuf:"bytes,6,opt,name=search_term,json=searchTerm,proto3" json:"search_term,omitempty"`
}

func (x *CompletePageRequest) Reset() {
	*x = CompletePageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletePageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletePageRequest) ProtoMessage() {}

func (x *CompletePageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_worker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletePageRequest.ProtoReflect.Descriptor instead.
func (*CompletePageRequest) Descriptor() ([]byte, []int) {
	return file_worker_proto_rawDescGZIP(), []int{2}
}

func (x *CompletePageRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CompletePageRequest) GetScrapeId() string {
	if x != nil {
		return x.ScrapeId
	}
	return ""
}

func (x *CompletePageRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CompletePageRequest) GetGroupSlug() string {
	if x != nil {
		return x.GroupSlug
	}
	return ""
}

func (x *CompletePageRequest) GetTaskId() int64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *CompletePageRequest) GetSearchTerm() string {
	if x != nil {
		return x.SearchTerm
	}
	return ""
}

var File_worker_proto protoreflect.FileDescriptor

var file_worker_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x6b, 0x72, 0x61, 0x6e, 0x67, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x6c, 0x75, 0x67, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x12,
	0x21, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x6b, 0x72, 0x61, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x73, 0x22, 0x3b, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x69, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x74, 0x65, 0x72, 0x6d, 0x22,
	0xad, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x63, 0x72, 0x61, 0x70,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x63, 0x72, 0x61,
	0x70, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x73, 0x6c, 0x75, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x53, 0x6c, 0x75, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x32,
	0x93, 0x01, 0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3c, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x16, 0x2e,
	0x6b, 0x72, 0x61, 0x6e, 0x67, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x44, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x12,
	0x1a, 0x2e, 0x6b, 0x72, 0x61, 0x6e, 0x67, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x72, 0x61, 0x6e, 0x67, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x6c, 0x69,
	0x6e, 0x6b, 0x2f, 0x69, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_worker_proto_rawDescOnce sync.Once
	file_worker_proto_rawDescData = file_worker_proto_rawDesc
)

func file_worker_proto_rawDescGZIP() []byte {
	file_worker_proto_rawDescOnce.Do(func() {
		file_worker_proto_rawDescData = protoimpl.X.CompressGZIP(file_worker_proto_rawDescData)
	})
	return file_worker_proto_rawDescData
}

var file_worker_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_worker_proto_goTypes = []interface{}{
	(*DispatchRequest)(nil),     // 0: krang.DispatchRequest
	(*Page)(nil),                // 1: krang.Page
	(*CompletePageRequest)(nil), // 2: krang.CompletePageRequest
	(*emptypb.Empty)(nil),       // 3: google.protobuf.Empty
}
var file_worker_proto_depIdxs = []int32{
	1, // 0: krang.DispatchRequest.pages:type_name -> krang.Page
	0, // 1: krang.WorkerService.Dispatch:input_type -> krang.DispatchRequest
	2, // 2: krang.WorkerService.CompletePage:input_type -> krang.CompletePageRequest
	3, // 3: krang.WorkerService.Dispatch:output_type -> google.protobuf.Empty
	3, // 4: krang.WorkerService.CompletePage:output_type -> google.protobuf.Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_worker_proto_init() }
func file_worker_proto_init() {
	if File_worker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_worker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DispatchRequest); i {
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
		file_worker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_worker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompletePageRequest); i {
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
			RawDescriptor: file_worker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_worker_proto_goTypes,
		DependencyIndexes: file_worker_proto_depIdxs,
		MessageInfos:      file_worker_proto_msgTypes,
	}.Build()
	File_worker_proto = out.File
	file_worker_proto_rawDesc = nil
	file_worker_proto_goTypes = nil
	file_worker_proto_depIdxs = nil
}
