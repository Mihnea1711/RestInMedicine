// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: idm.proto

package idm

import (
	proto_files "github.com/mihnea1711/POS_Project/services/gateway/idm/proto_files"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_idm_proto protoreflect.FileDescriptor

var file_idm_proto_rawDesc = []byte{
	0x0a, 0x09, 0x69, 0x64, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc4, 0x05, 0x0a, 0x03, 0x49, 0x44, 0x4d, 0x12, 0x2d, 0x0a,
	0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x49, 0x44,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0d, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x0d, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0e,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x12,
	0x12, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x45, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0e, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x45,
	0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x0e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3b, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x45, 0x6e, 0x68, 0x61, 0x6e, 0x63,
	0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x10, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x45, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x12, 0x41,
	0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x11, 0x2e, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x14, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x17, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x6c, 0x61, 0x63,
	0x6b, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x45, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x38, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x68, 0x6e, 0x65,
	0x61, 0x31, 0x37, 0x31, 0x31, 0x2f, 0x50, 0x4f, 0x53, 0x5f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x69, 0x64, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_idm_proto_goTypes = []interface{}{
	(*proto_files.RegisterRequest)(nil),       // 0: RegisterRequest
	(*proto_files.LoginRequest)(nil),          // 1: LoginRequest
	(*proto_files.EmptyRequest)(nil),          // 2: EmptyRequest
	(*proto_files.UserIDRequest)(nil),         // 3: UserIDRequest
	(*proto_files.UpdateUserRequest)(nil),     // 4: UpdateUserRequest
	(*proto_files.UpdateRoleRequest)(nil),     // 5: UpdateRoleRequest
	(*proto_files.UsernameRequest)(nil),       // 6: UsernameRequest
	(*proto_files.UpdatePasswordRequest)(nil), // 7: UpdatePasswordRequest
	(*proto_files.BlacklistRequest)(nil),      // 8: BlacklistRequest
	(*proto_files.IDInfoResponse)(nil),        // 9: IDInfoResponse
	(*proto_files.LoginResponse)(nil),         // 10: LoginResponse
	(*proto_files.UsersResponse)(nil),         // 11: UsersResponse
	(*proto_files.UserResponse)(nil),          // 12: UserResponse
	(*proto_files.EnhancedInfoResponse)(nil),  // 13: EnhancedInfoResponse
	(*proto_files.RoleResponse)(nil),          // 14: RoleResponse
	(*proto_files.PasswordResponse)(nil),      // 15: PasswordResponse
	(*proto_files.InfoResponse)(nil),          // 16: InfoResponse
}
var file_idm_proto_depIdxs = []int32{
	0,  // 0: IDM.Register:input_type -> RegisterRequest
	1,  // 1: IDM.Login:input_type -> LoginRequest
	2,  // 2: IDM.GetUsers:input_type -> EmptyRequest
	3,  // 3: IDM.GetUserByID:input_type -> UserIDRequest
	4,  // 4: IDM.UpdateUserByID:input_type -> UpdateUserRequest
	3,  // 5: IDM.DeleteUserByID:input_type -> UserIDRequest
	3,  // 6: IDM.GetUserRole:input_type -> UserIDRequest
	5,  // 7: IDM.UpdateUserRole:input_type -> UpdateRoleRequest
	6,  // 8: IDM.GetUserPassword:input_type -> UsernameRequest
	7,  // 9: IDM.UpdateUserPassword:input_type -> UpdatePasswordRequest
	8,  // 10: IDM.AddUserToBlacklist:input_type -> BlacklistRequest
	3,  // 11: IDM.CheckUserInBlacklist:input_type -> UserIDRequest
	3,  // 12: IDM.RemoveUserFromBlacklist:input_type -> UserIDRequest
	9,  // 13: IDM.Register:output_type -> IDInfoResponse
	10, // 14: IDM.Login:output_type -> LoginResponse
	11, // 15: IDM.GetUsers:output_type -> UsersResponse
	12, // 16: IDM.GetUserByID:output_type -> UserResponse
	13, // 17: IDM.UpdateUserByID:output_type -> EnhancedInfoResponse
	13, // 18: IDM.DeleteUserByID:output_type -> EnhancedInfoResponse
	14, // 19: IDM.GetUserRole:output_type -> RoleResponse
	13, // 20: IDM.UpdateUserRole:output_type -> EnhancedInfoResponse
	15, // 21: IDM.GetUserPassword:output_type -> PasswordResponse
	13, // 22: IDM.UpdateUserPassword:output_type -> EnhancedInfoResponse
	16, // 23: IDM.AddUserToBlacklist:output_type -> InfoResponse
	16, // 24: IDM.CheckUserInBlacklist:output_type -> InfoResponse
	13, // 25: IDM.RemoveUserFromBlacklist:output_type -> EnhancedInfoResponse
	13, // [13:26] is the sub-list for method output_type
	0,  // [0:13] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_idm_proto_init() }
func file_idm_proto_init() {
	if File_idm_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_idm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idm_proto_goTypes,
		DependencyIndexes: file_idm_proto_depIdxs,
	}.Build()
	File_idm_proto = out.File
	file_idm_proto_rawDesc = nil
	file_idm_proto_goTypes = nil
	file_idm_proto_depIdxs = nil
}
