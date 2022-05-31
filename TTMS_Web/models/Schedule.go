package models

//存放有关演出计划模型

type ScheduleIn struct {
	ID        int64 `db:"id" json:"id,string"`
	CinemaId  int64 `db:"cinema_id" json:"cinema_id,string"`
	MovieId   int64 `db:"movie_id" json:"movie_id,string"`
	DateDay   int64 `db:"date_day" json:"date_day,string"`
	StartTime int64 `db:"start_time" json:"start_time,string"`
	EndTime   int64 `db:"end_time" json:"end_time,string"`
	Price	  	int		`db:"price" json:"price"`
}

type ScheduleOut struct {
	ID        int64 `db:"id" json:"id,string"`
	CinemaId  int64 `db:"cinema_id" json:"cinema_id,string"`
	MovieId   int64 `db:"movie_id" json:"movie_id,string"`
	DateDay   int64 `db:"date_day" json:"date_day,string"`
	StartTime int64 `db:"start_time" json:"start_time,string"`
	EndTime   int64 `db:"end_time" json:"end_time,string"`
}


type SCheduledata struct {
	ID        int64 	`db:"id" json:"id,string"`
	CinemaId  int64 	`db:"cinema_id" json:"cinema_id,string"`
	MovieId   int64 	`db:"movie_id" json:"movie_id,string"`
	DateDay   int64 	`db:"date_day" json:"date_day,string"`
	StartTime int64 	`db:"start_time" json:"start_time,string"`
	EndTime   int64 	`db:"end_time" json:"end_time,string"`
	IsDelete  int		`db:"is_delete" json:"is_delete"`
	IsShow    int 		`db:"is_show" json:"is_show"`

	Price	  	int		`db:"price" json:"price"`
	CinemaName	string	`db:"cinema_name" json:"cinema_name"`
	Type	  	string 	`db:"type" json:"type"`
}

type ScheduleList struct {
	Total 	int 			`json:"Total"`
	List 	[]ScheduleOut 	`json:"list"`
}