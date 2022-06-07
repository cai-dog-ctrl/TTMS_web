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
	Total     int
	OrderList []OrderInfo
}

type Seatinfo struct {
	Row   int
	Col   int
	Price float32
}

type OrderFront struct {
	OrderID      int64
	Date         string
	Price        float32
	Status       int
	MovieName    string
	CoverImgPath string
	CinemaName   string
	SeatList     []*Seatinfo
	ScheduleDate int64
	ScheduleTime int64
}

type OrderFrontList struct {
	Total          int
	OrderFrontList []*OrderFront
}

type OrderFrontRet struct {
	OrderID      int64
	Date         string
	Price        float32
	Status       int
	MovieName    string
	CoverImgPath string
	CinemaName   string
	SeatList     []*Seatinfo
	ScheduleDate string
	ScheduleTime string
}

type OrderFrontListRet struct {
	Total          int
	OrderFrontList []*OrderFrontRet
}

type OrderData struct {
	MovieName    string			`db:"name" json:"movie_name"`
	CoverImgPath string			`db:"img" json:"move_img"`
	TotalPrice   float32		`db:"price" json:"movie_total_price"`
}

type OrderDataList struct {
	Total 	float32				`json:"all_total_price"`
	List 	[]OrderDataList
}

type MovieId struct {
	Id 		int64 				`db:"id" json:"id"`
}

type MovieIds struct {
	IDS 	[]int64 			`json:"IDS""`
}