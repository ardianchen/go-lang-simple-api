package routers

import (
	apiController "github.com/ardianchen/simple-api/src/controllers"
	"github.com/ardianchen/simple-api/src/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		// add header Access-Control-Allow-Origin
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	//API route for version 1
	v1 := r.Group("/api/v1")

	//If you want to pass your route through specific middlewares
	v1.Use(middlewares.UserMiddlewares())
	{
		v1.GET("user-list", apiController.UserList)
		v1.POST("user-list", apiController.CreateUser)
	}

	return r

}
