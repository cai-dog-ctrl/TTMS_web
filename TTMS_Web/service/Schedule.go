package service

import (
	"TTMS/dao/mysql"
	"TTMS/models"
	"TTMS/pkg/snowflake"
	"TTMS/pkg/utils"
	"errors"
	"strings"
)

//有关演出计划的业务逻辑代码

//AddSchedule 添加演出计划
func AddSchedule(p *models.ParamAddSchedule) (bool ,error) {

	err := mysql.RefreshSchedule()
	if err != nil {
		return false, errors.New("RefreshSchedule in serve.AddSchedule failed")
	}

	Sche := new(models.ScheduleIn)
	Sche.ID = snowflake.GenID()
	Sche.CinemaId = utils.ShiftToNum64(p.CinemaId)
	Sche.MovieId = utils.ShiftToNum64(p.MovieId)

	str_day  := strings.Split(p.DateDay, "-")
	date_day := str_day[0] + str_day[1] + str_day[2] 

	str_start_time 	:= strings.Split(p.StartTime, ":")
	start_time 		:= str_start_time[0] + str_start_time[1]

	Sche.DateDay = utils.ShiftToNum64(date_day)
	Sche.StartTime = utils.ShiftToNum64(start_time)
	Sche.Price = p.Price

  	ret, err := mysql.InsertSchedule(Sche)
	if err != nil {
		return false, err
	}

	if err == nil && !ret {
		return false, nil
	}

	return true, nil
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

func GetAllScheduleMsgByMovieId(page_num, page_size int, movie_id int64) (*models.ScheduleRetList, error) {
	
	err := mysql.RefreshSchedule()
	if err != nil {
		return nil, errors.New("RefreshSchedule in serve.GetAllScheduleMsgByMovieId failed")
	}

	p1, err := mysql.GetAllScheduleByMovieId(page_num,page_size,movie_id)
	if err != nil {
		return nil, err
	}

	ret := new(models.ScheduleRetList)
	ret.Total = p1.Total
	for _, it := range p1.List {
		date_day := ""
		date_day += utils.ShiftToStringFromInt64(int64(it.DateDay / 10000))
		date_day += "-"
		day_time := it.DateDay % 10000
		date_day += utils.ShiftToStringFromInt64(int64(day_time / 100))
		date_day += "-"
		date_day += utils.ShiftToStringFromInt64(int64(day_time % 100))

		start_time := ""
		start_time += utils.ShiftToStringFromInt64(int64(it.StartTime / 100))
		start_time += ":"
		start_time += utils.ShiftToStringFromInt64(int64(it.StartTime % 100))

		end_time := ""
		end_time += utils.ShiftToStringFromInt64(int64(it.EndTime / 100))
		end_time += ":"
		end_time += utils.ShiftToStringFromInt64(int64(it.EndTime % 100))

		ret.List = append(ret.List, models.ScheduleRet{
			ID: utils.ShiftToStringFromInt64(it.ID),
			CinemaId: utils.ShiftToStringFromInt64(it.CinemaId),
			MovieId: utils.ShiftToStringFromInt64(it.MovieId),
			MovieName: it.MovieName,
			DateDay: date_day,
			StartTime: start_time,
			EndTime: end_time,
			Price: it.Price,
			CinemaName: it.CinemaName,
			Type: it.Type,
		})
	}

	return ret, nil
}

func GetAllScheduleMsgByCinemaId(page_num, page_size int, cinema_id int64) (*models.ScheduleRetList, error) {
	
	err := mysql.RefreshSchedule()
	if err != nil {
		return nil, errors.New("RefreshSchedule in serve.GetAllScheduleMsgByCinemaId failed")
	}
	
	p1, err := mysql.GetAllScheduleByCinemaId(page_num,page_size,cinema_id)
	
	if err != nil {
		return nil, err
	}

	ret := new(models.ScheduleRetList)
	ret.Total = p1.Total
	for _, it := range p1.List {
		date_day := ""
		date_day += utils.ShiftToStringFromInt64(int64(it.DateDay / 10000))
		date_day += "-"
		day_time := it.DateDay % 10000
		date_day += utils.ShiftToStringFromInt64(int64(day_time / 100))
		date_day += "-"
		date_day += utils.ShiftToStringFromInt64(int64(day_time % 100))

		start_time := ""
		start_time += utils.ShiftToStringFromInt64(int64(it.StartTime / 100))
		start_time += ":"
		start_time += utils.ShiftToStringFromInt64(int64(it.StartTime % 100))

		end_time := ""
		end_time += utils.ShiftToStringFromInt64(int64(it.EndTime / 100))
		end_time += ":"
		end_time += utils.ShiftToStringFromInt64(int64(it.EndTime % 100))

		ret.List = append(ret.List, models.ScheduleRet{
			ID: utils.ShiftToStringFromInt64(it.ID),
			CinemaId: utils.ShiftToStringFromInt64(it.CinemaId),
			MovieId: utils.ShiftToStringFromInt64(it.MovieId),
			MovieName: it.MovieName,
			DateDay: date_day,
			StartTime: start_time,
			EndTime: end_time,
			Price: it.Price,
			CinemaName: it.CinemaName,
			Type: it.Type,
		})
	}

	return ret, nil
}

func GetAllScheduleByMovieIdandDay(movie_id, day int64) ( *models.ScheduleRetList, error) {
	
	err := mysql.RefreshSchedule()
	if err != nil {
		return nil, errors.New("RefreshSchedule in serve.GetAllScheduleByMovieIdandDay failed")
	}
	
	p1, err := mysql.GetAllScheduleByMovieIdandDay(movie_id, day)

	if err != nil {
		return nil, err
	}

	ret := new(models.ScheduleRetList)
	ret.Total = p1.Total
	for _, it := range p1.List {
		date_day := ""
		date_day += utils.ShiftToStringFromInt64(int64(it.DateDay / 10000))
		date_day += "-"
		day_time := it.DateDay % 10000
		date_day += utils.ShiftToStringFromInt64(int64(day_time / 100))
		date_day += "-"
		date_day += utils.ShiftToStringFromInt64(int64(day_time % 100))

		start_time := ""
		start_time += utils.ShiftToStringFromInt64(int64(it.StartTime / 100))
		start_time += ":"
		start_time += utils.ShiftToStringFromInt64(int64(it.StartTime % 100))

		end_time := ""
		end_time += utils.ShiftToStringFromInt64(int64(it.EndTime / 100))
		end_time += ":"
		end_time += utils.ShiftToStringFromInt64(int64(it.EndTime % 100))

		ret.List = append(ret.List, models.ScheduleRet{
			ID: utils.ShiftToStringFromInt64(it.ID),
			CinemaId: utils.ShiftToStringFromInt64(it.CinemaId),
			MovieId: utils.ShiftToStringFromInt64(it.MovieId),
			MovieName: it.MovieName,
			DateDay: date_day,
			StartTime: start_time,
			EndTime: end_time,
			Price: it.Price,
			CinemaName: it.CinemaName,
			Type: it.Type,
		})
	}

	return ret, nil
}


func UpdateSchedule(p *models.ParamsUpdateScheduleMsg) (bool, error) {
	
	err := mysql.RefreshSchedule()
	if err != nil {
		return false, errors.New("RefreshSchedule in serve.UpdateSchedule failed")
	}
	
	Sch := new(models.SCheduledata)
	Sch.ID 			= utils.ShiftToNum64(p.ID)
	Sch.CinemaId 	= utils.ShiftToNum64(p.CinemaId)
	Sch.MovieId 	= utils.ShiftToNum64(p.MovieId)

	str_day  := strings.Split(p.DateDay, "-")
	date_day := str_day[0] + str_day[1] + str_day[2] 

	str_start_time 	:= strings.Split(p.StartTime, ":")
	start_time 		:= str_start_time[0] + str_start_time[1]

	Sch.DateDay		= utils.ShiftToNum64(date_day)
	Sch.StartTime 	= utils.ShiftToNum64(start_time)
	Sch.Price 		= p.Price

	ret, err := mysql.UpdateSchedule(Sch)
	if err != nil {
		return false, err
	}
	if err == nil && !ret{
		return false, nil
	}
	return true, nil
}

func DeleteSchedule(id int64) error {

	err := mysql.RefreshSchedule()
	if err != nil {
		return errors.New("RefreshSchedule in serve.DeleteSchedule failed")
	}

	err = mysql.DeleteSchedule(id)
	if err != nil {
		return err
	}
	return nil
}

func GetAllScheduleDayByMovieID(movie_id int64) (*models.ScheRets, error) {
	
	err := mysql.RefreshSchedule()
	if err != nil {
		return nil, errors.New("RefreshSchedule in serve.GetAllScheduleDayByMovieID failed")
	}


	p1, err := mysql.GetAllScheduleDayByMovieID(movie_id)
	if err != nil {
		return nil, err
	}

	p := new(models.ScheRets)
	for _, it := range p1.Time {
		day := ""
		day += utils.ShiftToStringFromInt64(int64(it.Time / 10000))
		day += "-"
		time := it.Time % 10000
		day += utils.ShiftToStringFromInt64(int64(time / 100))
		day += "-"
		day += utils.ShiftToStringFromInt64(int64(time % 100))
		p.Time = append(p.Time, models.ScheDayRet{
			Time: day,
		})
	}

	return p, nil
}	

func GetAllScheduleMsgByDay(day string) (*models.ScheduleRetList, error) {

	err := mysql.RefreshSchedule()
	if err != nil {
		return nil, errors.New("RefreshSchedule in serve.GetAllScheduleMsgByDay failed")
	}

	day_string := strings.Trim(day, "-")

	day_int64 := utils.ShiftToNum64(day_string)

	p1, err := mysql.GetAllScheduleMsgByDay(day_int64)

	p := new(models.ScheduleRetList)
	p.Total = p1.Total
	for _, it := range p1.List {
		date_day := ""
		date_day += utils.ShiftToStringFromInt64(int64(it.DateDay / 10000))
		date_day += "-"
		day_time := it.DateDay % 10000
		date_day += utils.ShiftToStringFromInt64(int64(day_time / 100))
		date_day += "-"
		date_day += utils.ShiftToStringFromInt64(int64(day_time % 100))

		start_time := ""
		start_time += utils.ShiftToStringFromInt64(int64(it.StartTime / 100))
		start_time += ":"
		start_time += utils.ShiftToStringFromInt64(int64(it.StartTime % 100))

		end_time := ""
		end_time += utils.ShiftToStringFromInt64(int64(it.EndTime / 100))
		end_time += ":"
		end_time += utils.ShiftToStringFromInt64(int64(it.EndTime % 100))

		p.List = append(p.List, models.ScheduleRet{
			ID: utils.ShiftToStringFromInt64(it.ID),
			CinemaId: utils.ShiftToStringFromInt64(it.CinemaId),
			MovieId: utils.ShiftToStringFromInt64(it.MovieId),
			MovieName: it.MovieName,
			DateDay: date_day,
			StartTime: start_time,
			EndTime: end_time,
			Price: it.Price,
			CinemaName: it.CinemaName,
			Type: it.Type,
		})
	}

	if err != nil {
		return nil, err
	}
	return p, nil

}