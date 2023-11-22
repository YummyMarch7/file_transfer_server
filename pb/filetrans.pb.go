// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: pb/filetrans.proto

package file_transfer_server_pb

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

type Path struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path  string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Depth int32  `protobuf:"varint,2,opt,name=depth,proto3" json:"depth,omitempty"`
}

func (x *Path) Reset() {
	*x = Path{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_filetrans_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Path) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Path) ProtoMessage() {}

func (x *Path) ProtoReflect() protoreflect.Message {
	mi := &file_pb_filetrans_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Path.ProtoReflect.Descriptor instead.
func (*Path) Descriptor() ([]byte, []int) {
	return file_pb_filetrans_proto_rawDescGZIP(), []int{0}
}

func (x *Path) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Path) GetDepth() int32 {
	if x != nil {
		return x.Depth
	}
	return 0
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsFile bool    `protobuf:"varint,1,opt,name=is_file,json=isFile,proto3" json:"is_file,omitempty"`
	IsErr  bool    `protobuf:"varint,2,opt,name=is_err,json=isErr,proto3" json:"is_err,omitempty"`
	Name   string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Size   uint64  `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	SubDir []*Item `protobuf:"bytes,5,rep,name=sub_dir,json=subDir,proto3" json:"sub_dir,omitempty"`
	Files  []*Item `protobuf:"bytes,6,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_filetrans_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_pb_filetrans_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_pb_filetrans_proto_rawDescGZIP(), []int{1}
}

func (x *Item) GetIsFile() bool {
	if x != nil {
		return x.IsFile
	}
	return false
}

func (x *Item) GetIsErr() bool {
	if x != nil {
		return x.IsErr
	}
	return false
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Item) GetSubDir() []*Item {
	if x != nil {
		return x.SubDir
	}
	return nil
}

func (x *Item) GetFiles() []*Item {
	if x != nil {
		return x.Files
	}
	return nil
}

var File_pb_filetrans_proto protoreflect.FileDescriptor

var file_pb_filetrans_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x30, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x22, 0xbb, 0x01, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x69, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f,
	0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x45, 0x72, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x5f,
	0x64, 0x69, 0x72, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x06, 0x73, 0x75, 0x62, 0x44, 0x69, 0x72, 0x12, 0x2b, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x32, 0x46, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x12, 0x39, 0x0a, 0x07, 0x44, 0x69, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e,
	0x50, 0x61, 0x74, 0x68, 0x1a, 0x15, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x42, 0x19, 0x5a,
	0x17, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_filetrans_proto_rawDescOnce sync.Once
	file_pb_filetrans_proto_rawDescData = file_pb_filetrans_proto_rawDesc
)

func file_pb_filetrans_proto_rawDescGZIP() []byte {
	file_pb_filetrans_proto_rawDescOnce.Do(func() {
		file_pb_filetrans_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_filetrans_proto_rawDescData)
	})
	return file_pb_filetrans_proto_rawDescData
}

var file_pb_filetrans_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_filetrans_proto_goTypes = []interface{}{
	(*Path)(nil), // 0: file_trans.file.Path
	(*Item)(nil), // 1: file_trans.file.Item
}
var file_pb_filetrans_proto_depIdxs = []int32{
	1, // 0: file_trans.file.Item.sub_dir:type_name -> file_trans.file.Item
	1, // 1: file_trans.file.Item.files:type_name -> file_trans.file.Item
	0, // 2: file_trans.file.FileTrans.DirList:input_type -> file_trans.file.Path
	1, // 3: file_trans.file.FileTrans.DirList:output_type -> file_trans.file.Item
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pb_filetrans_proto_init() }
func file_pb_filetrans_proto_init() {
	if File_pb_filetrans_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_filetrans_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Path); i {
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
		file_pb_filetrans_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
			RawDescriptor: file_pb_filetrans_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_filetrans_proto_goTypes,
		DependencyIndexes: file_pb_filetrans_proto_depIdxs,
		MessageInfos:      file_pb_filetrans_proto_msgTypes,
	}.Build()
	File_pb_filetrans_proto = out.File
	file_pb_filetrans_proto_rawDesc = nil
	file_pb_filetrans_proto_goTypes = nil
	file_pb_filetrans_proto_depIdxs = nil
}