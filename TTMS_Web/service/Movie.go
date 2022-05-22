package service

//有关电影的业务逻辑代码

import (
	"TTMS/dao/mysql"
	"TTMS/models"
)

func GetFrontPage(p *models.ParamsFrontPage) (*models.FrontPage, error) {
	var err error
	front_page := new(models.FrontPage)
	front_page.CarouselList_, err = mysql.GetCarouselList(p.CarouselNum)
	if err != nil {
		return nil, err
	}
	front_page.ShowingList_, err = mysql.GetShowingList(p.ShowingNUm, 1)
	if err != nil {
		return nil, err
	}
	front_page.ComingList_, err = mysql.GetComingList(p.ComingNum, 1)
	if err != nil {
		return nil, err
	}
	front_page.ScoreRankingList_, err = mysql.GetScoreRankingList(p.ScoreRankingNum, 1)
	if err != nil {
		return nil, err
	}
	front_page.BoxOfficeRankingList_, err = mysql.GetBoxOfficeRankingList(p.BoxOfficeRankingNum, 1)
	if err != nil {
		return nil, err
	}
	return front_page, nil
}

func GetMovieInfo(p *models.ParamsMovie) (*models.MovieInfo, error) {
	movie, err := mysql.GetMovieInfo(p.Id)
	if err != nil {
		return nil, err
	}
	return movie, nil
}
