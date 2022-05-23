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
	Id int64 `json:"id" form:"id" binding:"required"`
}

type ParamsAdminmsg struct {
	Username    string `json:"username" form:"username" binding:"required"`
	Password    string `json:"password" form:"password" binding:"required"`
	Email       string `json:"email" form:"email" binding:"required"`
	PhoneNumber int64  `json:"phone_number" form:"phone_number" binding:"required"`
	Identity    int    `json:"identity" from:"identity" binding:"required"`
}

type ParamsUpdateMsg struct {
	Id          int64  `json:"id" from:"id" binding:"required"`
	Username    string `json:"username" form:"username" binding:"required"`
	Password    string `json:"password" form:"password" binding:"required"`
	Email       string `json:"email" form:"email" binding:"required"`
	PhoneNumber int64  `json:"phone_number" form:"phone_number" binding:"required"`
	IsDelete    int    `json:"is_delete" from:"id_delete" binding:"required"`
	Identity    int    `json:"identity" from:"identity" binding:"required"`
	IsLogin     int    `json:"is_login" from:"is_login" binding:"required"`
}

type ParamsGetShowingMovies struct {
	Num      int `json:"page" form:"num" binding:"required"`
	Page_num int `json:"page_num" form:"page_num" binding:"required"`
}

type ParamsGetComingMovies struct {
	Num      int `json:"page" form:"num" binding:"required"`
	Page_num int `json:"page_num" form:"page_num" binding:"required"`
}

type ParamsGetScoreRankingMovies struct {
	Num      int `json:"page" form:"num" binding:"required"`
	Page_num int `json:"page_num" form:"page_num" binding:"required"`
}

type ParamsGetBoxOfficeRankingMovies struct {
	Num      int `json:"page" form:"num" binding:"required"`
	Page_num int `json:"page_num" form:"page_num" binding:"required"`
}

type ParamsAddNewMovie struct {
	Name            string `json:"name" form:"name" binding:"required"`
	Description     string `json:"description" form:"description" binding:"required"`
	Tag             string `json:"tag" form:"tag" binding:"required"`
	Duration        int    `json:"duration" form:"duration" binding:"required"`
	Up_time         int    `json:"up_time" form:"up_time" binding:"required"`
	Down_time       int    `json:"down_time" form:"down_time" binding:"required"`
	CoverImgPath    string `json:"coverImgPath" form:"coverImgPath" binding:"required"`
	CarouselImgPath string `json:"carouselImgPath" form:"carouselImgPath" binding:"required"`
}

type ParamsModifyMovie struct {
	Id              int64   `json:"id" form:"id" binding:"required"`
	Name            string  `json:"name" form:"name" binding:"required"`
	Description     string  `json:"description" form:"description" binding:"required"`
	Tag             string  `json:"tag" form:"tag" binding:"required"`
	Duration        int     `json:"duration" form:"duration" binding:"required"`
	Up_time         int     `json:"up_Time" form:"up_time" binding:"required"`
	Score           float64 `json:"score" form:"score" binding:"required"`
	BoxOffice       float64 `json:"box_office" form:"box_office" binding:"required"`
	CoverImgPath    string  `json:"cover_img_path" form:"cover_img_path" binding:"required"`
	IsDelete        int     `json:"isDelete" form:"isDelete" binding:"required"`
	CarouselImgPath string  `json:"carousel_img_path" form:"carousel_img_path" binding:"required"`
	Down_time       int     `json:"down_Time" form:"down_time" binding:"required"`
}
