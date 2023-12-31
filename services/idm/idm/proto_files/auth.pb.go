// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto_files/auth.proto

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

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserCredentials *UserCredenetials `protobuf:"bytes,1,opt,name=UserCredentials,proto3" json:"UserCredentials,omitempty"`
	Role            string            `protobuf:"bytes,2,opt,name=Role,proto3" json:"Role,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_proto_files_auth_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetUserCredentials() *UserCredenetials {
	if x != nil {
		return x.UserCredentials
	}
	return nil
}

func (x *RegisterRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserCredentials *UserCredenetials `protobuf:"bytes,1,opt,name=UserCredentials,proto3" json:"UserCredentials,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_proto_files_auth_proto_rawDescGZIP(), []int{1}
}

func (x *LoginRequest) GetUserCredentials() *UserCredenetials {
	if x != nil {
		return x.UserCredentials
	}
	return nil
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	Info  *Info  `protobuf:"bytes,2,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_proto_files_auth_proto_rawDescGZIP(), []int{2}
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginResponse) GetInfo() *Info {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_proto_files_auth_proto protoreflect.FileDescriptor

var file_proto_files_auth_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x0f, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a,
	0x0f, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x65, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x4b,
	0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b,
	0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x65, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x0f, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x22, 0x40, 0x0a, 0x0d, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x19, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x40, 0x5a,
	0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x68, 0x6e,
	0x65, 0x61, 0x31, 0x37, 0x31, 0x31, 0x2f, 0x50, 0x4f, 0x53, 0x5f, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x69, 0x64, 0x6d, 0x2f,
	0x69, 0x64, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_files_auth_proto_rawDescOnce sync.Once
	file_proto_files_auth_proto_rawDescData = file_proto_files_auth_proto_rawDesc
)

func file_proto_files_auth_proto_rawDescGZIP() []byte {
	file_proto_files_auth_proto_rawDescOnce.Do(func() {
		file_proto_files_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_files_auth_proto_rawDescData)
	})
	return file_proto_files_auth_proto_rawDescData
}

var file_proto_files_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_files_auth_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),  // 0: RegisterRequest
	(*LoginRequest)(nil),     // 1: LoginRequest
	(*LoginResponse)(nil),    // 2: LoginResponse
	(*UserCredenetials)(nil), // 3: UserCredenetials
	(*Info)(nil),             // 4: Info
}
var file_proto_files_auth_proto_depIdxs = []int32{
	3, // 0: RegisterRequest.UserCredentials:type_name -> UserCredenetials
	3, // 1: LoginRequest.UserCredentials:type_name -> UserCredenetials
	4, // 2: LoginResponse.Info:type_name -> Info
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_files_auth_proto_init() }
func file_proto_files_auth_proto_init() {
	if File_proto_files_auth_proto != nil {
		return
	}
	file_proto_files_helper_proto_init()
	file_proto_files_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_files_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_proto_files_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_proto_files_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
			RawDescriptor: file_proto_files_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_files_auth_proto_goTypes,
		DependencyIndexes: file_proto_files_auth_proto_depIdxs,
		MessageInfos:      file_proto_files_auth_proto_msgTypes,
	}.Build()
	File_proto_files_auth_proto = out.File
	file_proto_files_auth_proto_rawDesc = nil
	file_proto_files_auth_proto_goTypes = nil
	file_proto_files_auth_proto_depIdxs = nil
}
