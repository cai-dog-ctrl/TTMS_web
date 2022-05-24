package controllers

import (
	"TTMS/models"
	"TTMS/service"
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

//
//// DeleteSchedule 删除演出计划
//func DeleteSchedule() error {
//
//}
//
//// UpdateSchedule 修改演出计划
//func UpdateSchedule() error {
//
//}
