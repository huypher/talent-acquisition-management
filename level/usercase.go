package level

import (
	"context"

	"github.com/huypher/talent-acquisition-management/domain"
)

type levelUsecase struct {
	levelRepository domain.LevelRepository
}

func NewLevelUsecase(levelRepository domain.LevelRepository) *levelUsecase {
	return &levelUsecase{
		levelRepository: levelRepository,
	}
}

func (u *levelUsecase) GetAll(context context.Context) ([]domain.Level, error) {
	return u.levelRepository.GetAll(context)
}
