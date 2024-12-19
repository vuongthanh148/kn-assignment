package authsvc

import (
	"context"

	"kn-assignment/internal/constant"
	"kn-assignment/internal/core/domain"
	errors "kn-assignment/internal/core/error"
	"kn-assignment/internal/util"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) RegisterUser(ctx context.Context, user domain.CreateUserRequest) error {
	existingUser, err := s.repo.GetUserByUsername(ctx, user.Username)

	if err != nil {
		return err
	}

	if existingUser != nil {
		return errors.NewCustomError(constant.ErrCodeDuplicateUser)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.repo.CreateUser(ctx, user)
}

func (s *service) AuthenticateUser(ctx context.Context, username, password string) (domain.LoginResponse, error) {
	user, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return domain.LoginResponse{}, err
	}
	if user == nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return domain.LoginResponse{}, errors.NewCustomError(constant.ErrCodeInvalidCredential)
	}

	accessToken, err := util.GenerateAccessToken(user.ID, user.Username, string(user.Role))
	if err != nil {
		return domain.LoginResponse{}, errors.NewCustomError(constant.ErrCodeGenerateToken)
	}

	refreshToken, err := util.GenerateRefreshToken(user.ID, user.Username, string(user.Role))
	if err != nil {
		return domain.LoginResponse{}, errors.NewCustomError(constant.ErrCodeGenerateToken)
	}

	return domain.LoginResponse{
		User: domain.User{
			ID:        user.ID,
			Username:  user.Username,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		AccessToken:  accessToken,
		RefreshToken: &refreshToken,
	}, nil
}
