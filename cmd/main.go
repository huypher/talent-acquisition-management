package main

import (
	"context"

	"github.com/huypher/talent-acquisition-management/app"
)

func main() {
	application, cleanup, err := app.InitApplication(context.Background())
	if err != nil {
		panic(err)
	}

	defer func() {
		cleanup()
	}()

	application.Run()
}
