package models

//存放有关影院模型
type SeatInfo struct {
	ID       int64 `json:"id" db:"id"`
	CinemaID int64 `json:"cinema_id" db:"cinema_id"`
	Row      int   `json:"row" db:"roww"`
	Col      int   `json:"col" db:"coll"`
	Status   int   `json:"state" db:"status"`
	Flag     int   `json:"flag" db:"flag"`
}

type Seat struct {
	ID     int64 `json:"id,string" db:"id"`
	Status int   `json:"status" db:"status"`
	Flag   int   `json:"flag" db:"flag"`
}

type SeatList struct {
	ID       int64    `json:"id,string"`
	SeatList [][]Seat `json:"seat_list"`
}

type CinemaInfo struct {
	ID       int64  `json:"id,string" db:"id"`
	Name     string `json:"name" db:"cinema_name"`
	MaxRow   int    `json:"row" db:"roww"`
	MaxCol   int    `json:"col" db:"coll"`
	Tag      string `json:"tag" db:"tag"`
	IsDelete int    `json:"is_delete" db:"is_delete"`
}

type CinemaList struct {
	Number     int           `json:"cinema_num"`
	CinemaList []*CinemaInfo `json:"cinemaList"`
}
