package main

import (
	"github.com/iammukeshm/structured-logging-golang-zap/utils"
	"go.uber.org/zap"
)

func main() {
	utils.InitializeLogger()
	utils.Logger.Info("Hello World")
	utils.Logger.Error("Not able to reach blog.",zap.String("url", "codewithmukesh.com"))
}