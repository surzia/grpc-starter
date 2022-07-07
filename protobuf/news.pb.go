// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: protobuf/news.proto

package protobuf

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Page     int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size     int32  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	IsFilter int32  `protobuf:"varint,4,opt,name=is_filter,json=isFilter,proto3" json:"is_filter,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_news_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_news_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_protobuf_news_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Request) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Request) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Request) GetIsFilter() int32 {
	if x != nil {
		return x.IsFilter
	}
	return 0
}

type News struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	News []*New `protobuf:"bytes,1,rep,name=news,proto3" json:"news,omitempty"`
}

func (x *News) Reset() {
	*x = News{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_news_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *News) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*News) ProtoMessage() {}

func (x *News) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_news_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use News.ProtoReflect.Descriptor instead.
func (*News) Descriptor() ([]byte, []int) {
	return file_protobuf_news_proto_rawDescGZIP(), []int{1}
}

func (x *News) GetNews() []*New {
	if x != nil {
		return x.News
	}
	return nil
}

type New struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uniquekey     string               `protobuf:"bytes,1,opt,name=uniquekey,proto3" json:"uniquekey,omitempty"`
	Title         string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Date          *timestamp.Timestamp `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Category      string               `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	AuthorName    string               `protobuf:"bytes,5,opt,name=author_name,json=authorName,proto3" json:"author_name,omitempty"`
	Url           string               `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	ThumbnailPicS string               `protobuf:"bytes,7,opt,name=thumbnail_pic_s,json=thumbnailPicS,proto3" json:"thumbnail_pic_s,omitempty"`
	IsContent     int32                `protobuf:"varint,8,opt,name=is_content,json=isContent,proto3" json:"is_content,omitempty"`
}

func (x *New) Reset() {
	*x = New{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_news_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *New) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*New) ProtoMessage() {}

func (x *New) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_news_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use New.ProtoReflect.Descriptor instead.
func (*New) Descriptor() ([]byte, []int) {
	return file_protobuf_news_proto_rawDescGZIP(), []int{2}
}

func (x *New) GetUniquekey() string {
	if x != nil {
		return x.Uniquekey
	}
	return ""
}

func (x *New) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *New) GetDate() *timestamp.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *New) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *New) GetAuthorName() string {
	if x != nil {
		return x.AuthorName
	}
	return ""
}

func (x *New) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *New) GetThumbnailPicS() string {
	if x != nil {
		return x.ThumbnailPicS
	}
	return ""
}

func (x *New) GetIsContent() int32 {
	if x != nil {
		return x.IsContent
	}
	return 0
}

var File_protobuf_news_proto protoreflect.FileDescriptor

var file_protobuf_news_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x62, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x73, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x22, 0x29, 0x0a, 0x04, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x21, 0x0a, 0x04,
	0x6e, 0x65, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4e, 0x65, 0x77, 0x52, 0x04, 0x6e, 0x65, 0x77, 0x73, 0x22,
	0xff, 0x01, 0x0a, 0x03, 0x4e, 0x65, 0x77, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x6e, 0x69, 0x71,
	0x75, 0x65, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x70, 0x69, 0x63, 0x5f, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x50, 0x69,
	0x63, 0x53, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x32, 0x42, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x34, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x74, 0x54, 0x6f, 0x70, 0x4e, 0x65, 0x77, 0x73,
	0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4e,
	0x65, 0x77, 0x73, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_news_proto_rawDescOnce sync.Once
	file_protobuf_news_proto_rawDescData = file_protobuf_news_proto_rawDesc
)

func file_protobuf_news_proto_rawDescGZIP() []byte {
	file_protobuf_news_proto_rawDescOnce.Do(func() {
		file_protobuf_news_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_news_proto_rawDescData)
	})
	return file_protobuf_news_proto_rawDescData
}

var file_protobuf_news_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protobuf_news_proto_goTypes = []interface{}{
	(*Request)(nil),             // 0: protobuf.Request
	(*News)(nil),                // 1: protobuf.News
	(*New)(nil),                 // 2: protobuf.New
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_protobuf_news_proto_depIdxs = []int32{
	2, // 0: protobuf.News.news:type_name -> protobuf.New
	3, // 1: protobuf.New.date:type_name -> google.protobuf.Timestamp
	0, // 2: protobuf.NewService.GetHotTopNews:input_type -> protobuf.Request
	1, // 3: protobuf.NewService.GetHotTopNews:output_type -> protobuf.News
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protobuf_news_proto_init() }
func file_protobuf_news_proto_init() {
	if File_protobuf_news_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_news_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_protobuf_news_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*News); i {
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
		file_protobuf_news_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*New); i {
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
			RawDescriptor: file_protobuf_news_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_news_proto_goTypes,
		DependencyIndexes: file_protobuf_news_proto_depIdxs,
		MessageInfos:      file_protobuf_news_proto_msgTypes,
	}.Build()
	File_protobuf_news_proto = out.File
	file_protobuf_news_proto_rawDesc = nil
	file_protobuf_news_proto_goTypes = nil
	file_protobuf_news_proto_depIdxs = nil
}
