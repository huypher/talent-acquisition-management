package app

import (
	"net/http"

	authDelivery "github.com/pghuy/talent-acquistion-management/auth"
	"github.com/pghuy/talent-acquistion-management/domain"
	"github.com/pghuy/talent-acquistion-management/infra"
	"github.com/pghuy/talent-acquistion-management/level"
	"github.com/pghuy/talent-acquistion-management/talent"
	"github.com/uptrace/bun"
)

func ProvideConfig() (*infra.Config, error) {
	return infra.NewConfig()
}

func ProvidePostgres(cfg *infra.Config) (*bun.DB, func(), error) {
	return infra.NewPostgres(cfg)
}

func ProvideHttpHandler(
	authDelivery domain.AuthDelivery,
	talentDelivery domain.TalentDelivery,
) http.Handler {
	return infra.NewHttpHandler(authDelivery, talentDelivery)
}

func ProvideRestService(httpHandler http.Handler) *http.Server {
	return infra.NewRestService(httpHandler)
}

func ProvideAuthDelivery(authUsecase domain.AuthUsecase) domain.AuthDelivery {
	return authDelivery.NewAuthDelivery(authUsecase)
}

func ProvideAuthUsecase(userUsecase domain.TalentUsecase) domain.AuthUsecase {
	return authDelivery.NewAuthUsecase(userUsecase)
}

func ProvideTalentRepository(db *bun.DB) (domain.TalentRepository, error) {
	return talent.NewTalentRepository(db)
}

func ProvideTalentUsecase(userRepository domain.TalentRepository) domain.TalentUsecase {
	return talent.NewTalentUsecase(userRepository)
}

func ProvideTalentDelivery(talentUsecase domain.TalentUsecase) domain.TalentDelivery {
	return talent.NewTalentDelivery(talentUsecase)
}

func ProvideLevelRepository(db *bun.DB) (domain.LevelRepository, error) {
	return level.NewLevelRepository(db)
}

func ProvideLevelUsecase(levelRepository domain.LevelRepository) domain.LevelUsecase {
	return level.NewLevelUsecase(levelRepository)
}

func ProvideLevelDelivery(levelUsecase domain.LevelUsecase) domain.LevelDelivery {
	return level.NewLevelDelivery(levelUsecase)
}
