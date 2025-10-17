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

	connString := os.Getenv("DATABASE_URL")

	if connString == "" {
		connString = db.GetConnectionString(&c.Database)
	}

	dbHandler := db.Init(connString, l)
	router := gin.Default()

	port := os.Getenv("PORT")

	if port == "" {
		port = fmt.Sprintf("%d", c.HTTP.Port)
	}

	serverUrl := fmt.Sprintf(":%s", port)

	controller.SetupRoutes(router, dbHandler, l)
	router.Run(serverUrl)
}