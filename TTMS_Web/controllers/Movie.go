package controllers

import (
	"TTMS/models"
	"TTMS/pkg/utils"
	"TTMS/service"
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
	err := c.ShouldBind(&p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("Movie ShouldBind Error")
		return
	}
	movie, err := service.GetMovieInfoByID(p)
	if err != nil {
		return
	}
	ResponseSuccess(c, movie)
}
