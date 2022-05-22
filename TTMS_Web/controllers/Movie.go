package controllers

import (
	"TTMS/models"
	"TTMS/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//有关电影的controller代码

func GetFrontPage(c *gin.Context) {
	p := new(models.ParamsFrontPage)
	err := c.ShouldBind(&p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("FrontPage ShouldBind Error")
		return
	}
	front_page, err := service.GetFrontPage(p)
	if err != nil {
		return
	}
	ResponseSuccess(c, front_page)
}

func GetMovieInfo(c *gin.Context) {
	p := new(models.ParamsMovie)
	err := c.ShouldBind(&p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("Movie ShouldBind Error")
		return
	}
	movie, err := service.GetMovieInfo(p)
	if err != nil {
		return
	}
	ResponseSuccess(c, movie)
}
