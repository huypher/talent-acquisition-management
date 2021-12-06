package main

import (
	"context"

	"github.com/pghuy/dobi-oms/app"
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
