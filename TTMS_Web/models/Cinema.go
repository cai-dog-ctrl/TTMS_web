package models

//存放有关影院模型
type SeatInfo struct {
	ID    int64 `json:"id" db:"id"`
	Row   int   `json:"row" db:"row"`
	Col   int   `json:"col" db:"col"`
	State int   `json:"state", db:"state"`
}

type SeatList struct {
	ID       int64   `json:"id"`
	SeatList [][]int `json:"seat_list"`
}

type CinemaInfo struct {
	ID       int64  `json:"id" db:"id"`
	MaxRow   int    `json:"row" db:"row"`
	MaxCol   int    `json:"col" db:"col"`
	Tag      string `json:"tag" db:"tag"`
	IsDelete int    `json:"is_delete" db:"is_delete"`
}
