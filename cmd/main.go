package main

import (
	"context"

	"github.com/pghuy/talent-acquistion-management/app"
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
