<div align="center">

<h1> JWT LIB MICRO SERVICE</h1>
 
</div>
 
## About 

Базовая биба с проверкой авторизии jwt, для микросервиса. Нужная минималка есть чтоб бы не выносить себе мозги и сразу делать бизнесовое 

    go get -u github.com/slavaWins/go-jwt-microservice-template@master
    go get -u github.com/slavaWins/go-jwt-microservice-template



    
	gjmt_db_service.Connect()
	gjmt_db_service.Migrate()

## Example usage 



    fmt.Println("Start")
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		fmt.Errorf("Не настроен env! %w", err)
		return
	}

	db_service.Connect()
	usecases.Migrate()

	r := gin.Default()

	protected2 := r.Group("/")
	protected2.Use(midlwares.RateLimitMiddleware(4))
	{
		protected2.POST("/login", controllers.LoginController)
		protected2.POST("/code", controllers.CodeController)
	}

	protected := r.Group("/")
	protected.Use(midlwares.AuthMiddleware())
	{
		protected.GET("/user", controllers.GetUser)
	}

	r.Run(":8081")
