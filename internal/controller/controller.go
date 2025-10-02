package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/SwanPoi/bmstu_rsoi_lab1/internal/handlers"
	"github.com/SwanPoi/bmstu_rsoi_lab1/internal/services"
	"github.com/SwanPoi/bmstu_rsoi_lab1/internal/repositories"
	"github.com/SwanPoi/bmstu_rsoi_lab1/pkg/logger"
)

func SetupRoutes(routes *gin.Engine, db *gorm.DB, logger *logger.Logger) {
	personRepo := repositories.NewPersonRepository(db)
	personService := services.NewPersonService(personRepo)
	personsHandler := handlers.NewPersonsHandler(personService, logger)

	api := routes.Group("/api/v1")
	{
		persons := api.Group("/persons")
		persons.GET("", personsHandler.GetPersons)
		persons.POST("", personsHandler.AddPerson)
	}
}