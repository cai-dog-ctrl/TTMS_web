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
	p := c.Query("id")
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
	movie_id := c.Query("movie_id")
	date_day := c.Query("date_day")

	if movie_id == "" || date_day == "" {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("GetTicketByMovieIdAndDateDay get movie_id or date_day  Error")
		return
	}

	p1, err := service.GetTicketByMovieIdAndDateDay(movie_id, date_day)

	if err != nil {
		zap.L().Error("service.GetTicketByMovieIdAndDateDay error", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, "获取失败")
		return
	}

	ResponseSuccess(c, p1)
}

// 根据影厅cinema_id和date_day查询票
func GetTicketByCinemaIdAndDateDay(c *gin.Context) {
	cinema_id 	:= c.Query("cinema_id")
	date_day 	:= c.Query("date_day")

	if cinema_id == "" || date_day == "" {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("GetTicketByCinemaIdAndDateDay get cinema_id or date_day  Error")
		return
	}

	p1, err := service.GetTicketByCinemaIdAndDateDay(cinema_id, date_day)

	if err != nil {
		zap.L().Error("service.GetTicketByCinemaIdAndDateDay error", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, "获取失败")
		return
	}

	ResponseSuccess(c, p1)
}

// 买票
func SaleTicket(c *gin.Context) {
	p := new(models.ParamsSaleTicket)
	err := c.ShouldBind(&p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("SaleTicket ShouldBind Error", zap.Error(err))
	}
	is, price, err1 := service.SaleTicket(p)
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
func Refund(c *gin.Context) {
	ticket_id 	:= c.Query("ticket_id")
	user_id 	:= c.Query("user_id")

	if ticket_id == "" || user_id == "" {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("Refund get ticket_id or user_id  Error")
		return
	}

	err := service.Refund(utils.ShiftToNum64(ticket_id), utils.ShiftToNum64(user_id))
	if err != nil {
		ResponseError(c, CodeServerBusy)
		zap.L().Error("service.Refund Error", zap.Error(err))
		return
	}
	
	ResponseSuccess(c, "Refund ticket successful.")
}