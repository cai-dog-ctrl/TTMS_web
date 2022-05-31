package models

//存放前端发送请求参数的结构体

// ParamsLogin 登录接收的参数 form为接收的名称，required为必填字段
type ParamsLogin struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type ParamsRegister struct {
	Username    string `json:"username" form:"username" binding:"required"`
	Password    string `json:"password" form:"password" binding:"required"`
	Email       string `json:"email" form:"email" binding:"required"`
	PhoneNumber int64  `json:"phone_number" form:"phone_number" binding:"required"`
}

type ParamsFrontPage struct {
	CarouselNum         int `json:"carousel_num" form:"carousel_num" binding:"required"`
	ShowingNum          int `json:"showing_num" form:"showing_num" binding:"required"`
	ComingNum           int `json:"coming_num" form:"coming_num" binding:"required"`
	ScoreRankingNum     int `json:"scoreRanking_num" form:"scoreRanking_num" binding:"required"`
	BoxOfficeRankingNum int `json:"boxofficeRanking_num" form:"boxofficeRanking_num" binding:"required"`
}

type ParamsMovie struct {
	Id string `json:"id" form:"id" binding:"required"`
}

type ParamsAdminmsg struct {
	Username    string `json:"username" form:"username" binding:"required"`
	Password    string `json:"password" form:"password" binding:"required"`
	Email       string `json:"email" form:"email" binding:"required"`
	PhoneNumber int64  `json:"phone_number" form:"phone_number" binding:"required"`
}

type ParamsUpdateMsg struct {
	Id          string `json:"id" form:"id" binding:"required"`
	Username    string `json:"username" form:"username" binding:"required"`
	Password    string `json:"password" form:"password" binding:"required"`
	Email       string `json:"email" form:"email" binding:"required"`
	PhoneNumber int64  `json:"phone_number" form:"phone_number" binding:"required"`
	IsDelete    int    `json:"is_delete" form:"is_delete" binding:"required"`
	Identity    int    `json:"identity" form:"identity" binding:"required"`
	IsLogin     int    `json:"is_login" form:"is_login" binding:"required"`
}

type ParamAddSchedule struct {
	CinemaId  int64 `json:"cinema_id" form:"cinema_id" binding:"required"`
	MovieId   int64 `json:"movie_id" form:"movie_id" binding:"required"`
	DateDay   int64 `json:"date_day" form:"date_dat" binding:"required"`
	StartTime int64 `json:"start_time" form:"start_time" binding:"required"`
	Price     int   `json:"price" form:"price" binding:"required"`
}

type ParamsUpdateScheduleMsg struct {
	ID        string `json:"id" form:"id" binding:"required"`
	CinemaId  string `json:"cinema_id" form:"cinema_id" binding:"required"`
	MovieId   string `json:"movie_id" form:"movie_id" binding:"required"`
	DateDay   string `json:"date_day" form:"date_day" binding:"required"`
	StartTime string `json:"start_time" form:"start_time" binding:"required"`
	EndTime   string `json:"end_time" form:"end_time" binding:"required"`
	IsDelete  int    `json:"is_delete" form:"is_delete" binding:"required"`
	IsShow    int    `json:"is_show" form:"is_show" binding:"required"`
	Price     int    `json:"price" form:"price" binding:"required"`
}

type ParamsGetShowingMovies struct {
	Num      int `json:"num" form:"num" binding:"required"`
	Page_num int `json:"page_num" form:"page_num" binding:"required"`
}

type ParamsGetComingMovies struct {
	Num      int `json:"num" form:"num" binding:"required"`
	Page_num int `json:"page_num" form:"page_num" binding:"required"`
}

type ParamsGetScoreRankingMovies struct {
	Num      int `json:"num" form:"num" binding:"required"`
	Page_num int `json:"page_num" form:"page_num" binding:"required"`
}

type ParamsGetBoxOfficeRankingMovies struct {
	Num      int `json:"num" form:"num" binding:"required"`
	Page_num int `json:"page_num" form:"page_num" binding:"required"`
}

type ParamsAddNewMovie struct {
	Name            string `json:"name" form:"name" binding:"required"`
	Description     string `json:"description" form:"description" binding:"required"`
	Tag             string `json:"tag" form:"tag" binding:"required"`
	Duration        int    `json:"duration" form:"duration" binding:"required"`
	Up_time         int    `json:"up_time" form:"up_time" binding:"required"`
	Down_time       int    `json:"down_time" form:"down_time" binding:"required"`
	CoverImgPath    string `json:"cover_img_path" form:"cover_img_path" binding:"required"`
	CarouselImgPath string `json:"carousel_img_path" form:"carousel_img_path" binding:"required"`
}

type ParamsModifyMovie struct {
	Id              int64   `json:"id" form:"id" binding:"required"`
	Name            string  `json:"name" form:"name" binding:"required"`
	Description     string  `json:"description" form:"description" binding:"required"`
	Tag             string  `json:"tag" form:"tag" binding:"required"`
	Duration        int     `json:"duration" form:"duration" binding:"required"`
	Up_time         int     `json:"up_time" form:"up_time" binding:"required"`
	Score           float64 `json:"score" form:"score" binding:"required"`
	BoxOffice       float64 `json:"box_office" form:"box_office" binding:"required"`
	CoverImgPath    string  `json:"cover_img_path" form:"cover_img_path" binding:"required"`
	IsDelete        int     `json:"is_delete" form:"is_delete" binding:"required"`
	CarouselImgPath string  `json:"carousel_img_path" form:"carousel_img_path" binding:"required"`
	Down_time       int     `json:"down_time" form:"down_time" binding:"required"`
}

type ParamsGetCinemaByID struct {
	ID int64 `json:"id" form:"id" binding:"required"`
}

type ParamsGetAllCinemas struct {
	Num      int `json:"num" form:"num" binding:"required"`
	Page_num int `json:"page_num" form:"page_num" binding:"required"`
}

type ParamsAddNewCinema struct {
	ID     int64  `json:"id" form:"id" binding:"required"`
	MaxRow int    `json:"row" form:"row" binding:"required"`
	MaxCol int    `json:"col" form:"col" binding:"required"`
	Tag    string `json:"tag" form:"tag" binding:"required"`
}

type ParamsModifyCinema struct {
	ID       int64  `json:"id" form:"id" binding:"required"`
	MaxRow   int    `json:"row" form:"row" binding:"required"`
	MaxCol   int    `json:"col" form:"col" binding:"required"`
	Tag      string `json:"tag" form:"tag" binding:"required"`
	IsDelete int    `json:"is_delete" form:"is_delete" binding:"required"`
}

type ParamsGetSeatByCinemaID struct {
	ID int64 `json:"id" form:"id" binding:"required"`
}

type ParamsModifySeat struct {
	ID       int64   `json:"id" form:"id" binding:"required"`
	SeatList [][]int `json:"seat_list" form:"seat_list" binding:"required"`
}

type ParamsGetAllMovies struct {
	Num           int    `json:"num" form:"num" binding:"required"`
	Page_num      int    `json:"page_num" form:"page_num" binding:"required"`
	FlagOfType    string `json:"flag_type" form:"flag_type" binding:"required"`
	FlagOfZone    string `json:"flag_zone" form:"flag_zone" binding:"required"`
	FlagOfShowing string `json:"flag_showing" form:"flag_showing" binding:"required"`
	FlagOfOrder   string `json:"flag_order" form:"flag_order" binding:"required"`
}

type ParamsGetMovieByName struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Num      int    `json:"num" form:"num" binding:"required"`
	Page_num int    `json:"page_num" form:"page_num" binding:"required"`
}
