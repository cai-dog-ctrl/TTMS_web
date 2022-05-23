package controllers

import (
	"TTMS/models"
	"TTMS/pkg/utils"
	"TTMS/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//有关电影的controller代码

func GetFrontPage(c *gin.Context) {
	p := new(models.ParamsFrontPage)
	p.CarouselNum = utils.ShiftToNum(c.Query("CarouselNum"))
	p.ShowingNum = utils.ShiftToNum(c.Query("ShowingNum"))
	p.ComingNum = utils.ShiftToNum(c.Query("ComingNum"))
	p.ScoreRankingNum = utils.ShiftToNum(c.Query("ScoreRankingNum"))
	p.BoxOfficeRankingNum = utils.ShiftToNum(c.Query("BoxOfficeRankingNum"))

	front_page, err := service.GetFrontPage(p)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		return
	}
	ResponseSuccess(c, front_page)
}

func GetMovieInfoByID(c *gin.Context) {
	p := new(models.ParamsMovie)
	p.Id = utils.ShiftToNum64(c.Param("Id"))
	fmt.Println(p)
	movie, err := service.GetMovieInfoByID(p)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		return
	}
	relevantMovies, err := service.GetRelevantMovies(movie.Tag)
	ResponseSuccess(c, gin.H{
		"movieInfo":      movie,
		"relevantMovies": relevantMovies,
	})
}

func GetShowingMovies(c *gin.Context) {
	p := new(models.ParamsGetShowingMovies)
	p.Num = utils.ShiftToNum(c.Query("Num"))
	p.Page_num = utils.ShiftToNum(c.Query("Page_num"))
	movie, err := service.GetAllShowingMovies(p)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		return
	}
	ResponseSuccess(c, movie)
}

func GetComingMovies(c *gin.Context) {
	p := new(models.ParamsGetComingMovies)
	p.Num = utils.ShiftToNum(c.Query("Num"))
	p.Page_num = utils.ShiftToNum(c.Query("Page_num"))
	movie, err := service.GetAllComingMovies(p)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		return
	}
	ResponseSuccess(c, movie)
}

func GetScoreRankingMovies(c *gin.Context) {
	p := new(models.ParamsGetScoreRankingMovies)
	p.Num = utils.ShiftToNum(c.Query("Num"))
	p.Page_num = utils.ShiftToNum(c.Query("Page_num"))
	movie, err := service.GetAllScoreRankingMovies(p)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		return
	}
	ResponseSuccess(c, movie)
}

func GetBoxOfficeRankingMovies(c *gin.Context) {
	p := new(models.ParamsGetBoxOfficeRankingMovies)
	p.Num = utils.ShiftToNum(c.Query("Num"))
	p.Page_num = utils.ShiftToNum(c.Query("Page_num"))
	movie, err := service.GetAllBoxOfficeRankingMovies(p)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		return
	}
	ResponseSuccess(c, movie)
}

func AddNewMovie(c *gin.Context) {
	p := new(models.ParamsAddNewMovie)
	err := c.ShouldBind(&p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("AddNewMovie ShouldBind Error", zap.Error(err))
		return
	}
	err = service.AddNewMovie(p)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		zap.L().Error("service.AddNewMovie Error", zap.Error(err))
		return
	}
	ResponseSuccess(c, "add new movie successful.")
}

func ModifyMovieByID(c *gin.Context) {
	p := new(models.ParamsModifyMovie)
	err := c.ShouldBind(p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("ModifyMovieByID ShouldBind Error", zap.Error(err))
		return
	}
	err = service.ModifyMovieByID(p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("service.ModifyMovieByID Error", zap.Error(err))
		return
	}
	ResponseSuccess(c, "modify movie successful.")
}
