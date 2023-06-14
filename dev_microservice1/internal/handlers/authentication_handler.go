package handlers

import (
	"context"
	"log"
	"net/http"

	pb "github.com/Nazerkh09/gajap/dev_microservice1/internal/protobuf"
)

type AuthenticationHandler struct {
	authClient pb.AuthenticationClient
}

// .
func NewAuthenticationHandler(authClient pb.AuthenticationClient) *AuthenticationHandler {
	return &AuthenticationHandler{
		authClient: authClient,
	}
}

func (h *AuthenticationHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Parse request parameters and validate inputs
	// ...

	// Call the authentication service
	response, err := h.authClient.Login(context.Background(), &pb.LoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Printf("Failed to authenticate user: %v", err)
		http.Error(w, "Failed to authenticate user", http.StatusInternalServerError)
		return
	}

	// Handle the authentication response
	// ...
}

func (h *AuthenticationHandler) Register(w http.ResponseWriter, r *http.Request) {
	// Parse request parameters and validate inputs
	// ...

	// Call the authentication service
	response, err := h.authClient.Register(context.Background(), &pb.RegisterRequest{
		Username: username,
		Password: password,
		Email:    email,
	})
	if err != nil {
		log.Printf("Failed to register user: %v", err)
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	// Handle the registration response
	// ...
}

// Additional handler functions for other authentication-related operations
// ...
