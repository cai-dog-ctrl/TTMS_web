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
	g.GET("/GetFrontPage", controllers.GetFrontPage)
	g.GET("/GetMovieInfoById/:Id", controllers.GetMovieInfoByID)
	g.GET("/GetAllShowingMovies", controllers.GetShowingMovies)
	g.GET("/GetAllComingMovies", controllers.GetComingMovies)
	g.GET("/GetAllScoreRankingMovies", controllers.GetScoreRankingMovies)
	g.GET("/GetAllBoxOfficeRankingMovies", controllers.GetBoxOfficeRankingMovies)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
