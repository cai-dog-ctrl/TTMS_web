package mysql

import (
	"TTMS/models"
	"TTMS/pkg/utils"
	"fmt"
	"go.uber.org/zap"
	"time"
)

//有关电影的持久化代码

func GetCarouselList(num int) (*models.CarouselList, error) {
	carouselList := new(models.CarouselList)
	sqlStr := "select id, img from movie_info"
	err := db.Select(&carouselList.CarouselList, sqlStr)
	if err != nil {
		zap.L().Error(sqlStr)
		return nil, err
	}
	return carouselList, nil
}

func GetComingList(num int, page_num int) (*models.ComingList, error) {
	comingList := new(models.ComingList)
	day := time.Now().Day()
	month := utils.ShiftToNum(time.Now().Month().String())
	year := time.Now().Year()
	date := day + month*100 + year*100*100
	sqlStr1 := fmt.Sprintf("select count(*) from movie_info where %v < date and is_delete = 0", date)
	err := db.Get(&comingList.Number, sqlStr1)
	if err != nil {
		zap.L().Error(sqlStr1)
		return nil, err
	}
	startIdx := num * page_num
	sqlStr2 := fmt.Sprintf("select id, name, cover_img from movie_info where %v < date and is_delete = 0 limit %v, %v", date, startIdx, num)
	err = db.Select(&comingList.ComingList, sqlStr2)
	if err != nil {
		zap.L().Error(sqlStr2)
		return nil, err
	}
	return comingList, nil
}

func GetShowingList(num int, page_num int) (*models.ShowingList, error) {
	showingList := new(models.ShowingList)
	day := time.Now().Day()
	month := int(time.Now().Month())
	year := time.Now().Year()
	date := day + month*100 + year*100*100
	sqlStr1 := fmt.Sprintf("select count(*) from movie_info where %v >= date and %v < down_time and is_delete = 0", date, date)
	err := db.Get(&showingList.Number, sqlStr1)
	if err != nil {
		zap.L().Error(sqlStr1)
		return nil, err
	}
	startIdx := num * page_num
	sqlStr2 := fmt.Sprintf("select id, name, score, cover_img from movie_info where %v >= date and %v < down_time and is_delete = 0 limit %v, %v", date, startIdx, page_num, num)
	err = db.Select(&showingList.ShowingList, sqlStr2)
	if err != nil {
		zap.L().Error(sqlStr2)
		return nil, err
	}
	return showingList, nil
}

func GetScoreRankingList(num int, page_num int) (*models.ScoreRankingList, error) {
	scoreRankingList := new(models.ScoreRankingList)
	startIdx := num * page_num
	sqlStr := fmt.Sprintf("select id, name, score, cover_img from movie_info order by score desc limit %v, %v", startIdx, num)
	err := db.Select(&scoreRankingList.ScoreRankingList, sqlStr)
	if err != nil {
		zap.L().Error(sqlStr)
		return nil, err
	}
	return scoreRankingList, nil
}

func GetBoxOfficeRankingList(num int, page_num int) (*models.BoxOfficeRankingList, error) {
	boxOfficeRankingList := new(models.BoxOfficeRankingList)
	startIdx := num * page_num
	sqlStr := fmt.Sprintf("select id, name, pf, cover_img from movie_info order by pf desc limit %v, %v", startIdx, num)
	err := db.Select(&boxOfficeRankingList.BoxOfficeRankingList, sqlStr)
	if err != nil {
		zap.L().Error(sqlStr)
		return nil, err
	}
	return boxOfficeRankingList, nil
}

func GetMovieInfoByID(id int64) (*models.MovieInfo, error) {
	movie := new(models.MovieInfo)
	sqlStr := fmt.Sprintf("select * from movie_info where id = %v", id)
	err := db.Get(movie, sqlStr)
	if err != nil {
		zap.L().Error(sqlStr)
		return nil, err
	}
	return movie, nil
}
