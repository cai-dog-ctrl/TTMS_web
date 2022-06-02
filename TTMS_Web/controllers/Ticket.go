package controllers

import (
	"TTMS/models"
	"TTMS/pkg/utils"
	"TTMS/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//有关票务的controller代码

// 根据演出计划ID查看票
func GetTicketByScheduleId(c *gin.Context) {
	p := c.Param("id")
	if p == "" {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("GetTicketByScheduleId getid Error")
		return
	}

	p1, err := service.GetTicketByScheduleId(utils.ShiftToNum64(p))

	if err != nil {
		zap.L().Error("service.GetTicketByScheduleId error", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, "获取失败")
		return
	}

	ResponseSuccess(c, p1)

}

// 根据movie_id和时间date_day查询票
func GetTicketByMovieIdAndDateDay(c *gin.Context) {

}

// 根据影厅cinema_id和date_day查询票
func GetTicketByCinemaIdAndDateDay(c *gin.Context) {

}

// 买票
func SaleTicket(c *gin.Context) {
	p := new(models.ParamsSaleTicket)
	err := c.ShouldBind(&p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("SaleTicket ShouldBind Error", zap.Error(err))
	}
	is, err1, price := service.SaleTicket(p)
	if err1 != nil || !is {
		ResponseError(c, CodeServerBusy)
		zap.L().Error("service.SaleTicket Error", zap.Error(err))
		return
	}
	ResponseSuccess(c, gin.H{
		"TotalPrice": price,
	})
}

// 退票
