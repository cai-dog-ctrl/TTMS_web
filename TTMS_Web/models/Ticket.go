package models

//存放有关票务模型

// 座位状态
type SeatStatus struct {
	ID          int64  `db:"id" json:"id,string"`
	CinemaId 	int64  `db:"cinema_id" json:"cinema_id,string"`
	Roww		int  `db:"roww" json:"roww"`
	Coll 		int  `db:"coll" json:"coll"`
	Status		int	 `db:"status" json:"status"`
}

type Seats struct {
	Total int 	 `json:"total"`
	List []SeatStatus  `json:"List"`
}

type Order struct {
	UserID int64 `json:"user_id"`
	TicketList []string `json:"ticket_list"`
}