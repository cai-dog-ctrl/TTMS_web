package controllers

import (
	"TTMS/models"
	"TTMS/service"
	"TTMS/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//有关演出假话的controller代码

// AddSchedule 添加演出计划
func AddSchedule(c *gin.Context) {
	p := new(models.ParamAddSchedule)
	err := c.ShouldBind(&p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("AddSchedule ShouldBind Error")
		return
	}
	err = service.AddSchedule(p)
	if err != nil {
		zap.L().Error("service.AddSchedule error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, "添加成功")
}


// // DeleteSchedule 删除演出计划
// func DeleteSchedule(c *gin.Context) {
// 	p := c.Query("id")

// 	err := service.DeleteSchedule(p)
// 	if err != nil {
// 		zap.L().Error("service.DeleteSchedule error", zap.Error(err))
// 		ResponseError(c, CodeServerBusy)
// 		return
// 	}
// 	ResponseSuccess(c, "删除成功")
// }

// UpdateSchedule 修改演出计划
func UpdateSchedule(c *gin.Context) {
	p := new(models.ParamsUpdateScheduleMsg)
	err := c.ShouldBind(&p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("UpdateSchedule ShouldBind Error", zap.Error(err))
		return
	}
	err = service.UpdateSchedule(p)
	if err != nil {
		zap.L().Error("service.UpdateSchedule error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, "修改成功")
}

// 通过电影ID获取所有演出计划信息
func GetAllScheduleMsgByMovieId(c *gin.Context) {
	pageNum := utils.ShiftToNum(c.Query("page_num"))
	pageSize := utils.ShiftToNum(c.Query("page_size"))
	movie_id := utils.ShiftToNum(c.Query("movie_id"))

	p1, err := service.GetAllScheduleMsgByMovieId(pageNum,pageSize,movie_id)

	if err != nil {
		zap.L().Error("service.GetAllScheduleMsgByMovieId error", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, "获取失败")

		return
	}
	ResponseSuccess(c, p1)
}

// 通过影厅ID获取所有演出计划信息
func GetAllScheduleMsgByCinemaId(c *gin.Context) {
	pageNum := utils.ShiftToNum(c.Query("page_num"))
	pageSize := utils.ShiftToNum(c.Query("page_size"))
	cinema_id := utils.ShiftToNum(c.Query("cinema_id"))

	p1, err := service.GetAllScheduleMsgByCinemaId(pageNum,pageSize,cinema_id)

	if err != nil {
		zap.L().Error("service.GetAllScheduleMsgByCinemaId error", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, "获取失败")

		return
	}
	ResponseSuccess(c, p1)
}

func GetAllScheduleByMovieIdandDay(c *gin.Context) {
	movie_id := utils.ShiftToNum64(c.Query("movie_id"))
	day := utils.ShiftToNum64(c.Query("date_day"))

	p1, err := service.GetAllScheduleByMovieIdandDay(movie_id, day)

	if err != nil {
		zap.L().Error("service.GetAllScheduleByMovieIdandDay error", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, "获取失败")

		return
	}
	ResponseSuccess(c, p1)

}