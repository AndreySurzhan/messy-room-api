package main

import (
	"flag"
	"fmt"
	"github.com/AndreySurzhan/messy-room-api/internal/app"
	"github.com/AndreySurzhan/messy-room-api/internal/config"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "./config.yaml", "path to config file")
	flag.Parse()

	fmt.Printf("config file: %+v\n", configFile)

	cfg, err := config.New(configFile)
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
