package wrappers

import "github.com/mihnea1711/POS_Project/services/gateway/idm/proto_files"

// --------------------------------------------------------- LoginResponse ---------------------------------------------------------
type LoginResponse struct {
	Response *proto_files.LoginResponse
}

func (ur *LoginResponse) IsResponseNil() bool {
	return ur.Response == nil
}

func (ur *LoginResponse) IsInfoNil() bool {
	return ur.Response != nil && ur.Response.Info == nil
}

func (ur *LoginResponse) IsTokenEmpty() bool {
	return ur.Response != nil && ur.Response.Token == ""
}

// --------------------------------------------------------- UsersResponse ---------------------------------------------------------
type UsersResponse struct {
	Response *proto_files.UsersResponse
}

func (ur *UsersResponse) IsResponseNil() bool {
	return ur.Response == nil
}

func (ur *UsersResponse) IsInfoNil() bool {
	return ur.Response != nil && ur.Response.Info == nil
}

func (ur *UsersResponse) IsUsersNil() bool {
	return ur.Response != nil && ur.Response.Users == nil
}

// --------------------------------------------------------- UsersResponse ---------------------------------------------------------
type UserResponse struct {
	Response *proto_files.UserResponse
}

func (ur *UserResponse) IsResponseNil() bool {
	return ur.Response == nil
}

func (ur *UserResponse) IsInfoNil() bool {
	return ur.Response != nil && ur.Response.Info == nil
}

func (ur *UserResponse) IsUserNil() bool {
	return ur.Response != nil && ur.Response.User == nil
}
