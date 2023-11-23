package middleware

import (
	"context"
	"log"

	"errors"
	"fmt"

	"github.com/mihnea1711/POS_Project/services/idm/idm/proto_files"
	"github.com/mihnea1711/POS_Project/services/idm/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ValidateRequestInterceptor is a gRPC interceptor for request validation.
func ValidateRequestInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	// Perform validation based on the method name
	switch info.FullMethod {
	case utils.IDMServiceName + utils.RegisterMethodName:
		return validateRegisterRequest(ctx, req, info, handler)
	case utils.IDMServiceName + utils.LoginMethodName:
		return validateLoginRequest(ctx, req, info, handler)
	case utils.IDMServiceName + utils.UpdateUserMethodName:
		return validateUpdateUserRequest(ctx, req, info, handler)
	case utils.IDMServiceName + utils.UpdateUserRoleMethodName:
		return validateUpdateUserRoleRequest(ctx, req, info, handler)
	case utils.IDMServiceName + utils.UpdateUserPasswordMethodName:
		return validateUpdateUserPasswordRequest(ctx, req, info, handler)
	default:
		// For methods without specific validation
		log.Printf("[IDM_VALIDATION] No validation needed for method: %s", info.FullMethod)
		return handler(ctx, req)
	}
}

func validateRegisterRequest(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Your validation logic for RegisterRequest...
	registerReq, ok := req.(*proto_files.RegisterRequest)
	if !ok {
		log.Printf("[IDM_VALIDATION] Invalid request type for method %s: expected %T, got %T", info.FullMethod, &proto_files.RegisterRequest{}, req)
		return nil, status.Error(codes.InvalidArgument, "invalid request type")
	}

	// Add your validation logic for UserCredentials, Role, etc.
	if err := validateUsername(registerReq.UserCredentials.Username); err != nil {
		log.Printf("[IDM_VALIDATION] Invalid username in %s: %v", info.FullMethod, err)
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid username: %v", err))
	}

	if err := validatePassword(registerReq.UserCredentials.Password); err != nil {
		log.Printf("[IDM_VALIDATION] Invalid password in %s: %v", info.FullMethod, err)
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid password: %v", err))
	}

	if err := validateRole(registerReq.Role); err != nil {
		log.Printf("[IDM_VALIDATION] Invalid role in %s: %v", info.FullMethod, err)
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid role: %v", err))
	}

	log.Printf("[IDM_VALIDATION] Validation successful for method: %s", info.FullMethod)
	// Call the next middleware or the actual RPC method
	return handler(ctx, req)
}

// Add logs and comments to validateLoginRequest
func validateLoginRequest(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	loginReq, ok := req.(*proto_files.LoginRequest)
	if !ok {
		log.Printf("[IDM_VALIDATION] Invalid request type for method %s: expected %T, got %T", info.FullMethod, &proto_files.LoginRequest{}, req)
		return nil, status.Error(codes.InvalidArgument, "invalid request type")
	}

	// Add your validation logic for UserCredentials...
	if err := validateUsername(loginReq.UserCredentials.Username); err != nil {
		log.Printf("[IDM_VALIDATION] Invalid username in %s: %v", info.FullMethod, err)
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid username: %v", err))
	}

	if err := validatePassword(loginReq.UserCredentials.Password); err != nil {
		log.Printf("[IDM_VALIDATION] Invalid password in %s: %v", info.FullMethod, err)
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid password: %v", err))
	}

	log.Printf("[IDM_VALIDATION] Validation successful for method: %s", info.FullMethod)
	// Call the next middleware or the actual RPC method
	return handler(ctx, req)
}

// Add logs and comments to validateUpdateUserRequest
func validateUpdateUserRequest(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	updateUserReq, ok := req.(*proto_files.UpdateUserRequest)
	if !ok {
		log.Printf("[IDM_VALIDATION] Invalid request type for method %s: expected %T, got %T", info.FullMethod, &proto_files.UpdateUserRequest{}, req)
		return nil, status.Error(codes.InvalidArgument, "invalid request type")
	}

	// Add your validation logic for UserData...
	if err := validateUserID(updateUserReq.UserData.IDUser); err != nil {
		log.Printf("[IDM_VALIDATION] Invalid UserID in %s: %v", info.FullMethod, err)
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid UserID: %v", err))
	}

	if err := validateUsername(updateUserReq.UserData.Username); err != nil {
		log.Printf("[IDM_VALIDATION] Invalid username in %s: %v", info.FullMethod, err)
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid username: %v", err))
	}

	log.Printf("[IDM_VALIDATION] Validation successful for method: %s", info.FullMethod)
	// Call the next middleware or the actual RPC method
	return handler(ctx, req)
}

// Add logs and comments to validateUpdateUserRoleRequest
func validateUpdateUserRoleRequest(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	updateRoleReq, ok := req.(*proto_files.UpdateRoleRequest)
	if !ok {
		log.Printf("[IDM_VALIDATION] Invalid request type for method %s: expected %T, got %T", info.FullMethod, &proto_files.UpdateRoleRequest{}, req)
		return nil, status.Error(codes.InvalidArgument, "invalid request type")
	}

	// Add your validation logic for UserID and Role...
	if err := validateUserID(updateRoleReq.UserID); err != nil {
		log.Printf("[IDM_VALIDATION] Invalid UserID in %s: %v", info.FullMethod, err)
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid UserID: %v", err))
	}

	if err := validateRole(updateRoleReq.Role); err != nil {
		log.Printf("[IDM_VALIDATION] Invalid role in %s: %v", info.FullMethod, err)
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid role: %v", err))
	}

	log.Printf("[IDM_VALIDATION] Validation successful for method: %s", info.FullMethod)
	// Call the next middleware or the actual RPC method
	return handler(ctx, req)
}

// Add logs and comments to validateUpdateUserPasswordRequest
func validateUpdateUserPasswordRequest(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	updatePasswordReq, ok := req.(*proto_files.UpdatePasswordRequest)
	if !ok {
		log.Printf("[IDM_VALIDATION] Invalid request type for method %s: expected %T, got %T", info.FullMethod, &proto_files.UpdatePasswordRequest{}, req)
		return nil, status.Error(codes.InvalidArgument, "invalid request type")
	}

	// Add your validation logic for UserID and Password...
	if err := validateUserID(updatePasswordReq.UserID); err != nil {
		log.Printf("[IDM_VALIDATION] Invalid UserID in %s: %v", info.FullMethod, err)
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid UserID: %v", err))
	}

	if err := validatePassword(updatePasswordReq.Password); err != nil {
		log.Printf("[IDM_VALIDATION] Invalid password in %s: %v", info.FullMethod, err)
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid password: %v", err))
	}

	log.Printf("[IDM_VALIDATION] Validation successful for method: %s", info.FullMethod)
	// Call the next middleware or the actual RPC method
	return handler(ctx, req)
}

func validateUserID(userID *proto_files.UserID) error {
	if userID == nil {
		return errors.New("UserID cannot be nil")
	}

	if userID.ID <= 0 {
		return errors.New("UserID must be a positive integer")
	}

	// Add more UserID validation logic if needed

	return nil
}

func validateUsername(username string) error {
	if username == "" {
		return errors.New("username cannot be empty")
	}
	// Add more username validation logic if needed
	return nil
}

func validatePassword(password string) error {
	if password == "" {
		return errors.New("password cannot be empty")
	}
	// Add more password validation logic if needed
	return nil
}

func validateRole(role string) error {
	allowedRoles := []string{"admin", "patient", "doctor"}
	for _, r := range allowedRoles {
		if r == role {
			return nil
		}
	}
	return errors.New("invalid role. role must be one of: admin, patient, doctor")
}
