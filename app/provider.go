package app

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/huypher/talent-acquisition-management/user"

	authDelivery "github.com/huypher/talent-acquisition-management/auth"
	"github.com/huypher/talent-acquisition-management/domain"
	"github.com/huypher/talent-acquisition-management/infra"
	"github.com/huypher/talent-acquisition-management/level"
	"github.com/huypher/talent-acquisition-management/talent"
)

func ProvideConfig() (*infra.Config, error) {
	return infra.NewConfig()
}

func ProvidePostgres(cfg *infra.Config) (*gorm.DB, func(), error) {
	return infra.NewPostgres(cfg)
}

func ProvideHttpHandler(
	authDelivery domain.AuthDelivery,
	talentDelivery domain.TalentDelivery,
	userDelivery domain.UserDelivery,
) http.Handler {
	return infra.NewHttpHandler(authDelivery, talentDelivery, userDelivery)
}

func ProvideRestService(httpHandler http.Handler) *http.Server {
	return infra.NewRestService(httpHandler)
}

func ProvideUserRepository(db *gorm.DB) (domain.UserRepository, error) {
	return user.NewUserRepository(db)
}

func ProvideUserUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	return user.NewUserUsecase(userRepository)
}

func ProvideUserDelivery(userUsecase domain.UserUsecase) domain.UserDelivery {
	return user.NewProductDelivery(userUsecase)
}

func ProvideAuthDelivery(authUsecase domain.AuthUsecase) domain.AuthDelivery {
	return authDelivery.NewAuthDelivery(authUsecase)
}

func ProvideAuthUsecase(userUsecase domain.UserUsecase) domain.AuthUsecase {
	return authDelivery.NewAuthUsecase(userUsecase)
}

func ProvideTalentRepository(db *gorm.DB) (domain.TalentRepository, error) {
	return talent.NewTalentRepository(db)
}

func ProvideTalentUsecase(userRepository domain.TalentRepository) domain.TalentUsecase {
	return talent.NewTalentUsecase(userRepository)
}

func ProvideTalentDelivery(talentUsecase domain.TalentUsecase) domain.TalentDelivery {
	return talent.NewTalentDelivery(talentUsecase)
}

func ProvideLevelRepository(db *gorm.DB) (domain.LevelRepository, error) {
	return level.NewLevelRepository(db)
}

func ProvideLevelUsecase(levelRepository domain.LevelRepository) domain.LevelUsecase {
	return level.NewLevelUsecase(levelRepository)
}

func ProvideLevelDelivery(levelUsecase domain.LevelUsecase) domain.LevelDelivery {
	return level.NewLevelDelivery(levelUsecase)
}
