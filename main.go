package main

import "api-sqlite/config"

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Iniciando configuracao error: %v", err)
		return
	}

}
