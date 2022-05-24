package mysql

import "TTMS/models"

//有关演出计划的持久化代码

// InsertSchedule 添加演出计划
func InsertSchedule(p *models.ScheduleOut) error {
	sqlStr := "insert into showschdule (id, cinema_id, movie_id, date_day, start_time, end_time) values (?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(sqlStr, p.ID, p.CinemaId, p.MovieId, p.DateDay, p.StartTime, p.EndTime)
	return err
}
