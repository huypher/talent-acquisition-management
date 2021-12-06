//+build wireinject

package app

import (
	"context"

	"github.com/google/wire"
)

func InitApplication(ctx context.Context) (*App, func(), error) {
	wire.Build(
		set,
		App{},
	)

	return nil, nil, nil
}
