package internal

import (
	"context"

	pb "github.com/Nazerkh09/gajap/dev_microservice1/proto"
)

type authenticationService struct {
}

// NewAuthenticationService creates a new instance of the authentication service.
func NewAuthenticationService() pb.AuthenticationServiceServer {
	return &authenticationService{}
}

func (s *authenticationService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// Implement the login logic here
	// Validate the username and password, generate a token, etc.

	token := "example-token"
	response := &pb.LoginResponse{
		Token: token,
	}

	return response, nil
}

func (s *authenticationService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// Implement the registration logic here
	// Create a new user with the provided username, password, and email

	message := "Registration successful"
	response := &pb.RegisterResponse{
		Message: message,
	}

	return response, nil
}
