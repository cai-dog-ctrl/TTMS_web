package models

//存放有关电影模型

const (
	IS_SHOWING = 0
)

type MovieInfo struct {
	Id              int64   `json:"id,string" db:"id"`
	Name            string  `json:"name" db:"name"`
	Description     string  `json:"description" db:"description"`
	Tag             string  `json:"tag" db:"tag"`
	Duration        int     `json:"duration" db:"movie_time"`
	Up_time         string  `json:"up_Time" db:"date"`
	Score           float64 `json:"score" db:"score"`
	BoxOffice       float64 `json:"box_office" db:"pf"`
	CoverImgPath    string  `json:"cover_img_path" db:"cover_img"`
	IsDelete        int     `json:"is_delete" db:"is_delete"`
	CarouselImgPath string  `json:"carousel_img_path" db:"img"`
	Down_time       string  `json:"down_Time" db:"down_time"`
	Zone            string  `json:"zone" db:"zone"`
}

type Carousel struct {
	Id              int64  `json:"id,string" db:"id"`
	CarouselImgPath string `json:"carousel_img_path" db:"img"`
}

type Showing struct {
	Id           int64   `json:"id,string" db:"id"`
	Name         string  `json:"name" db:"name"`
	Score        float64 `json:"score" db:"score"`
	CoverImgPath string  `json:"cover_img_path" db:"cover_img"`
}

type Coming struct {
	Id           int64  `json:"id,string" db:"id"`
	Name         string `json:"name" db:"name"`
	CoverImgPath string `json:"cover_img_path" db:"cover_img"`
}

type ScoreRanking struct {
	Id           int64   `json:"id,string" db:"id"`
	Name         string  `json:"name" db:"name"`
	Score        float64 `json:"score" db:"score"`
	CoverImgPath string  `json:"cover_img_path" db:"cover_img"`
}

type BoxOfficeRanking struct {
	Id           int64   `json:"id,string" db:"id"`
	Name         string  `json:"name" db:"name"`
	BoxOffice    float64 `json:"box_office" db:"pf"`
	CoverImgPath string  `json:"cover_img_path" db:"cover_img"`
}

type MovieList struct {
	Total     int          `json:"total"`
	MovieList []*MovieInfo `json:"movieList"`
}

type CarouselList struct {
	CarouselList []*Carousel `json:"CarouselList"`
}

type ShowingList struct {
	Number      int        `json:"Showing_num"`
	ShowingList []*Showing `json:"ShowingList"`
}

type ComingList struct {
	Number     int       `json:"Coming_num"`
	ComingList []*Coming `json:"ComingList"`
}

type ScoreRankingList struct {
	ScoreRankingList []*ScoreRanking `json:"ScoreRankingList"`
}

type BoxOfficeRankingList struct {
	BoxOfficeRankingList []*BoxOfficeRanking `json:"BoxOfficeRankingList"`
}

type FrontPage struct {
	CarouselList_         *CarouselList         `json:"CarouselList"`
	ShowingList_          *ShowingList          `json:"ShowingList"`
	ComingList_           *ComingList           `json:"ComingList"`
	ScoreRankingList_     *ScoreRankingList     `json:"ScoreRankingList"`
	BoxOfficeRankingList_ *BoxOfficeRankingList `json:"BoxOfficeRankingList"`
}
