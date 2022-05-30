package service

import (
	"TTMS/dao/mysql"
	"TTMS/models"
	"TTMS/pkg/snowflake"
	"TTMS/pkg/utils"
)

//有关演出计划的业务逻辑代码

//AddSchedule 添加演出计划
func AddSchedule(p *models.ParamAddSchedule) error {
	Sche := new(models.ScheduleIn)
	Sche.ID = snowflake.GenID()
	Sche.CinemaId = p.CinemaId
	Sche.MovieId = p.MovieId
	Sche.DateDay = p.DateDay
	Sche.StartTime = p.StartTime
	Sche.Price = p.Price

	Sche.EndTime = 10

	err := mysql.InsertSchedule(Sche)
	if err != nil {
		return err
	}
	return nil
}

// // DeleteSchedule 删除演出计划
// func DeleteSchedule(p string) error {

// 	id := utils.ShiftToNum64(p)

// 	err := mysql.DeleteSchedule(id)

// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func GetAllScheduleMsgByMovieId(page_num, page_size int, movie_id int) (p *models.ScheduleList, err error) {
	
	p1, err := mysql.GetAllScheduleByMovieId(page_num,page_size,movie_id)
	
	
	if err != nil {
		return nil, err
	}
	return p1, nil
}

func GetAllScheduleMsgByCinemaId(page_num, page_size int, cinema_id int) (p *models.ScheduleList, err error) {
	
	p1, err := mysql.GetAllScheduleByCinemaId(page_num,page_size,cinema_id)
	
	
	if err != nil {
		return nil, err
	}
	return p1, nil
}

func GetAllScheduleByMovieIdandDay(movie_id, day int64) (p *models.ScheduleList, err error) {
	p1, err := mysql.GetAllScheduleByMovieIdandDay(movie_id, day)

	if err != nil {
		return nil, err
	}
	return p1, nil
}


func UpdateSchedule(p *models.ParamsUpdateScheduleMsg) error {
	Sch := new(models.SCheduledata)
	Sch.ID 			= utils.ShiftToNum64(p.ID)
	Sch.CinemaId 	= utils.ShiftToNum64(p.CinemaId)
	Sch.MovieId 	= utils.ShiftToNum64(p.MovieId)
	Sch.DateDay		= utils.ShiftToNum64(p.DateDay)
	Sch.StartTime 	= utils.ShiftToNum64(p.StartTime)
	Sch.EndTime 	= utils.ShiftToNum64(p.EndTime)
	Sch.IsDelete	= p.IsDelete
	Sch.IsShow 		= p.IsShow

	err := mysql.UpdateSchedule(Sch)
	if err != nil {
		return err
	}
	return nil
}