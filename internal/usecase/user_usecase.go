package usecase

import (
	"context"

	"go-setup/internal/entity"
	"go-setup/internal/repository"
	"go-setup/pkg/errors"
)

type UserUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (uc *UserUsecase) CreateUser(ctx context.Context, user *entity.User) error {
	// Validate user
	if user.Name == "" {
		return errors.ErrUserNameRequired
	}
	if user.Email == "" {
		return errors.ErrUserEmailRequired
	}
	
	// Check if user exists
	_, err := uc.userRepo.GetByEmail(ctx, user.Email)
	if err == nil {
		return errors.ErrConflict
	}
	
	return uc.userRepo.Create(ctx, user)
}

func (uc *UserUsecase) GetUserByID(ctx context.Context, id int64) (*entity.User, error) {
	return uc.userRepo.GetByID(ctx, id)
}

func (uc *UserUsecase) ListUsers(ctx context.Context) ([]*entity.User, error) {
	return uc.userRepo.List(ctx)
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, user *entity.User) error {
	return uc.userRepo.Update(ctx, user)
}

func (uc *UserUsecase) DeleteUser(ctx context.Context, id int64) error {
	return uc.userRepo.Delete(ctx, id)
}