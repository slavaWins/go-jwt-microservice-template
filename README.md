<div align="center">

<h1> JWT LIB MICRO SERVICE</h1>
 
</div>
 
## About 

Базовая биба с проверкой авторизии jwt, для микросервиса. Нужная минималка есть чтоб бы не выносить себе мозги и сразу делать бизнесовое 

    go get -u github.com/slavaWins/go-jwt-microservice-template@master
    go get -u github.com/slavaWins/go-jwt-microservice-template



Controller:

    func GetUser(c *gin.Context) {

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}

	user := gjmt_models.User{}
	db := db_service.Connect()

	if db.First(&user, userIDUint).Error != nil {
		c.JSON(http.StatusNotFound, gjmt_models.NewErrorResponse("Error code id"))
		return
	}

	c.JSON(http.StatusOK, gjmt_models.NewSuccessResponse(user))
    }
    

Main

	gjmt_db_service.Connect()
	gjmt_db_service.Migrate()

	db_service.Migrate()

	r := gin.Default()

	protected := r.Group("/")
	protected.Use(gjmt_midlwares.AuthMiddleware())
	{
		protected.GET("/user", controllers.GetUser)
	}

	r.Run(":8081")
 