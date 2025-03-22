package repositories

import (
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent"
	"MSaaS-Framework/MSaaS/cmd/wizcraft/app/ent/user"
	"context"
	"errors"
	"github.com/google/uuid"
)

type UserRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) *UserRepository {
	return &UserRepository{
		client: client,
	}
}

// Create는 전달받은 유저 데이터를 기반으로 새로운 유저를 생성합니다.
func (r *UserRepository) Create(ctx context.Context, user *ent.User) (*ent.User, error) {
	u, err := r.client.User.Create().
		SetEmail(user.Email).
		SetPassword(user.Password).
		SetSalt(user.Salt).
		SetName(user.Name).
		SetRole(user.Role).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) createTx(ctx context.Context, tx *ent.Tx, user *ent.User) (*ent.User, error) {
	u, err := tx.User.Create().
		SetEmail(user.Email).
		SetPassword(user.Password).
		SetSalt(user.Salt).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// GetByID는 전달받은 ID에 해당하는 유저를 반환합니다.
func (r *UserRepository) GetByID(ctx context.Context, uuid uuid.UUID) (*ent.User, error) {
	u, err := r.client.User.
		Query().
		Where(user.ID(uuid)).
		Only(ctx)
	if err != nil {
		// TODO: 데이터가 없는 경우, 조회에 실패한 경우 에러 처리 방법 검토 필요
		// 1. status code 반환
		// 2. 지정된 error 반환 후, service 레이어에서 분류하여 처리
		if ent.IsNotFound(err) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return u, nil
}

// getbyemail은 전달받은 이메일에 해당하는 유저를 반환합니다.
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*ent.User, error) {
	u, err := r.client.User.
		Query().
		Where(user.Email(email)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return u, nil
}

// Update는 전달받은 ID에 해당하는 유저를 업데이트합니다.
func (r *UserRepository) Update(ctx context.Context, user *ent.User) (*ent.User, error) {
	updatedUser, err := r.client.User.
		UpdateOneID(user.ID).
		SetPassword(user.Password).
		SetSalt(user.Salt).
		SetName(user.Name).
		SetActive(user.Active).
		Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return updatedUser, nil
}

// Delete는 전달받은 ID에 해당하는 유저를 삭제합니다.
func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := r.client.User.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return errors.New("user not found")
		}
		return err
	}
	return nil
}
