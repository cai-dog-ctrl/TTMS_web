package models

//存放有关演出计划模型

type ScheduleIn struct {
	ID        int64 
	CinemaId  int64 
	MovieId   int64
	DateDay   int64 
	StartTime int64 
	EndTime   int64 
	Price	  	float32	
}

type ScheduleOut struct {
	ID        int64 	`db:"id" json:"id"`
	CinemaId  int64 	`db:"cinema_id" json:"cinema_id"`
	MovieId   int64 	`db:"movie_id" json:"movie_id"`
	MovieName string	`db:"name" json:"name"`
	DateDay   int64 	`db:"date_day" json:"date_day"`
	StartTime int64 	`db:"start_time" json:"start_time"`
	EndTime   int64 	`db:"end_time" json:"end_time"`
	Price 	  int		`db:"price" json:"price"`
	CinemaName	string	`db:"cinema_name" json:"cinema_name"`
	Type	  	string	`db:"tag" json:"type"`
}

type ScheduleRet struct {
	ID        string 	`json:"id"`
	CinemaId  string	`json:"cinema_id"`
	MovieId   string 	`json:"movie_id"`
	MovieName string	`json:"name"`
	DateDay   string 	`json:"date_day"`
	StartTime string 	`json:"start_time"`
	EndTime   string 	`json:"end_time"`
	Price 	  int		`json:"price"`
	CinemaName	string	`json:"cinema_name"`
	Type	  	string	`json:"type"`
}

type ScheduleRetList struct {
	Total 	int 			`json:"Total"`
	List 	[]ScheduleRet 	`json:"list"`
}

type SCheduledata struct {
	ID        int64 	
	CinemaId  int64 	
	MovieId   int64 	
	DateDay   int64 	
	StartTime int64 	
	Price	  	int	
}

type ScheduleList struct {
	Total 	int 			`json:"Total"`
	List 	[]ScheduleOut 	`json:"list"`
}

type Sche struct {
	Time 	int64		`db:"movie_time" json:"movie_time"`
	Id 	int64			`db:"id" json:"id"`
}

type ScheDateDay struct {
	Time 	int64		`db:"date_day" json:"data_day"`
}

type ScheDayRet struct {
	Time 	string		`json:"data_day"`
}

type ScheDay struct {
	Time 	[]ScheDateDay 	`json:"Time"`
}


type ScheRets struct {
	Time 	[]ScheDayRet 	`json:"Time"`
}