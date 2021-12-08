package talent

import (
	"context"

	"github.com/huypher/kit/container"

	"github.com/huypher/talent-acquisition-management/domain"
)

type talentUsecase struct {
	talentRepository domain.TalentRepository
}

func NewTalentUsecase(userRepository domain.TalentRepository) *talentUsecase {
	return &talentUsecase{
		talentRepository: userRepository,
	}
}

func (u *talentUsecase) GetByID(ctx context.Context, id int) (*domain.Talent, error) {
	return u.talentRepository.GetByID(ctx, id)
}

func (u *talentUsecase) GetList(ctx context.Context, filter container.Map, pageID, perPage int) ([]domain.Talent, error) {
	return u.talentRepository.GetList(ctx, filter, pageID, perPage)
}
