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