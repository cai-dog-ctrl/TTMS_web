package routes

import (
	"TTMS/controllers"
	"TTMS/logger"
	"TTMS/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Cors())
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	g := r.Group("/api")
	g.POST("/login", controllers.Login)
	g.POST("/register", controllers.Register)
	g.Use()
	{

	}
	g.GET("/frontPage", controllers.GetFrontPage)
	g.GET("/movieInfo", controllers.GetMovieInfoByID)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
