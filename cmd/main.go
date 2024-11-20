package main

import (
	"whatisyourtime/internal/configs"
	"whatisyourtime/internal/rest/server"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	config()
	router := server.NewEngine(gin.DebugMode)
	if err := router.Run(":8080"); err != nil {
		logrus.Panicf("failed to run router : %v", err)
	}
}

func config() {
	configs.Setup()
}
