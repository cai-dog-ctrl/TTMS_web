package routes

import (
	"TTMS/controllers"
	"TTMS/logger"
	"TTMS/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Cors())
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	g := r.Group("/api")
	g.POST("/login", controllers.Login)
	g.POST("/register", controllers.Register)
	g.POST("/addadmin", controllers.AddAdmin)
	g.GET("/getallmsg", controllers.GetAllMsg)
	g.GET("/getusermsgbyid/:id", controllers.GetUserMsgById)
	g.GET("/getpicturebyfilename/:img", controllers.GetPictureByFileName)
	g.PUT("/updatemsg", controllers.UpdateMsg)
	g.Use()
	{

	}
	//cat front page / movie info / showing movies / coming movies / scoreRanking movies / boxofficeRanking movies
	g.GET("/GetFrontPage", controllers.GetFrontPage)
	g.GET("/GetMovieInfoById/:Id", controllers.GetMovieInfoByID)
	g.GET("/GetAllShowingMovies", controllers.GetShowingMovies)
	g.GET("/GetAllComingMovies", controllers.GetComingMovies)
	g.GET("/GetAllScoreRankingMovies", controllers.GetScoreRankingMovies)
	g.GET("/GetAllBoxOfficeRankingMovies", controllers.GetBoxOfficeRankingMovies)

	//manage movie
	g.POST("/AddNewMovie", controllers.AddNewMovie)
	g.PUT("/ModifyMovie", controllers.ModifyMovieByID)
	
	//manage cinema
	g.POST("/AddNewCinema", controllers.AddNewCinema)
	g.PUT("/ModifyCinemaByID", controllers.ModifyCinemaByID)
	
	//manage seat
	g.GET("/GetSeatByCinemaID", controllers.GetSeatByCinemaID)
	g.PUT("/ModifySeat", controllers.ModifySeat)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
