package services

import (
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent"
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent/user"
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/repositories"
	"MSaaS-Framework/MSaaS/internal/security"
	"MSaaS-Framework/MSaaS/pkg/object"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type UserService struct {
	repository *repositories.UserRepository
	client     *ent.Client
}

func NewUserService(repository *repositories.UserRepository, client *ent.Client) *UserService {
	return &UserService{
		repository: repository,
		client:     client,
	}
}

// CreateUser는 새로운 유저를 생성합니다.
func (s *UserService) CreateUser(ctx context.Context, userDTO object.User) (*ent.User, error) {
	// check if user already exists
	if _, err := s.repository.GetByEmail(ctx, userDTO.Email); err == nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, salt, err := security.HashPassword(userDTO.Password)
	if err != nil {
		return nil, err
	}

	user := &ent.User{
		Email:    userDTO.Email,
		Password: hashedPassword,
		Salt:     salt,
		Role:     user.Role(userDTO.Role),
		Name:     userDTO.Username,
	}

	fmt.Println(user)
	return s.repository.Create(ctx, user)
}

// GetUserByID는 전달받은 UUID에 해당하는 유저 정보를 조회합니다.
func (s *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	return s.repository.GetByID(ctx, id)
}

// UpdateUser는 전달받은 UUID에 해당하는 유저의 정보를 업데이트합니다.
func (s *UserService) UpdateUser(ctx context.Context, id uuid.UUID, userDTO object.User) (*ent.User, error) {
	existingUser, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	user := &ent.User{
		ID: id,
	}

	if existingUser.Password != userDTO.Password {
		hashedPassword, salt, err := security.HashPassword(userDTO.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hashedPassword
		user.Salt = salt
	}

	// TODO: Update user fields
	return s.repository.Update(ctx, user)
}

// DeleteUser는 전달받은 UUID에 해당하는 유저를 삭제합니다.
func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return s.repository.Delete(ctx, id)
}
