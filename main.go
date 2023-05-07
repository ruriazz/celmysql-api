package main

import (
	"github.com/celmysql-api/common"
	"github.com/celmysql-api/controller"
	"github.com/celmysql-api/database"
	"github.com/celmysql-api/middleware"
	"github.com/celmysql-api/repository"
	"github.com/celmysql-api/routes"
	"github.com/celmysql-api/services"
	"github.com/celmysql-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	// "github.com/swaggo/swag/example/basic/docs"
	"github.com/celmysql-api/docs"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// err := sentry.Init(sentry.ClientOptions{
	// 	Dsn: "https://64f70a24f5ce416798908acd368f1d5f@o4503903109644288.ingest.sentry.io/4503903111544832",
	// 	// Set TracesSampleRate to 1.0 to capture 100%
	// 	// of transactions for performance monitoring.
	// 	// We recommend adjusting this value in production,
	// 	TracesSampleRate: 1.0,
	// })
	// if err != nil {
	// 	log.Fatalf("sentry.Init: %s", err)
	// }
	config, _ := utils.LoadConfig("./")
	docs.SwaggerInfo.Title = "Celestial - MySQL API"
	docs.SwaggerInfo.Description = "Celestial - MySQL  API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}

	db := database.GetConnection(config)
	validate := validator.New()

	bankRepository := repository.NewBankRepository()
	bankService := services.NewBankService(bankRepository, db, validate)
	bankController := controller.NewBankController(bankService)

	permissionPolicyUserRepository := repository.NewPermissionPolicyUserRepository()
	imageFileRepository := repository.NewImageFileRepository()
	imageFileService := services.NewImageFileService(imageFileRepository, db, validate)
	imageFileController := controller.NewImageFileController(imageFileService)

	rajaOngkirService := services.NewRajaOngkirService(db, validate)
	rajaOngkirController := controller.NewRajaOngkirController(rajaOngkirService)

	sendEmailService := services.NewSendEmailService(db, validate)
	sendEmailController := controller.NewSendEmailController(sendEmailService)

	authService := services.NewAuthService(permissionPolicyUserRepository, db, validate)
	authController := controller.NewAuthController(authService)

	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // maximum file 8 MiB
	router.Static("/public", "./public")
	makeRoutes(router)
	router.Use(gin.Logger())
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {

		if common.NotFoundErrors(c, recovered) {
			return
		}

		if common.ValidationErrors(c, recovered) {
			return
		}

		common.InternalServerError(c, recovered)

	}))

	router.SetTrustedProxies([]string{"192.168.1.2"})

	v1 := router.Group(docs.SwaggerInfo.BasePath)
	{
		v1.Use(middleware.JWT())

		bank := v1.Group("/bank")
		routes.BankRouter(bankController, bank)

		// imageFile := v1.Group("/image-file")
		// routes.ImageFileRouter(imageFileController, imageFile)
		// // imageFile.Static("/", "./*public")
		// rajaOngkir := v1.Group("/raja-ongkir")
		// routes.RajaOngkirRouter(rajaOngkirController, rajaOngkir)

		// sendemail := v1.Group("/send-email")
		// routes.SendEmailRouter(sendEmailController, sendemail)

		imageFile := v1.Group("/image-file")
		routes.ImageFileRouter(imageFileController, imageFile)
		// imageFile.Static("/", "./*public")
		rajaOngkir := v1.Group("/raja-ongkir")
		routes.RajaOngkirRouter(rajaOngkirController, rajaOngkir)

		sendemail := v1.Group("/send-email")
		routes.SendEmailRouter(sendEmailController, sendemail)
	}

	v2 := router.Group(docs.SwaggerInfo.BasePath)
	{
		auth := v2.Group("/auth")
		routes.AuthRouter(authController, auth)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// router.Use(CORSMiddleware())
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"PUT", "GET", "OPTIONS", "POST", "PATCH"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "https://github.com"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))
	// router.Static("/public", "./public")
	// router.StaticFS("/", gin.Dir("dist", false))
	errPort := router.Run(":3000")

	common.PanicIfError(errPort)

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func makeRoutes(r *gin.Engine) {
	cors := func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Content-Type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}
		c.Next()
	}
	r.Use(cors)
}
