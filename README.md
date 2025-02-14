<div align="center">

<h1> JWT LIB MICRO SERVICE</h1>
 
</div>
 
## About 

Базовая биба с проверкой авторизии jwt, для микросервиса. Нужная минималка есть чтоб бы не выносить себе мозги и сразу делать бизнесовое 

    go get -u github.com/slavaWins/go-jwt-microservice-template@master
    go get -u github.com/slavaWins/go-jwt-microservice-template



Controller:

    func GetUser(c fiber.Ctx) error {
	user, err := gjmt_midlwares.GetAuthUser(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(gjmt_models.ResponseWithError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(gjmt_models.ResponseWithValue(user))

    }
    

Main

    fmt.Println("============ Start ============ ")
	fmt.Println(os.Getenv("APP_NAME"))
	fmt.Println(os.Getenv("APP_PORT"))

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		fmt.Errorf("Не настроен env! %w", err)
		return
	}

	gjmt_db_service.Connect()
	gjmt_db_service.Migrate()

	db_service.Migrate()

	r := fiber.New()

	routes.ApiRoutes(r)

	// Маршрут для Swagger-документации
	if os.Getenv("APP_PRODUCTION") == "false" {
		docs.SwaggerInfo.Title = os.Getenv("APP_NAME")

		docs.SwaggerInfo.BasePath = os.Getenv("SWAGGER_USE_CUSTOM_BASEPATH")

		r.Get("/swagger/*", swagger.HandlerDefault)
		//r.Get("/swagger/*any", fiberSwagger.WrapHandler(swaggerFiles.Handler))
	}

	log.Fatal(r.Listen(":" + os.Getenv("APP_PORT")))
