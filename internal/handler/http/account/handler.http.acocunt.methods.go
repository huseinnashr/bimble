package account

import (
	"context"

	v1 "github.com/huseinnashr/bimble/api/v1"
)

// Signup implements v1.AccountServiceHTTPServer.
func (*Handler) Signup(context.Context, *v1.SignupRequest) (*v1.SignupResponse, error) {
	return &v1.SignupResponse{
		Message: "pong",
	}, nil
}
