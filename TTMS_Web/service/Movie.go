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
	front_page.ShowingList_, err = GetShowingMovies(p.ShowingNum, 0)
	if err != nil {
		return nil, err
	}
	front_page.ComingList_, err = GetComingMovies(p.ComingNum, 0)
	if err != nil {
		return nil, err
	}
	front_page.ScoreRankingList_, err = GetScoreRankingList(p.ScoreRankingNum, 0)
	if err != nil {
		return nil, err
	}
	front_page.BoxOfficeRankingList_, err = GetBoxOfficeRankingList(p.BoxOfficeRankingNum, 0)
	if err != nil {
		return nil, err
	}
	return front_page, nil
}

func GetMovieInfoByID(p *models.ParamsMovie) (*models.MovieInfo, error) {
	movie, err := mysql.GetMovieInfoByID(p.Id)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func GetAllShowingMovies(p *models.ParamsGetShowingMovies) (*models.ShowingList, error) {
	return GetShowingMovies(p.Num, p.Page_num-1)
}

func GetAllComingMovies(p *models.ParamsGetComingMovies) (*models.ComingList, error) {
	return GetComingMovies(p.Num, p.Page_num-1)
}

func GetAllScoreRankingMovies(p *models.ParamsGetScoreRankingMovies) (*models.ScoreRankingList, error) {
	return GetScoreRankingList(p.Num, p.Page_num-1)
}

func GetAllBoxOfficeRankingMovies(p *models.ParamsGetBoxOfficeRankingMovies) (*models.BoxOfficeRankingList, error) {
	return GetBoxOfficeRankingList(p.Num, p.Page_num-1)
}

func GetShowingMovies(num int, page_num int) (*models.ShowingList, error) {
	showingList, err := mysql.GetShowingList(num, page_num)
	if err != nil {
		return nil, err
	}
	return showingList, nil
}

func GetComingMovies(num int, page_num int) (*models.ComingList, error) {
	comingList, err := mysql.GetComingList(num, page_num)
	if err != nil {
		return nil, err
	}
	return comingList, nil
}

func GetScoreRankingList(num int, page_num int) (*models.ScoreRankingList, error) {
	scoreRankingList, err := mysql.GetScoreRankingList(num, page_num)
	if err != nil {
		return nil, err
	}
	return scoreRankingList, nil
}

func GetBoxOfficeRankingList(num int, page_num int) (*models.BoxOfficeRankingList, error) {
	boxOfficeRankingList, err := mysql.GetBoxOfficeRankingList(num, page_num)
	if err != nil {
		return nil, err
	}
	return boxOfficeRankingList, nil
}

func AddNewMovie(p *models.ParamsAddNewMovie) error {
	NewMovie := new(models.MovieInfo)
	NewMovie.Name = p.Name
	NewMovie.Description = p.Description
	NewMovie.Tag = p.Tag
	NewMovie.Duration = p.Duration
	NewMovie.Up_time = p.Up_time
	NewMovie.Down_time = p.Down_time
	NewMovie.CoverImgPath = p.CoverImgPath
	NewMovie.CarouselImgPath = p.CarouselImgPath
	return mysql.InsertMovie(NewMovie) 
}

func ModifyMovieByID(p *models.ParamsModifyMovie) error {
	movie := new(models.MovieInfo)
	movie.Id = p.Id
	movie.Name = p.Name
	movie.Description = p.Description
	movie.Tag = p.Tag
	movie.Duration = p.Duration
	movie.Up_time = p.Up_time
	movie.Score = p.Score
	movie.BoxOffice = p.BoxOffice
	movie.CoverImgPath = p.CoverImgPath
	movie.IsDelete = p.IsDelete
	movie.CarouselImgPath = p.CarouselImgPath
	movie.Down_time = p.Down_time
	return mysql.ModifyMovie(movie)
}