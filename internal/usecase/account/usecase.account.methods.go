package account

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (u *Usecase) Signup(ctx context.Context, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.InternalServer(err.Error(), "failed to hash password")
	}

	accountID, err := u.accountRepo.CreateAccount(ctx, email, string(hashedPassword))
	if err != nil {
		return errors.InternalServer(err.Error(), "failed to create account")
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		return errors.InternalServer(err.Error(), "failed to create verification token")
	}

	token := fmt.Sprintf("%d:%s", accountID, uuid.String())
	if err := u.accountRepo.SetVerificationToken(ctx, accountID, token); err != nil {
		return errors.InternalServer(err.Error(), "failed to set verification token")
	}

	encodedToken := base64.URLEncoding.EncodeToString([]byte(token))
	if err := u.accountRepo.SendVerificationToken(ctx, email, encodedToken); err != nil {
		return errors.InternalServer(err.Error(), "failed to send verification token")
	}

	return nil
}

func (u *Usecase) Verify(ctx context.Context, encodedToken string) error {
	token, err := base64.URLEncoding.DecodeString(encodedToken)
	if err != nil {
		return errors.InternalServer(err.Error(), "failed to decode token")
	}

	accountID, err := u.accountRepo.GetAccountIDFromToken(ctx, string(token))
	if err != nil {
		return errors.InternalServer(err.Error(), "failed to get account id")
	}

	if accountID == 0 {
		return errors.BadRequest("account id is not found", "Verification link is invalid")
	}

	if err := u.accountRepo.SetAccountToVerified(ctx, accountID); err != nil {
		return errors.InternalServer(err.Error(), "failed to set account to verified")
	}

	return nil
}

func (u *Usecase) Login(ctx context.Context, email, password string) (string, error) {
	accountRef, err := u.accountRepo.GetAccountRefFromEmail(ctx, email)
	if err != nil {
		return "", errors.InternalServer(err.Error(), "failed to retrieve account ref")
	}

	if accountRef == nil {
		return "", errors.Unauthorized("account ref is empty", "account doesn't exist")
	}

	bcrypt.CompareHashAndPassword([]byte(accountRef.HashedPassword), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return "", errors.Unauthorized("hashed password not matched", "email/password is invalid")
		}
		return "", errors.InternalServer(err.Error(), "failed to compare password")
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", errors.InternalServer(err.Error(), "failed to create verification token")
	}

	token := fmt.Sprintf("%d:%s", accountRef.AccountID, uuid.String())
	if err := u.accountRepo.SetSession(ctx, token, *accountRef); err != nil {
		return "", errors.InternalServer(err.Error(), "failed to set session")
	}

	encodedToken := base64.URLEncoding.EncodeToString([]byte(token))
	return encodedToken, nil
}
