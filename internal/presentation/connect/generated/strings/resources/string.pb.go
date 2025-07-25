// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: string_phone/strings/resources/string.proto

package resources

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StringType int32

const (
	StringType_STRING_TYPE_UNSPECIFIED StringType = 0
	StringType_STRING_TYPE_A           StringType = 1
	StringType_STRING_TYPE_B           StringType = 2
	StringType_STRING_TYPE_C           StringType = 3
	StringType_STRING_TYPE_D           StringType = 4
	StringType_STRING_TYPE_E           StringType = 5
	StringType_STRING_TYPE_F           StringType = 6
)

// Enum value maps for StringType.
var (
	StringType_name = map[int32]string{
		0: "STRING_TYPE_UNSPECIFIED",
		1: "STRING_TYPE_A",
		2: "STRING_TYPE_B",
		3: "STRING_TYPE_C",
		4: "STRING_TYPE_D",
		5: "STRING_TYPE_E",
		6: "STRING_TYPE_F",
	}
	StringType_value = map[string]int32{
		"STRING_TYPE_UNSPECIFIED": 0,
		"STRING_TYPE_A":           1,
		"STRING_TYPE_B":           2,
		"STRING_TYPE_C":           3,
		"STRING_TYPE_D":           4,
		"STRING_TYPE_E":           5,
		"STRING_TYPE_F":           6,
	}
)

func (x StringType) Enum() *StringType {
	p := new(StringType)
	*p = x
	return p
}

func (x StringType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StringType) Descriptor() protoreflect.EnumDescriptor {
	return file_string_phone_strings_resources_string_proto_enumTypes[0].Descriptor()
}

func (StringType) Type() protoreflect.EnumType {
	return &file_string_phone_strings_resources_string_proto_enumTypes[0]
}

func (x StringType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StringType.Descriptor instead.
func (StringType) EnumDescriptor() ([]byte, []int) {
	return file_string_phone_strings_resources_string_proto_rawDescGZIP(), []int{0}
}

type TemplateString struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Id                string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DefaultDurability int32                  `protobuf:"varint,3,opt,name=default_durability,json=defaultDurability,proto3" json:"default_durability,omitempty"`
	Type              StringType             `protobuf:"varint,4,opt,name=type,proto3,enum=string_phone.strings.resources.StringType" json:"type,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *TemplateString) Reset() {
	*x = TemplateString{}
	mi := &file_string_phone_strings_resources_string_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TemplateString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateString) ProtoMessage() {}

func (x *TemplateString) ProtoReflect() protoreflect.Message {
	mi := &file_string_phone_strings_resources_string_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateString.ProtoReflect.Descriptor instead.
func (*TemplateString) Descriptor() ([]byte, []int) {
	return file_string_phone_strings_resources_string_proto_rawDescGZIP(), []int{0}
}

func (x *TemplateString) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TemplateString) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TemplateString) GetDefaultDurability() int32 {
	if x != nil {
		return x.DefaultDurability
	}
	return 0
}

func (x *TemplateString) GetType() StringType {
	if x != nil {
		return x.Type
	}
	return StringType_STRING_TYPE_UNSPECIFIED
}

type String struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Length        int32                  `protobuf:"varint,3,opt,name=length,proto3" json:"length,omitempty"`
	Durability    int32                  `protobuf:"varint,4,opt,name=durability,proto3" json:"durability,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Type          StringType             `protobuf:"varint,7,opt,name=type,proto3,enum=string_phone.strings.resources.StringType" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *String) Reset() {
	*x = String{}
	mi := &file_string_phone_strings_resources_string_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *String) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*String) ProtoMessage() {}

func (x *String) ProtoReflect() protoreflect.Message {
	mi := &file_string_phone_strings_resources_string_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use String.ProtoReflect.Descriptor instead.
func (*String) Descriptor() ([]byte, []int) {
	return file_string_phone_strings_resources_string_proto_rawDescGZIP(), []int{1}
}

func (x *String) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *String) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *String) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *String) GetDurability() int32 {
	if x != nil {
		return x.Durability
	}
	return 0
}

func (x *String) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *String) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *String) GetType() StringType {
	if x != nil {
		return x.Type
	}
	return StringType_STRING_TYPE_UNSPECIFIED
}

var File_string_phone_strings_resources_string_proto protoreflect.FileDescriptor

const file_string_phone_strings_resources_string_proto_rawDesc = "" +
	"\n" +
	"+string_phone/strings/resources/string.proto\x12\x1estring_phone.strings.resources\x1a\x1fgoogle/protobuf/timestamp.proto\"\xa3\x01\n" +
	"\x0eTemplateString\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12-\n" +
	"\x12default_durability\x18\x03 \x01(\x05R\x11defaultDurability\x12>\n" +
	"\x04type\x18\x04 \x01(\x0e2*.string_phone.strings.resources.StringTypeR\x04type\"\x9a\x02\n" +
	"\x06String\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x16\n" +
	"\x06length\x18\x03 \x01(\x05R\x06length\x12\x1e\n" +
	"\n" +
	"durability\x18\x04 \x01(\x05R\n" +
	"durability\x129\n" +
	"\n" +
	"created_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12>\n" +
	"\x04type\x18\a \x01(\x0e2*.string_phone.strings.resources.StringTypeR\x04type*\x9b\x01\n" +
	"\n" +
	"StringType\x12\x1b\n" +
	"\x17STRING_TYPE_UNSPECIFIED\x10\x00\x12\x11\n" +
	"\rSTRING_TYPE_A\x10\x01\x12\x11\n" +
	"\rSTRING_TYPE_B\x10\x02\x12\x11\n" +
	"\rSTRING_TYPE_C\x10\x03\x12\x11\n" +
	"\rSTRING_TYPE_D\x10\x04\x12\x11\n" +
	"\rSTRING_TYPE_E\x10\x05\x12\x11\n" +
	"\rSTRING_TYPE_F\x10\x06BfZdgithub.com/K-Kizuku/ito-denwa/internal/presentation/connect/generated/string_phone/strings/resourcesb\x06proto3"

var (
	file_string_phone_strings_resources_string_proto_rawDescOnce sync.Once
	file_string_phone_strings_resources_string_proto_rawDescData []byte
)

func file_string_phone_strings_resources_string_proto_rawDescGZIP() []byte {
	file_string_phone_strings_resources_string_proto_rawDescOnce.Do(func() {
		file_string_phone_strings_resources_string_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_string_phone_strings_resources_string_proto_rawDesc), len(file_string_phone_strings_resources_string_proto_rawDesc)))
	})
	return file_string_phone_strings_resources_string_proto_rawDescData
}

var file_string_phone_strings_resources_string_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_string_phone_strings_resources_string_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_string_phone_strings_resources_string_proto_goTypes = []any{
	(StringType)(0),               // 0: string_phone.strings.resources.StringType
	(*TemplateString)(nil),        // 1: string_phone.strings.resources.TemplateString
	(*String)(nil),                // 2: string_phone.strings.resources.String
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_string_phone_strings_resources_string_proto_depIdxs = []int32{
	0, // 0: string_phone.strings.resources.TemplateString.type:type_name -> string_phone.strings.resources.StringType
	3, // 1: string_phone.strings.resources.String.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: string_phone.strings.resources.String.updated_at:type_name -> google.protobuf.Timestamp
	0, // 3: string_phone.strings.resources.String.type:type_name -> string_phone.strings.resources.StringType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_string_phone_strings_resources_string_proto_init() }
func file_string_phone_strings_resources_string_proto_init() {
	if File_string_phone_strings_resources_string_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_string_phone_strings_resources_string_proto_rawDesc), len(file_string_phone_strings_resources_string_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_string_phone_strings_resources_string_proto_goTypes,
		DependencyIndexes: file_string_phone_strings_resources_string_proto_depIdxs,
		EnumInfos:         file_string_phone_strings_resources_string_proto_enumTypes,
		MessageInfos:      file_string_phone_strings_resources_string_proto_msgTypes,
	}.Build()
	File_string_phone_strings_resources_string_proto = out.File
	file_string_phone_strings_resources_string_proto_goTypes = nil
	file_string_phone_strings_resources_string_proto_depIdxs = nil
}
