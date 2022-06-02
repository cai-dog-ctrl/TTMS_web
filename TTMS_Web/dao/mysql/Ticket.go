package mysql

import "TTMS/models"

//有关票务的持久化代码


func GetTicketByScheduleId(id int64) (*models.Ticks, error) {
	sqlStr := "select id, schedule_id, cinema_id, movie_id, seat_id, status from ticket where is_delete = -1"
	p := new(models.Ticks)
	err := db.Select(&p.List, sqlStr, id)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func GetTicketByMovieIdAndDateDay(movie_id, date_day int64) ( *models.Ticks, error) {
	p := new(models.Sche)
	sqlStr := "select id from showschdule where movie_id = ? and date_day = ? and is_delete = -1"
	err := db.Select(&p.Id, sqlStr, movie_id, date_day)
	if err != nil {
		return nil, err
	}

	sqlStr1 := "select id, schedule_id, cinema_id, movie_id, seat_id, status from ticket where schedule_id = ? and is_delete = -1 "
	p1 := new(models.Ticks)
	err = db.Select(&p1.List, sqlStr1, p.Id)
	if err != nil {
		return nil, err
	}
	return p1, nil
}

func GetTicketByCinemaIdAndDateDay(cinema_id, date_day int64) ( *models.Ticks, error) {
	p := new(models.Sche)
	sqlStr := "select id from showschdule where cinema_id = ? and date_day = ? and is_delete = -1"
	err := db.Select(&p.Id, sqlStr, cinema_id, date_day)
	if err != nil {
		return nil, err
	}

	sqlStr1 := "select id, schedule_id, cinema_id, movie_id, seat_id, status from ticket where schedule_id = ? and is_delete = -1 "
	p1 := new(models.Ticks)
	err = db.Select(&p1.List, sqlStr1, p.Id)
	if err != nil {
		return nil, err
	}
	return p1, nil
}

