package models

type OrderInfo struct {
	ID       int64   `json:"id,string" db:"id"`
	UserID   int64   `json:"user_id,string" db:"user_id"`
	TicketID int64   `json:"ticket_id,string" db:"ticket_id"`
	Date     string  `json:"date" db:"date"`
	Time     string  `json:"time" db:"time"`
	IsDelete int     `json:"is_delete" db:"is_delete"`
	Price    float32 `json:"price" db:"price"`
	Status   int     `json:"status" db:"status"`
}

type OrderList struct {
	Total     int         `json:"total"`
	OrderList []OrderInfo `json:"order_list"`
}

type Seatinfo struct {
	Row   int     `json:"row" db:"roww"`
	Col   int     `json:"col" db:"coll"`
	Price float32 `json:"price" db:"price"`
}

type OrderFront struct {
	OrderID      int64       `json:"order_id" db:"id"`
	Date         string      `json:"date" db:"date"`
	Price        float32     `json:"price"`
	Status       int         `json:"status" db:"status"`
	MovieName    string      `json:"movie_name" db:"name"`
	CoverImgPath string      `json:"cover_img_path" db:"cover_img"`
	CinemaName   string      `json:"cinema_name" db:"cinema_name"`
	SeatList     []*Seatinfo `json:"seat_list"`
	ScheduleDate int64       `json:"schedule_date" db:"date_day"`
	ScheduleTime int64       `json:"schedule_time" db:"start_time"`
}

type OrderFrontList struct {
	Total          int           `json:"total"`
	OrderFrontList []*OrderFront `json:"order_list"`
}

type OrderFrontRet struct {
	OrderID      int64       `json:"order_id,string"`
	Date         string      `json:"date"`
	Price        float32     `json:"price"`
	Status       int         `json:"status"`
	MovieName    string      `json:"movie_name"`
	CoverImgPath string      `json:"cover_img_path"`
	CinemaName   string      `json:"cinema_name"`
	SeatList     []*Seatinfo `json:"seat_list"`
	ScheduleDate string      `json:"schedule_date"`
	ScheduleTime string      `json:"schedule_time"`
}

type OrderFrontListRet struct {
	Total          int
	OrderFrontList []*OrderFrontRet
}

type OrderData struct {
	MovieName    string			`db:"name" json:"movie_name"`
	CoverImgPath string			`db:"img" json:"move_img"`
	TotalPrice   float32		`json:"movie_total_price"`
}

type OrderDataList struct {
	Total 	float32				`json:"all_total_price"`
	List 	[]OrderData
}

type MovieId struct {
	Id 				int64 		`db:"id" json:"id"`
	MovieName   	string		`db:"name" json"movie_name"`
	CoverImgPath 	string		`db:"img" json:"img"`
}

type MovieIds struct {
	IDS 	[]MovieId 			`json:"IDS""`
}
