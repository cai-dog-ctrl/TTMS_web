package service

import (
	"TTMS/dao/mysql"
	"TTMS/models"
	"TTMS/pkg/snowflake"
)

//有关演出计划的业务逻辑代码

//AddSchedule 添加演出计划
func AddSchedule(p *models.ParamAddSchedule) error {
	Sche := new(models.ScheduleOut)
	Sche.ID = snowflake.GenID()
	Sche.CinemaId = p.CinemaId
	Sche.MovieId = p.MovieId
	Sche.DateDay = p.DateDay
	Sche.StartTime = p.StartTime

	Sche.EndTime = 10

	err := mysql.InsertSchedule(Sche)
	if err != nil {
		return err
	}
	return nil
}

// // DeleteSchedule 删除演出计划
// func DeleteSchedule(p int) error {

// }
