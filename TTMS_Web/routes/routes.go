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
	g.GET("/getallmsg", controllers.GetAllMsg)
	g.GET("/getpicturebyfilename/:img", controllers.GetPictureByFileName)


	//movie
	g.GET("/GetFrontPage", controllers.GetFrontPage)
	g.GET("/GetMovieByName", controllers.GetMovieByName)
	g.GET("/GetMovieInfoByID/:Id", controllers.GetMovieInfoByID)
	g.GET("/GetAllMovies", controllers.GetAllMovies)
	g.GET("/GetAllShowingMovies", controllers.GetShowingMovies)
	g.GET("/GetAllComingMovies", controllers.GetComingMovies)
	g.GET("/GetAllScoreRankingMovies", controllers.GetScoreRankingMovies)
	g.GET("/GetAllBoxOfficeRankingMovies", controllers.GetBoxOfficeRankingMovies)
	g.GET("/GetMovieSort", controllers.GetMovieSort)
	
	//schedule
	g.GET("/getallschedulemsbyday", controllers.GetAllScheduleMsgByDay)
	g.GET("/getallschedulemsgbymovieid", controllers.GetAllScheduleMsgByMovieId)
	g.GET("/getallschedulemsgbycinemaid", controllers.GetAllScheduleMsgByCinemaId)
	g.GET("/getallschedulebymovieidandday", controllers.GetAllScheduleByMovieIdandDay)


	// order
	g.GET("/CountAllSales", controllers.CountAllSales)
	g.GET("/CountSalesByDay/:day", controllers.CountSalesByDay)
	g.GET("/CountSalesByMonth/:month", controllers.CountSalesByMonth)
	g.GET("/CountSalesByYear/:year", controllers.CountSalesByYear)
	
	g.Use(middleware.JWTAuthMiddleware())
	{
		g.GET("/getusermsgbyid/:id", controllers.GetUserMsgById)
		g.POST("/addadmin", controllers.AddAdmin)
		g.PUT("/updatemsg", controllers.UpdateMsg)
	}
	//cat front page / movie info / showing movies / coming movies / scoreRanking movies / boxofficeRanking movies


	//manage movie
	g.POST("/AddNewMovie", controllers.AddNewMovie)
	g.PUT("/ModifyMovie", controllers.ModifyMovieByID)
	g.POST("/UploadPicture", controllers.UploadPicture)

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
	g.PUT("/deleteschedule/:id", controllers.DeleteSchedule)
	g.GET("/getallscheduledaybymovieid", controllers.GetAllScheduleDayByMovieId)
	g.GET("/GetScheduleMsgById/:id", controllers.GetScheduleMsgById)

	//sale
	g.PUT("/SaleTicket", controllers.SaleTicket)
	g.GET("/GetTicketByScheduleId", controllers.GetTicketByScheduleId)
	g.PUT("/Refund", controllers.Refund)
	// g.GET("/GetTicketByMovieIdAndDateDay", controllers.GetTicketByMovieIdAndDateDay)
	// g.GET("/GetTicketByCinemaIdAndDateDay", controllers.GetTicketByCinemaIdAndDateDay)

	
	//cat and manage order
	g.GET("/GetOrderByID/:ID", controllers.GetOrderByID)
	g.GET("/GetOrderByUserID", controllers.GetOrderByUserID)
	g.GET("/PayMoneyByOrderID/:ID", controllers.PayMoneyByOrderID)
	
	g.GET("/RefundOrder/:id", controllers.RefundOrder)
	return r
}
