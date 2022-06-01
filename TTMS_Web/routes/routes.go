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
	g.GET("/GetMovieByName", controllers.GetMovieByName)
	g.GET("/GetMovieInfoByID/:Id", controllers.GetMovieInfoByID)
	g.GET("/GetAllMovies", controllers.GetAllMovies)
	g.GET("/GetAllShowingMovies", controllers.GetShowingMovies)
	g.GET("/GetAllComingMovies", controllers.GetComingMovies)
	g.GET("/GetAllScoreRankingMovies", controllers.GetScoreRankingMovies)
	g.GET("/GetAllBoxOfficeRankingMovies", controllers.GetBoxOfficeRankingMovies)
	g.GET("/GetMovieSort", controllers.GetMovieSort)

	//manage movie
	g.POST("/AddNewMovie", controllers.AddNewMovie)
	g.PUT("/ModifyMovie", controllers.ModifyMovieByID)

	//manage cinema
	g.GET("/GetCinemaByID/:ID", controllers.GetCinemaByID)
	g.GET("/GetAllCinemas", controllers.GetAllCinemas)
	g.POST("/AddNewCinema", controllers.AddNewCinema)
	g.PUT("/ModifyCinemaByID", controllers.ModifyCinemaByID)
	g.PUT("/DeleteCinemaByID/:ID", controllers.DeleteCinemaByID)

	//manage seat
	g.GET("/GetSeatByCinemaID/:ID", controllers.GetSeatByCinemaID)
	g.PUT("/ModifySeat", controllers.ModifySeat)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	//manage schedule
	g.POST("/addschedule", controllers.AddSchedule)
	g.PUT("/updateSchedule", controllers.UpdateSchedule)
	g.GET("/getallschedulemsgbymovieid", controllers.GetAllScheduleMsgByMovieId)
	g.GET("/getallschedulemsgbycinemaid", controllers.GetAllScheduleMsgByCinemaId)
	g.GET("/getallschedulebymovieidandday", controllers.GetAllScheduleByMovieIdandDay)
	g.PUT("/deleteschedule/:id", controllers.DeleteSchedule)
	g.GET("/getalldcheduledaybymovieid/:movie_id", controllers.GetAllScheduleDayByMovieId)
	g.GET("/getallschedulemsbyday/:date_day", controllers.GetAllScheduleMsgByDay)
	return r
}
