package main

import (
	"github.com/lenuse/mall"
	"github.com/lenuse/mall/config"
	"github.com/lenuse/mall/repository"
)

func main() {
	_ = repository.Open()
	defer repository.Close()
	app := mall.New()
	_ = app.Run(config.New().Server.Port)
}
