package main

import (
	"time"

	"github.com/avtara/vip-management-system-api/config"
	"github.com/avtara/vip-management-system-api/controller"
	"github.com/avtara/vip-management-system-api/repository"
	"github.com/avtara/vip-management-system-api/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	cors "github.com/itsjamie/gin-cors"
	"gorm.io/gorm"
)

var (
	db                *gorm.DB                     = config.SetupDatabaseConnection()
	athleteRepository repository.AthleteRepository = repository.NewAthleteRepository(db)
	athleteService    service.AthleteService       = service.NewAthleteService(athleteRepository)
	athleteController controller.AthleteController = controller.NewAthleteController(athleteService)
)

func main() {
	defer config.CloseDatabaseConnection(db)

	server := gin.Default()
	server.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE, OPTIONS",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
		date, ok := fl.Field().Interface().(time.Time)
		if ok {
			today := time.Now()
			if today.After(date) {
				return false
			}
		}
		return true
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	commonRoutes := server.Group("api", gin.BasicAuth(gin.Accounts{
		"receptionist": "1234",
		"internal":     "1234",
	}))
	{
		commonRoutes.GET("/vips", athleteController.AllAthlete)
		commonRoutes.GET("/vips/:id", athleteController.DetailAthlete)
	}

	receptionistRoutes := server.Group("api", gin.BasicAuth(gin.Accounts{
		"receptionist": "1234",
	}))
	{
		receptionistRoutes.PATCH("/vips/:id/arrived", athleteController.UpdateArrivedStatus)
	}

	internalRoutes := server.Group("api", gin.BasicAuth(gin.Accounts{
		"internal": "1234",
	}))
	{
		internalRoutes.POST("/vips", athleteController.InsertAthlete)
	}

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.Run()
}
