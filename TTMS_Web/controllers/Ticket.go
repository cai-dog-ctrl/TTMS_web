package controllers

import (
	"TTMS/models"
	"TTMS/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//有关票务的controller代码

// 根据演出计划查看票

// 根据movie_id和时间date_day查询票

// 根据影厅cinema_id和date_day查询票

// 买票
func SaleTicket(c *gin.Context){
	p := new(models.ParamsSaleTicket)
	err := c.ShouldBind(&p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("SaleTicket ShouldBind Error", zap.Error(err))
	}
	err = service.SaleTicket(p)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		zap.L().Error("service.SaleTicket Error", zap.Error(err))
		return
	}
	ResponseSuccess(c, "sale ticket successful.")
}

// 退票
