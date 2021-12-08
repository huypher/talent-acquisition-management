package user

import (
	"context"

	"github.com/pghuy/talent-acquisition-management/domain"
)

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) *userUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) GetByUserName(ctx context.Context, userName string) (domain.User, error) {
	return u.userRepository.GetByUserName(ctx, userName)
}

func (u *userUsecase) GetByID(ctx context.Context, id int) (domain.User, error) {
	return u.userRepository.GetByID(ctx, id)
}
