package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/SwanPoi/bmstu_rsoi_lab1/internal/models"
	"github.com/SwanPoi/bmstu_rsoi_lab1/internal/services"
	"github.com/SwanPoi/bmstu_rsoi_lab1/pkg/logger"
)

type PersonsHandler struct {
	service *services.PersonService
	Logger *logger.Logger
}

func NewPersonsHandler(service *services.PersonService, logger *logger.Logger) *PersonsHandler {
	return &PersonsHandler{
		service: service,
		Logger: logger,
	}
}

func (h *PersonsHandler) GetPersons(ctx *gin.Context) {
	persons, err := h.service.GetAll()

	if err != nil {
		h.Logger.Errorf(err.Error())
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{ Message: err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, persons)
}

func (h *PersonsHandler) AddPerson(ctx *gin.Context) {
	var req models.PersonUpsert

	if err := ctx.ShouldBindJSON(&req); err != nil {
		var validationErrors validator.ValidationErrors

		if ok := errors.As(err, &validationErrors); ok {
			resp := models.ValidationErrorResponse{
				Message: "Validation Error",
				Errors: make(map[string]string),
			}

			for _, fieldErr := range validationErrors {
				field := fieldErr.Field()
				tag := fieldErr.Tag()

				var message string

				switch tag {
				case "required":
					message = field + " is required"
				case "min":
					if field == "Age" {
						message = "Age must be non-negative"
					} else {
						message = field + " is too short"
					}
				default:
					message = "Invalid value of " + field	
				}

				h.Logger.Errorf(message)
				jsonName := fieldErr.StructField()
				resp.Errors[jsonName] = message
			}

			ctx.JSON(http.StatusBadRequest, resp )
			return
		}

		ctx.JSON(http.StatusBadRequest, models.ValidationErrorResponse{ Message: err.Error() })
		return
	}

	id, err := h.service.AddPerson(&req)

	if err != nil {
		h.Logger.Errorf(err.Error())

		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{ Message: err.Error() })
		return
	}

	location := "/api/v1/persons/" + strconv.FormatInt(int64(id), 10)
	ctx.Header("Location", location)
	ctx.Status(http.StatusCreated)
}