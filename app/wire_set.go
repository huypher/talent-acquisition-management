package app

import "github.com/google/wire"

var set = wire.NewSet(
	ProvideConfig,
	ProvidePostgres,
	ProvideHttpHandler,
	ProvideRestService,

	ProvideAuthDelivery,
	ProvideAuthUsecase,

	ProvideTalentRepository,
	ProvideTalentUsecase,
	ProvideTalentDelivery,

	ProvideLevelRepository,
	ProvideLevelUsecase,
	ProvideLevelDelivery,
)
