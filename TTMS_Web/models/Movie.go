package models

//存放有关电影模型
type MovieInfo struct {
	Id              int64   `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Tag             string  `json:"tag"`
	Duration        int     `json:"duration"`
	Up_time         string  `json:"up_Time"`
	Score           float64 `json:"score"`
	BoxOffice       float64 `json:"box_office"`
	CoverImgPath    string  `json:"cover_img_path"`
	IsDelete        int     `json:"isDelete"`
	CarouselImgPath string  `json:"carousel_img_path"`
	Down_time       string  `json:"down_Time"`
}

type Carousel struct {
	Id              int64  `json:"id"`
	CarouselImgPath string `json:"carousel_img_path"`
}

type Showing struct {
	Id           int64   `json:"id"`
	Name         string  `json:"name"`
	Score        float64 `json:"score"`
	CoverImgPath string  `json:"cover_img_path"`
}

type Coming struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	CoverImgPath string `json:"cover_img_path"`
}

type ScoreRanking struct {
	Id           int64   `json:"id"`
	Name         string  `json:"name"`
	Score        float64 `json:"score"`
	CoverImgPath string  `json:"cover_img_path"`
}

type BoxOfficeRanking struct {
	Id           int64   `json:"id"`
	Name         string  `json:"name"`
	BoxOffice    float64 `json:"box_office"`
	CoverImgPath string  `json:"cover_img_path"`
}

type CarouselList struct {
	CarouselList []*Carousel `json:"CarouselList"`
}

type ShowingList struct {
	ShowingList []*Showing `json:"ShowingList"`
}

type ComingList struct {
	ComingList []*Coming `json:"ComingList"`
}

type ScoreRankingList struct {
	ScoreRankingList []*ScoreRanking `json:"ScoreRankingList"`
}

type BoxOfficeRankingList struct {
	BoxOfficeRankingList []*BoxOfficeRanking `json:"BoxOfficeRankingList"`
}

type FrontPage struct {
	CarouselList_         *CarouselList
	ShowingList_          *ShowingList
	ComingList_           *ComingList
	ScoreRankingList_     *ScoreRankingList
	BoxOfficeRankingList_ *BoxOfficeRankingList
}
