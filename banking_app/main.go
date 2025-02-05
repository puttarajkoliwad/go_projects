package main

import (
	"github.com/puttarajkoliwad/go_projects/banking_app/app"
	"github.com/puttarajkoliwad/go_projects/banking_app/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}