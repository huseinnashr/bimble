package account

import (
	"context"

	v1 "github.com/huseinnashr/bimble/api/v1"
	"github.com/huseinnashr/bimble/internal/pkg/tracer"
)

// Signup implements v1.AccountServiceHTTPServer.
func (h *Handler) Signup(ctx context.Context, req *v1.SignupRequest) (*v1.SignupResponse, error) {
	ctx, span := tracer.Start(ctx, "handler.Signup")
	defer span.End()

	if err := h.accountUsecase.Signup(ctx, req.GetEmail(), req.GetPassword()); err != nil {
		return nil, err
	}

	return &v1.SignupResponse{
		Message: "Signup Success! Check your email for account verification link",
	}, nil
}

// Verify implements v1.AccountServiceHTTPServer.
func (h *Handler) Verify(ctx context.Context, req *v1.VerifyRequest) (*v1.VerifyResponse, error) {
	if err := h.accountUsecase.Verify(ctx, req.GetToken()); err != nil {
		return nil, err
	}

	return &v1.VerifyResponse{
		Message: "Your account has been verified",
	}, nil
}

// Login implements v1.AccountServiceHTTPServer.
func (h *Handler) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	encodedToken, err := h.accountUsecase.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, err
	}

	return &v1.LoginResponse{
		Token: encodedToken,
	}, nil
}
