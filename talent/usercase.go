package talent

import (
	"context"

	"github.com/pghuy/talent-acquistion-management/domain"
)

type talentUsecase struct {
	talentRepository domain.TalentRepository
}

func NewTalentUsecase(userRepository domain.TalentRepository) *talentUsecase {
	return &talentUsecase{
		talentRepository: userRepository,
	}
}

func (u *talentUsecase) GetByUserName(ctx context.Context, userName string) (*domain.Talent, error) {
	return u.talentRepository.GetByUserName(ctx, userName)
}

func (u *talentUsecase) GetByID(ctx context.Context, id int) (*domain.Talent, error) {
	return u.talentRepository.GetByID(ctx, id)
}
