package mysql

import "TTMS/models"

//有关演出计划的持久化代码

// InsertSchedule 添加演出计划
func InsertSchedule(p *models.ScheduleIn) error {
	sqlStr := "insert into showschdule (id, cinema_id, movie_id, date_day, start_time, end_time, price) values (?, ?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(sqlStr, p.ID, p.CinemaId, p.MovieId, p.DateDay, p.StartTime, p.EndTime, p.Price)
	return err
}

// // DeleteSchedule 删除演出计划
// func DeleteSchedule(p int64) error {
// 	sqlStr := "update showschdule set is_delete=? where id = ?"
// 	_, err := db.Exec(sqlStr, 1, p)
// 	return err
// }

func GetAllScheduleByMovieId(page_num, page_size int, movie_id int) (*models.ScheduleList, error) {
	Sches := new(models.ScheduleList)
	sqlStr := "select s.id, s.cinema_id, s.movie_id, s.date_day, s.start_time, s.end_time, s.is_show, s.price, c.cinema_name, c.tag  from showschdule s,cinema_info c where s.movie_id = ? and c.is_delete = -1 and s.cinema_id = c.id limit ?,?"
	err := db.Select(&Sches.List, sqlStr, movie_id, (page_num-1)*page_size, page_size)
	
	if err != nil {
		return nil, err
	}
	
	sqlStr1 := "select count(id) from user where is_delete = -1 and movie_id = ?"

	err = db.Get(&Sches.Total, sqlStr1, movie_id)

	if err != nil {
		return nil, err
	}
	return Sches, nil

}

func GetAllScheduleByCinemaId(page_num, page_size int, cinema_id int) (*models.ScheduleList, error) {
	Sches := new(models.ScheduleList)
	sqlStr := "select s.id, s.cinema_id, s.movie_id, s.date_day, s.start_time, s.end_time, s.is_show, s.price, c.cinema_name, c.tag  from showschdule s,cinema_info c where s.cinema_id = c.id and s.is_delete = -1 and s.cinema_id = ? limit ?,?"
	err := db.Select(&Sches.List, sqlStr, cinema_id, (page_num-1)*page_size, page_size)
	
	if err != nil {
		return nil, err
	}
	
	sqlStr1 := "select count(id) from user where is_delete = -1 and cinema_id = ?"

	err = db.Get(&Sches.Total, sqlStr1, cinema_id)

	if err != nil {
		return nil, err
	}
	return Sches, nil

}

func GetAllScheduleByMovieIdandDay(movie_id, day int64) (*models.ScheduleList, error) {
	Sches := new(models.ScheduleList)
	sqlStr := "select s.id, s.cinema_id, s.movie_id, s.date_day, s.start_time, s.end_time, s.is_show, s.price, c.cinema_name, c.tag  from showschdule s,cinema_info c where s.movie_id = ? and s.date_day = ? and s.is_delete = -1"
	err := db.Select(&Sches.List, sqlStr, movie_id, day)
	
	if err != nil {
		return nil, err
	}
	
	sqlStr1 := "select count(id) from user where is_delete = -1 and movie_id = ? and date_day = day"

	err = db.Get(&Sches.Total, sqlStr1, movie_id)

	if err != nil {
		return nil, err
	}
	return Sches, nil
}


func UpdateSchedule(p *models.SCheduledata) error {
	sqlStr := "update showschdule set date_day=?, start_time=?, end_time=?, is_delete=?, is_show=? where id = ?"
	_, err := db.Exec(sqlStr, p.DateDay, p.StartTime, p.EndTime, p.IsDelete, p.IsShow, p.ID)
	return err
}