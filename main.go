package main

import (
	"fmt"

	"github.com/wmoura85/gopportunities/config"
	"github.com/wmoura85/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
    logger = config.GetLogger("main")
	//iniciar Configs
	err error := config.Init()
	if err =! nil {
		logger.Errorf("config initialization error: %v",err)
		return
	}


	//iIniciar Router
	router.Initialize()
}
	