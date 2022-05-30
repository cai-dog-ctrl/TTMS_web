package controllers

import (
	"TTMS/models"
	"TTMS/pkg/utils"
	"TTMS/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

//有关影院的controller代码
func GetCinemaByID(c *gin.Context){
	
}

func GetAllCinemas(c *gin.Context){
	p := new(models.ParamsGetAllCinemas)
	p.Num = utils.ShiftToNum(c.Query("Num"))
	p.Page_num = utils.ShiftToNum(c.Query("Page_num"))
	cinema, err := service.GetAllCinemas(p)
	if err != nil{
		zap.L().Error("", zap.Error(err))
		return
	}
	ResponseSuccess(c, cinema)
}

func AddNewCinema(c *gin.Context){
	p := new(models.ParamsAddNewCinema)
	err := c.ShouldBind(&p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("AddNewCinema ShouldBind Error", zap.Error(err))
	}
	err = service.AddNewCinema(p)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		zap.L().Error("service.AddNewCinema Error", zap.Error(err))
		return
	}
	ResponseSuccess(c, "add new cinema successful.")
}

func ModifyCinemaByID(c *gin.Context) {
	p := new(models.ParamsModifyCinema)
	err := c.ShouldBind(p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("ModifyCinemaByID ShouldBind Error", zap.Error(err))
		return
	}
	err = service.ModifyCinemaByID(p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("service.ModifyCinemaByID Error", zap.Error(err))
		return
	}
	ResponseSuccess(c, "modify cinema successful.")
}

func GetSeatByCinemaID(c *gin.Context){
	p := new(models.ParamsGetSeatByCinemaID)
	err := c.ShouldBind(&p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("GetSeatByCinemaID ShouldBind Error", zap.Error(err))
		return
	}
	seatlist, err1 := service.GetSeatByCinemaID(p)
	if err1 != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("service.GetSeatByCinemaID Error", zap.Error(err))
		return
	}
	ResponseSuccess(c, seatlist)
}

func ModifySeat(c *gin.Context) {
	p := new(models.ParamsModifySeat)
	err := c.ShouldBind(p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("ModifySeat ShouldBind Error", zap.Error(err))
		return
	}
	err = service.ModifySeat(p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("service.ModifySeat Error", zap.Error(err))
		return
	}
	ResponseSuccess(c, "modify seat successful.")
}