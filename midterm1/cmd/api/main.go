package main

import (
	"github.com/Tamir1205/midterm1/internal"
	"github.com/Tamir1205/midterm1/internal/config"
)

func main() {
	// init app
	cfg, err := config.NewConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	err = internal.NewApp(cfg).Run()
	if err != nil {
		panic(err)
	}
}
