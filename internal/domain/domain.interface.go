package domain

import (
	"context"
)

type IAccountUsecase interface {
	Signup(ctx context.Context, email, password string) error
	Verify(ctx context.Context, encodedToken string) error
	Login(ctx context.Context, email, password string) (string, error)
}

type IAccountRepo interface {
	HashPassword(ctx context.Context, password string) (string, error)
	CreateAccount(ctx context.Context, email, hashedPassword string) (int64, error)
	SetVerificationToken(ctx context.Context, accountID int64, encodedToken string) error
	SendVerificationToken(ctx context.Context, email, encodedToken string) error
	GetAccountIDFromToken(ctx context.Context, token string) (int64, error)
	SetAccountToVerified(ctx context.Context, accountID int64) error
	GetAccountRefFromEmail(ctx context.Context, email string) (AccountRef, error)
	SetSession(ctx context.Context, token string, accountRef AccountRef) error
}

type AccountRef struct {
	AccountID          int64  `json:"account_id"`
	DatingProfileID    int64  `json:"dating_profile_id"`
	DatingPreferenceID int64  `json:"dating_preference_id"`
	HashedPassword     string `json:"-"`
}
