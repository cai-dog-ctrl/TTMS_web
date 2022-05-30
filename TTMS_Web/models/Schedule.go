package models

//存放有关演出计划模型

type ScheduleIn struct {
	ID        int64 `db:"id" json:"id"`
	CinemaId  int64 `db:"cinema_id" json:"cinema_id"`
	MovieId   int64 `db:"movie_id" json:"movie_id"`
	DateDay   int64 `db:"date_day" json:"date_day"`
	StartTime int64 `db:"start_time" json:"start_time"`
}

type ScheduleOut struct {
	ID        int64 `db:"id" json:"id"`
	CinemaId  int64 `db:"cinema_id" json:"cinema_id"`
	MovieId   int64 `db:"movie_id" json:"movie_id"`
	DateDay   int64 `db:"date_day" json:"date_day"`
	StartTime int64 `db:"start_time" json:"start_time"`
	EndTime   int64 `db:"end_time" json:"end_time"`
}
