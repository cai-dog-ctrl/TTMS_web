package controllers

import (
	"TTMS/models"
	"TTMS/pkg/utils"
	"TTMS/service"
	// "fmt"

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
		zap.L().Error("service.GetOrderByID ERROR", zap.Error(err))
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

	p, err := service.CountAllSales()
	if err != nil {
		zap.L().Error("service.CountAllSales ERROR", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, p)
}

func CountSalesByDay(c *gin.Context) {

	day := c.Param("day")

	p, err := service.CountSalesByDay(day)
	if err != nil {
		zap.L().Error("service.CountSalesByDay ERROR", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, p)
}

func CountSalesByMonth(c *gin.Context) {
	month := c.Param("month")

	p, err := service.CountSalesByMonth(month)
	if err != nil {
		zap.L().Error("service.CountSalesByMonth ERROR", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, p)
}

func CountSalesByYear(c *gin.Context) {
	year := c.Param("year")

	p, err := service.CountSalesByYear(year)
	if err != nil {
		zap.L().Error("service.CountSalesByYear ERROR", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, p)
}


func RefundOrder(c *gin.Context) {
	Id := c.Param("id")

	ret, err := service.RefundOrder(Id)
	if err != nil || !ret{
		zap.L().Error("service.CountSalesByYear ERROR", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, "取消订单失败")
		return
	}
	ResponseSuccess(c, "取消订单")
}