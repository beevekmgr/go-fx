package message

func registerRoutes(handler *handler.Handler) {
	handler.Gin.GET("/message",func(c *gin.Context){
		c.JSON(200,gin.H("message": "Yo, whats up"))
	})
}

var Module = fx.Options(fx.Invoke(registerRoutes))
