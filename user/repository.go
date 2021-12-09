package user

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/huypher/talent-acquisition-management/domain"
)

type user struct {
	ID       int
	UserName string
	Password string
	Name     string
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) (*userRepository, error) {
	return &userRepository{
		db: db,
	}, nil
}

func (r *userRepository) GetByUserName(ctx context.Context, userName string) (domain.User, error) {
	model := user{}

	if err := r.db.WithContext(ctx).Model(&model).Where("username = ?", userName).Find(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, nil
		}
		return domain.User{}, err
	}

	return domain.User{
		ID:       model.ID,
		UserName: model.UserName,
		Password: model.Password,
		Name:     model.Name,
	}, nil
}

func (r *userRepository) GetByID(ctx context.Context, id int) (domain.User, error) {
	model := user{}

	if err := r.db.WithContext(ctx).Model(&model).Where("id = ?", id).Find(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, nil
		}
		return domain.User{}, err
	}

	return domain.User{
		ID:       model.ID,
		UserName: model.UserName,
		Password: model.Password,
		Name:     model.Name,
	}, nil
}
