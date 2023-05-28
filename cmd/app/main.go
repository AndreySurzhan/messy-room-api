package main

import (
	"gitlab.stripchat.dev/myclub/go-service/internal/config"
	"gitlab.stripchat.dev/myclub/go-service/internal/pkg/app"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	a, err := app.New(cfg)
	if err != nil {
		panic(err)
	}

	err = a.Run()
	if err != nil {
		panic(err)
	}
}
