package models

//存放有关票务模型

type Seats struct {
	Total int 	 `json:"total"`
	List []SeatInfo  `json:"List"`
}

type Order struct {
	UserID int64 `json:"user_id"`
	TicketList []string `json:"ticket_list"`
}

type Tick struct {
	Id  			int64	`db:"id" json:"id"`
	ScheduleId 		int64	`db:"schedule_id" json:"schedule_id"`
	CinemaId 		int64 	`db:"cinema_id" json:"cinema_id"`
	MovieId			int64	`db:"movie_id" json:"movie_id"`
	SeatId 			int64	`db:"seat_id" json:"seat_id"`
	Status 			int 	`db:"status" json:"status"`
}

type Ticks struct {
	RowNum 	int 	`json:"row"`
	Total 	int		`json:"total"`
	List 	[]Tick 	`json:"list"`
}

type TickRet struct {
	Id  			string	`json:"id"`
	ScheduleId 		string	`json:"schedule_id"`
	CinemaId 		string	`json:"cinema_id"`
	MovieId			string	`json:"movie_id"`
	SeatId 			string	`json:"seat_id"`
	Status 			int 	`json:"status"`
}

type TickRets struct {
	Total 	int			`json:"total"`
	List 	[]TickRet 	`json:"list"`
}

type TickRetss struct {
	List 	[][]TickRet 	`json:"list"`
}