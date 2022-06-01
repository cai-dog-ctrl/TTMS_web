package mysql

import (
	"TTMS/models"
	"errors"

	"TTMS/pkg/snowflake"

)

//有关演出计划的持久化代码

// InsertSchedule 添加演出计划，并生成票
func InsertSchedule(p *models.ScheduleIn) (bool, error) {

	sqlStr1 := "select movie_time from movie_info where id = ?"
	T := new(models.Sche)
	T.Id = 0
	err := db.Get(&T.Time, sqlStr1, p.MovieId)
	if err != nil {
		return false, nil
	}
	T.Time += 10  //预留10分钟退场和进场
	hour := T.Time / 60
	minute := T.Time % 60

	p.EndTime = p.StartTime + int64(hour) * 100 + int64(minute)

	if p.EndTime > int64(2400) {
		// err =  errors.New("the show time has expired, please reschedule")
		return false, nil
	}

	// sqlStr1 = "select id from showschedule where cinema_id = ? and movie_id = ? and ((start_time > ? and start_time < ?) or (start_time <= ? and end_time >= ?) or (start_time < ? and end_time > ?))"
	// err = db.Get(&T.Id, sqlStr1, p.CinemaId, p.MovieId, p.StartTime, p.EndTime, p.StartTime, p.EndTime, p.StartTime, p.StartTime)
	// if err == nil &&  T.Id != 0{
	// 	return errors.New("the show time has expired, please reschedule")
	// }

	sqlStr := "select id, cinema_id, movie_id, date_day, start_time, end_time, s.price, from showschdule where is_delete = -1 and cinema_id = ? and date_day = ?"
	Sches := new(models.ScheduleList)
	err = db.Select(&Sches.List, sqlStr, p.CinemaId, p.DateDay)
	if err != nil {
		return false, err
	}

	for _, it := range Sches.List{
		if p.StartTime <= it.StartTime && p.EndTime > it.StartTime { //已有演出计划的开始时间在新演出计划的S-E之间
			return false, nil
		}else if p.StartTime > it.StartTime && p.StartTime < it.EndTime { //新演出计划的开始时间在已有演出计划的S-E之间
			return false, nil
		}
	}

	tx, err := db.Beginx()  //开启事务
	if err != nil {
		return false, err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}else if err != nil {
			tx.Rollback()
		}else{
			err = tx.Commit()
		}
	}()

	sqlStr = "insert into showschdule (id, cinema_id, movie_id, date_day, start_time, end_time, price) values (?, ?, ?, ?, ?, ?, ?)"
	rs, err := tx.Exec(sqlStr, p.ID, p.CinemaId, p.MovieId, p.DateDay, p.StartTime, p.EndTime, p.Price)
	if err != nil {
		return false, err
	}

	n, err := rs.RowsAffected()
	if err != nil {
		return false, err
	}

	if n!=1 {
		return false, errors.New("exec sqlStr failed, insert showschedule failed")
	}

	Seats := new(models.Seats)
	sqlStr2 := "select id, cinema_id, roww, coll, status from seat_info where status = 1 and cinema_id = ?"
	err = db.Select(&Seats.List, sqlStr2, p.CinemaId)
	if err != nil {
		return false, err
	}

	for _, it := range Seats.List {
		sqlStr3 := "insert into ticket (id, schedule_id, cinema_id, movie_id, seat_id) values (?, ?, ?, ?, ?)"

		id := snowflake.GenID()

		rs, err := tx.Exec(sqlStr3, id, p.ID, p.CinemaId, p.MovieId, it.ID)
		if err != nil {
			return false, err
		}
	
		n, err := rs.RowsAffected()
		if err != nil {
			return false, err
		}
	
		if n!=1 {
			return false, errors.New("exec sqlStr failed, insert ticket failed")
		}
	}

	return true, err
}

// // DeleteSchedule 删除演出计划
// func DeleteSchedule(p int64) error {
// 	sqlStr := "update showschdule set is_delete=? where id = ?"
// 	_, err := db.Exec(sqlStr, 1, p)
// 	return err
// }

func GetAllScheduleByMovieId(page_num, page_size int, movie_id int64) (*models.ScheduleList, error) {
	Sches := new(models.ScheduleList)
	sqlStr := "select s.id, s.cinema_id, s.movie_id, s.date_day, s.start_time, s.end_time, s.is_show, s.price, c.cinema_name, c.tag  from showschdule s,cinema_info c where s.movie_id = ? and c.is_delete = -1 and s.cinema_id = c.id limit ?,?"
	err := db.Select(&Sches.List, sqlStr, movie_id, (page_num-1)*page_size, page_size)
	
	if err != nil {
		return nil, err
	}
	
	sqlStr1 := "select count(id) from showschdule where is_delete = -1 and movie_id = ?"

	err = db.Get(&Sches.Total, sqlStr1, movie_id)

	if err != nil {
		return nil, err
	}
	return Sches, nil
}

func GetAllScheduleByCinemaId(page_num, page_size int, cinema_id int64) (*models.ScheduleList, error) {
	Sches := new(models.ScheduleList)
	sqlStr := "select s.id, s.cinema_id, s.movie_id, s.date_day, s.start_time, s.end_time, s.is_show, s.price, c.cinema_name, c.tag  from showschdule s,cinema_info c where s.cinema_id = c.id and s.is_delete = -1 and s.cinema_id = ? limit ?,?"
	err := db.Select(&Sches.List, sqlStr, cinema_id, (page_num-1)*page_size, page_size)
	
	if err != nil {
		return nil, err
	}
	
	sqlStr1 := "select count(id) from showschdule where is_delete = -1 and cinema_id = ?"

	err = db.Get(&Sches.Total, sqlStr1, cinema_id)

	if err != nil {
		return nil, err
	}
	return Sches, nil

}

func GetAllScheduleByMovieIdandDay(movie_id, day int64) (*models.ScheduleList, error) {
	Sches := new(models.ScheduleList)
	sqlStr := "select s.id, s.cinema_id, s.movie_id, s.date_day, s.start_time, s.end_time, s.is_show, s.price, c.cinema_name, c.tag  from showschdule s,cinema_info c where s.movie_id = ? and s.date_day = ? and s.is_delete = -1 and s.cinema_id = c.id"
	err := db.Select(&Sches.List, sqlStr, movie_id, day)
	
	if err != nil {
		return nil, err
	}
	
	sqlStr1 := "select count(id) from showschdule where is_delete = -1 and movie_id = ? and date_day = day"

	err = db.Get(&Sches.Total, sqlStr1, movie_id)

	if err != nil {
		return nil, err
	}
	return Sches, nil
}



func UpdateSchedule(p *models.SCheduledata) ( bool, error) {

	sqlStr1 := "select movie_time from movie_info where id = ?"
	T := new(models.Sche)
	T.Id = 0
	err := db.Get(&T.Time, sqlStr1, p.MovieId)
	if err != nil {
		return false, nil
	}
	T.Time += 10  //预留10分钟退场和进场
	hour := T.Time / 60
	minute := T.Time % 60

	end_time := p.StartTime + int64(hour) * 100 + int64(minute)

	if end_time > int64(2400) {
		// err =  errors.New("the show time has expired, please reschedule")
		return false, nil
	}

	sqlStr := "select id, cinema_id, movie_id, date_day, start_time, end_time, s.price, from showschdule where is_delete = -1 and cinema_id = ? and date_day = ?"
	Sches := new(models.ScheduleList)
	err = db.Select(&Sches.List, sqlStr, p.CinemaId, p.DateDay)
	if err != nil {
		return false, err
	}

	for _, it := range Sches.List{
		if p.StartTime <= it.StartTime && end_time > it.StartTime { //已有演出计划的开始时间在新演出计划的S-E之间
			return false, nil
		}else if p.StartTime > it.StartTime && p.StartTime < it.EndTime { //新演出计划的开始时间在已有演出计划的S-E之间
			return false, nil
		}
	}

	sqlStr2 := "update showschdule set date_day=?, start_time=?, end_time=?, price=? where id = ?"
	_, err = db.Exec(sqlStr2, p.DateDay, p.StartTime, end_time, p.Price, p.ID)
	if err != nil {
		return false, err
	}
	return true,nil
}

func DeleteSchedule(id int64) error {
	sqlStr := "update showschdule set is_delete = 1 where id = ?"
	_,err := db.Exec(sqlStr, id)
	if err != nil {
		return err
	}

	sqlStr1 := "update ticket set is_delete = 1 where schedule_id = ?"
	_,err = db.Exec(sqlStr1, id)
	if err != nil {
		return err
	}

	return nil
}

func GetAllScheduleDayByMovieID(movie_id int64) (*models.ScheDay, error) {
	sqlStr := "select date_day, id from showschdule where movie_id = ? and is_delete = -1"
	Sches := new(models.ScheDay)
	err := db.Select(&Sches.Time, sqlStr, movie_id)

	if err != nil {
		return nil, err
	}
	return Sches, nil
}

func GetAllScheduleMsgByDay(day int64) (*models.ScheduleList, error) {
	
	sqlStr := "select s.id, s.cinema_id, s.movie_id, s.date_day, s.start_time, s.end_time, c.cinema_name, c.tag, m.name  from showschdule s,cinema_info c,movie_info m where s.is_delete = -1 and s.day = ?"
	Sches := new(models.ScheduleList)
	err := db.Select(&Sches.List, sqlStr, day)

	if err != nil {
		return nil, err
	}

	sqlStr1 := "select count(id) from showschdule where is_delete = -1 and date_day = ?"

	err = db.Get(&Sches.Total, sqlStr1, day)
	if err != nil {
		return nil, err
	}
	return Sches, nil
}