package routes

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/forkyid/go-utils/v1/middleware"
	"github.com/gin-gonic/gin"
	"go-rest-api/docs"
	"go-rest-api/src/connection"
	"gorm.io/gorm"

	userController "go-rest-api/src/controller/v1"
	userModels "go-rest-api/src/models/v1"
	userService "go-rest-api/src/service/v1"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var master *gorm.DB
var router = gin.Default()

type DB struct {
	Master *gorm.DB
}

func Run() {	
	godotenv.Load()
	RouterSetup()
	router.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}

func RouterSetup() *gin.Engine {
	// set up
	router.SetTrustedProxies(nil)
	middleware := middleware.Middleware{}
	router.Use(middleware.CORS)

	// swagger
	docs.SwaggerInfo.Title = "Go Rest API"
	docs.SwaggerInfo.Description = "Go Rest API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	docs.SwaggerInfo.Schemes = []string{"https", "http"}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// database connection (type *gorm.DB)
	master = connection.DBMaster()

	// controller
	userModels := userModels.NewRepository(connection.DB{
		Master: master,
	})
	userService := userService.NewService(userModels)
	userController := userController.NewController(userService)

	// endpoint
	v1 := router.Group("/v1")
	users := v1.Group("/users")
	users.GET("register", userController.Register)

	return router
}
