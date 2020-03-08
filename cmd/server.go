package main

import (
	"github.com/lenuse/mall"
	"github.com/lenuse/mall/config"
	"github.com/lenuse/mall/repository"
)

func main() {
	repository.Init()
	defer repository.Close()
	app := mall.New()
	app.Run(config.New().Server.Port)
}

