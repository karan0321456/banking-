package main

import (
	"banking/app"
	"banking/logger"
)

func main(){
	logger.Info("Starting the Application")
	app.Start()
}