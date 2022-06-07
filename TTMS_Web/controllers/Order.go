package controllers

import (
	"TTMS/models"
	"TTMS/pkg/utils"
	"TTMS/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetOrderByUserID(c *gin.Context) {
	p := new(models.ParamsGetOrderByUserID)
	p.ID = c.Query("ID")
	p.Num = utils.ShiftToNum(c.Query("Num"))
	p.Page_num = utils.ShiftToNum(c.Query("Page_num"))
	order, err := service.GetOrderByUserID(p)
	if err != nil {
		zap.L().Error("service.GetOrderByUserID ERROR", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, order)
}

func GetOrderByID(c *gin.Context) {
	p := new(models.ParamsGetOrderByID)
	p.ID = c.Param("ID")
	order, err := service.GetOrderByID(p)
	if err != nil {
		zap.L().Error("service.GetOrderByUserID ERROR", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, order)
}

func PayMoneyByOrderID(c *gin.Context) {
	p := new(models.ParamsPayMoneyByOrderID)
	p.ID = c.Param("ID")
	err := service.PayMoneyByOrderID(p)
	if err != nil {
		zap.L().Error("service.PayMoneyByOrderID ERROR", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, "pay money successfully.")
}

func CountAllSales(c *gin.Context) {

}

func CountSalesByDay(c *gin.Context) {

}

func CountSalesByMonth(c *gin.Context) {

}

func CountSalesByYear(c *gin.Context) {

}
