// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto_files/user.proto

package proto_files

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

type UserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *UserID) Reset() {
	*x = UserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserID) ProtoMessage() {}

func (x *UserID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserID.ProtoReflect.Descriptor instead.
func (*UserID) Descriptor() ([]byte, []int) {
	return file_proto_files_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserID) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type UserIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID *UserID `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *UserIDRequest) Reset() {
	*x = UserIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIDRequest) ProtoMessage() {}

func (x *UserIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIDRequest.ProtoReflect.Descriptor instead.
func (*UserIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_files_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserIDRequest) GetUserID() *UserID {
	if x != nil {
		return x.UserID
	}
	return nil
}

type UserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IDUser   *UserID `protobuf:"bytes,1,opt,name=IDUser,proto3" json:"IDUser,omitempty"`
	Username string  `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *UserData) Reset() {
	*x = UserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_proto_files_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserData) GetIDUser() *UserID {
	if x != nil {
		return x.IDUser
	}
	return nil
}

func (x *UserData) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *UserData `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	Info *Info     `protobuf:"bytes,2,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_proto_files_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserResponse) GetUser() *UserData {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserResponse) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_proto_files_user_proto protoreflect.FileDescriptor

var file_proto_files_user_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x18, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x30, 0x0a, 0x0d,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x47,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x06, 0x49, 0x44,
	0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x52, 0x06, 0x49, 0x44, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x69, 0x68, 0x6e, 0x65, 0x61, 0x31, 0x37, 0x31, 0x31, 0x2f, 0x50, 0x4f, 0x53, 0x5f, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x72, 0x61, 0x62, 0x62, 0x69, 0x74, 0x2f, 0x69, 0x64, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_files_user_proto_rawDescOnce sync.Once
	file_proto_files_user_proto_rawDescData = file_proto_files_user_proto_rawDesc
)

func file_proto_files_user_proto_rawDescGZIP() []byte {
	file_proto_files_user_proto_rawDescOnce.Do(func() {
		file_proto_files_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_files_user_proto_rawDescData)
	})
	return file_proto_files_user_proto_rawDescData
}

var file_proto_files_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_files_user_proto_goTypes = []interface{}{
	(*UserID)(nil),        // 0: UserID
	(*UserIDRequest)(nil), // 1: UserIDRequest
	(*UserData)(nil),      // 2: UserData
	(*UserResponse)(nil),  // 3: UserResponse
	(*Info)(nil),          // 4: Info
}
var file_proto_files_user_proto_depIdxs = []int32{
	0, // 0: UserIDRequest.UserID:type_name -> UserID
	0, // 1: UserData.IDUser:type_name -> UserID
	2, // 2: UserResponse.User:type_name -> UserData
	4, // 3: UserResponse.Info:type_name -> Info
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_files_user_proto_init() }
func file_proto_files_user_proto_init() {
	if File_proto_files_user_proto != nil {
		return
	}
	file_proto_files_helper_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_files_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserID); i {
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
		file_proto_files_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIDRequest); i {
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
		file_proto_files_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserData); i {
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
		file_proto_files_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
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
			RawDescriptor: file_proto_files_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_files_user_proto_goTypes,
		DependencyIndexes: file_proto_files_user_proto_depIdxs,
		MessageInfos:      file_proto_files_user_proto_msgTypes,
	}.Build()
	File_proto_files_user_proto = out.File
	file_proto_files_user_proto_rawDesc = nil
	file_proto_files_user_proto_goTypes = nil
	file_proto_files_user_proto_depIdxs = nil
}
