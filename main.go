package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/SwanPoi/bmstu_rsoi_lab1/config"
	"github.com/SwanPoi/bmstu_rsoi_lab1/internal/controller"
	"github.com/SwanPoi/bmstu_rsoi_lab1/internal/db"
	"github.com/SwanPoi/bmstu_rsoi_lab1/pkg/logger"
)

func main() {
	configPath := "./config/envs"
	c, err := config.LoadConfig(configPath)

	if err != nil {
		log.Fatal(err)
	}

	loggerFile, err := os.OpenFile(
		c.Logger.File,
		os.O_APPEND | os.O_CREATE | os.O_WRONLY,
		0664,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer func(loggerFile *os.File) {
		err := loggerFile.Close()

		if err != nil {
			log.Fatal(err)
		}
	}(loggerFile)

	l := logger.New(c.Logger.Level, loggerFile)

	connString := db.GetConnectionString(&c.Database)
	fmt.Println("server url: ", connString)
	dbHandler := db.Init(connString, l)
	router := gin.Default()

	serverUrl := fmt.Sprintf(":%d",
		c.HTTP.Port,
	)

	fmt.Println("server url: ", serverUrl)
	controller.SetupRoutes(router, dbHandler, l)
	router.Run(serverUrl)
}