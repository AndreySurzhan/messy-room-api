package main

import (
	"github.com/AndreySurzhan/messy-room-api/internal/app"
	"github.com/AndreySurzhan/messy-room-api/internal/config"
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
