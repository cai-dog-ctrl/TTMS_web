package mysql

import (
	"TTMS/models"
	"fmt"
	"go.uber.org/zap"
	"time"
)

//有关电影的持久化代码

func GetCarouselList(num int) (*models.CarouselList, error) {
	carouselList := new(models.CarouselList)
	sqlStr := fmt.Sprintf("select id, img from movie_info limit 0, %v", num)
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
	month := int(time.Now().Month())
	year := time.Now().Year()
	date := day + month*100 + year*100*100
	sqlStr1 := fmt.Sprintf("select count(*) from movie_info where %v < date and is_delete = -1", date)
	fmt.Println(sqlStr1)
	err := db.Get(&comingList.Number, sqlStr1)
	if err != nil {
		zap.L().Error(sqlStr1)
		return nil, err
	}
	startIdx := num * page_num
	sqlStr2 := fmt.Sprintf("select id, name, cover_img from movie_info where %v < date and is_delete = -1 limit %v, %v", date, startIdx, num)
	fmt.Println(sqlStr2)
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
	sqlStr1 := fmt.Sprintf("select count(*) from movie_info where %v >= date and %v < down_time and is_delete = -1", date, date)
	err := db.Get(&showingList.Number, sqlStr1)
	if err != nil {
		zap.L().Error(sqlStr1)
		return nil, err
	}
	startIdx := num * page_num
	fmt.Println(sqlStr1)
	sqlStr2 := fmt.Sprintf("select id, name, score, cover_img from movie_info where %v >= date and %v < down_time and is_delete = -1 limit %v, %v", date, date, startIdx, num)
	fmt.Println(sqlStr2)
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
	day := time.Now().Day()
	month := int(time.Now().Month())
	year := time.Now().Year()
	date := day + month*100 + year*100*100
	sqlStr := fmt.Sprintf("select id, name, score, cover_img from movie_info where %v >= date and %v < down_time and is_delete = -1 order by score desc limit %v, %v", date, date, startIdx, num)
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
	day := time.Now().Day()
	month := int(time.Now().Month())
	year := time.Now().Year()
	date := day + month*100 + year*100*100
	sqlStr := fmt.Sprintf("select id, name, pf, cover_img from movie_info where %v >= date and %v < down_time and is_delete = -1 order by pf desc limit %v, %v", date, date, startIdx, num)
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

func InsertMovie(p *models.MovieInfo) error {
	sqlStr := "insert into movie_info (name, description, tag, movie_time, date, score, pf, img, is_delete, cover_img, down_time, zone) values (?, ?, ?, ?, ?, 0, 0, ?, -1, ?, ?, ?)"
	_, err := db.Exec(sqlStr, p.Name, p.Description, p.Tag, p.Duration, p.Up_time, p.CarouselImgPath, p.CoverImgPath, p.Down_time, p.Zone)
	if err != nil {
		zap.L().Error(sqlStr)
		return err
	}
	return nil
}

func ModifyMovie(p *models.MovieInfo) error {
	sqlStr := "update movie_info set name = ?, description = ?, tag = ?, movie_time = ?, date = ?, score = ?, pf = ?, img = ?, is_delete = ?, cover_img = ?, down_time = ?, zone = ? where id = ?"
	_, err := db.Exec(sqlStr, p.Name, p.Description, p.Tag, p.Duration, p.Up_time, p.Score, p.BoxOffice, p.CarouselImgPath, p.IsDelete, p.CoverImgPath, p.Down_time, p.Id, p.Zone)
	if err != nil {
		zap.L().Error(sqlStr)
		return err
	}
	return nil
}

func GetAllMovies(p *models.ParamsGetAllMovies, so int) (*models.MovieList, error) {
	movie := new(models.MovieList)
	var sqlStr string
	startIdx := p.Num * p.Page_num
	day := time.Now().Day()
	month := int(time.Now().Month())
	year := time.Now().Year()
	date := day + month*100 + year*100*100
	switch 11111 & so {
	case 0:
		sqlStr = "select * from movie_info"
	case 1:
		sqlStr = fmt.Sprintf("select * from movie_info where tag = %v", p.FlagOfType)
	case 2:
		sqlStr = fmt.Sprintf("select * from movie_info where zone = %v", p.FlagOfZone)
	case 3:
		sqlStr = fmt.Sprintf("select * from movie_info where tag = %v and zone = %v", p.FlagOfType, p.FlagOfZone)
	case 4:
		if p.FlagOfShowing == "showing" {
			sqlStr = fmt.Sprintf("select * from movie_info where %v >= date and %v < down_time and is_delete = -1 limit %v, %v", date, date, startIdx, p.Num)
		} else if p.FlagOfShowing == "coming" {
			sqlStr = fmt.Sprintf("select * from movie_info where %v < date and is_delete = -1 limit %v, %v", date, startIdx, p.Num)
		}
	case 5:
		if p.FlagOfShowing == "showing" {
			sqlStr = fmt.Sprintf("select * from movie_info where type = %v and %v >= date and %v < down_time and is_delete = -1 limit %v, %v", p.FlagOfType, date, date, startIdx, p.Num)
		} else if p.FlagOfShowing == "coming" {
			sqlStr = fmt.Sprintf("select * from movie_info where type = %v and %v < date and is_delete = -1 limit %v, %v", p.FlagOfType, date, startIdx, p.Num)
		}
	case 6:
		if p.FlagOfShowing == "showing" {
			sqlStr = fmt.Sprintf("select * from movie_info where zone = %v and %v >= date and %v < down_time and is_delete = -1 limit %v, %v", p.FlagOfZone, date, date, startIdx, p.Num)
		} else if p.FlagOfShowing == "coming" {
			sqlStr = fmt.Sprintf("select * from movie_info where zone = %v and %v < date and is_delete = -1 limit %v, %v", p.FlagOfZone, date, startIdx, p.Num)
		}
	case 7:
		if p.FlagOfShowing == "showing" {
			sqlStr = fmt.Sprintf("select * from movie_info where type = %v and zone = %v and %v >= date and %v < down_time and is_delete = -1 limit %v, %v", p.FlagOfType, p.FlagOfZone, date, date, startIdx, p.Num)
		} else if p.FlagOfShowing == "coming" {
			sqlStr = fmt.Sprintf("select * from movie_info where type = %v and zone = %v and %v < date and is_delete = -1 limit %v, %v", p.FlagOfType, p.FlagOfZone, date, startIdx, p.Num)
		}
	case 8:
		if p.FlagOfOrder == "score" {
			//Score
			sqlStr = fmt.Sprintf("select * from movie_info order by score desc limit %v, %v", startIdx, p.Num)
		} else if p.FlagOfOrder == "boxoffice" {
			//Boxoffice
			sqlStr = fmt.Sprintf("select * from movie_info order by pf desc limit %v, %v", startIdx, p.Num)
		} else if p.FlagOfOrder == "date" {
			//Date
			sqlStr = fmt.Sprintf("select * from movie_info order by date desc limit %v, %v", startIdx, p.Num)
		}
	case 9:
		if p.FlagOfOrder == "score" {
			//Score
			sqlStr = fmt.Sprintf("select * from movie_info where type = %v order by score desc limit %v, %v", p.FlagOfType, startIdx, p.Num)
		} else if p.FlagOfOrder == "boxoffice" {
			//Boxoffice
			sqlStr = fmt.Sprintf("select * from movie_info where type = %v order by pf desc limit %v, %v", p.FlagOfType, startIdx, p.Num)
		} else if p.FlagOfOrder == "date" {
			//Date
			sqlStr = fmt.Sprintf("select * from movie_info where type = %v order by date desc limit %v, %v", p.FlagOfType, startIdx, p.Num)
		}
	case 10:
		if p.FlagOfOrder == "score" {
			//Score
			sqlStr = fmt.Sprintf("select * from movie_info where zone = %v order by score desc limit %v, %v", p.FlagOfZone, startIdx, p.Num)
		} else if p.FlagOfOrder == "boxoffice" {
			//Boxoffice
			sqlStr = fmt.Sprintf("select * from movie_info where zone = %v order by pf desc limit %v, %v", p.FlagOfZone, startIdx, p.Num)
		} else if p.FlagOfOrder == "date" {
			//Date
			sqlStr = fmt.Sprintf("select * from movie_info where zone = %v order by date desc limit %v, %v", p.FlagOfZone, startIdx, p.Num)
		}
	case 11:
		if p.FlagOfOrder == "score" {
			//Score
			sqlStr = fmt.Sprintf("select * from movie_info where type = %v and zone = %v order by score desc limit %v, %v", p.FlagOfType, p.FlagOfZone, startIdx, p.Num)
		} else if p.FlagOfOrder == "boxoffice" {
			//Boxoffice
			sqlStr = fmt.Sprintf("select * from movie_info where type = %v and zone = %v order by pf desc limit %v, %v", p.FlagOfType, p.FlagOfZone, startIdx, p.Num)
		} else if p.FlagOfOrder == "date" {
			//Date
			sqlStr = fmt.Sprintf("select * from movie_info where type = %v and zone = %v order by date desc limit %v, %v", p.FlagOfType, p.FlagOfZone, startIdx, p.Num)
		}
	case 12:
		if p.FlagOfShowing == "showing" {
			if p.FlagOfOrder == "score" {
				//Score
				sqlStr = fmt.Sprintf("select * from movie_info where %v >= date and %v < down_time order by score desc limit %v, %v", date, date, startIdx, p.Num)
			} else if p.FlagOfOrder == "boxoffice" {
				//Boxoffice
				sqlStr = fmt.Sprintf("select * from movie_info where %v >= date and %v < down_time order by pf desc limit %v, %v", date, date, startIdx, p.Num)
			} else if p.FlagOfOrder == "date" {
				//Date
				sqlStr = fmt.Sprintf("select * from movie_info where %v >= date and %v < down_time order by date desc limit %v, %v", date, date, startIdx, p.Num)
			}
		} else if p.FlagOfShowing == "coming" {
			if p.FlagOfOrder == "score" {
				//Score
				sqlStr = fmt.Sprintf("select * from movie_info where %v < date order by score desc limit %v, %v", date, startIdx, p.Num)
			} else if p.FlagOfOrder == "boxoffice" {
				//Boxoffice
				sqlStr = fmt.Sprintf("select * from movie_info where %v < date order by pf desc limit %v, %v", date, startIdx, p.Num)
			} else if p.FlagOfOrder == "date" {
				//Date
				sqlStr = fmt.Sprintf("select * from movie_info where %v < date order by date desc limit %v, %v", date, startIdx, p.Num)
			}
		}
	case 13:
		if p.FlagOfShowing == "showing" {
			if p.FlagOfOrder == "score" {
				//Score
				sqlStr = fmt.Sprintf("select * from movie_info where %v >= date and %v < down_time and type = %v order by score desc limit %v, %v", date, date, p.FlagOfType, startIdx, p.Num)
			} else if p.FlagOfOrder == "boxoffice" {
				//Boxoffice
				sqlStr = fmt.Sprintf("select * from movie_info where %v >= date and %v < down_time and type = %v order by pf desc limit %v, %v", date, date, p.FlagOfType, startIdx, p.Num)
			} else if p.FlagOfOrder == "date" {
				//Date
				sqlStr = fmt.Sprintf("select * from movie_info where %v >= date and %v < down_time and type = %v order by date desc limit %v, %v", date, date, p.FlagOfType, startIdx, p.Num)
			}
		} else if p.FlagOfShowing == "coming" {
			if p.FlagOfOrder == "score" {
				//Score
				sqlStr = fmt.Sprintf("select * from movie_info where %v < date and type = %v order by score desc limit %v, %v", date, p.FlagOfType, startIdx, p.Num)
			} else if p.FlagOfOrder == "boxoffice" {
				//Boxoffice
				sqlStr = fmt.Sprintf("select * from movie_info where %v < date and type = %v order by pf desc limit %v, %v", date, p.FlagOfType, startIdx, p.Num)
			} else if p.FlagOfOrder == "date" {
				//Date
				sqlStr = fmt.Sprintf("select * from movie_info where %v < date and type = %v order by date desc limit %v, %v", date, p.FlagOfType, startIdx, p.Num)
			}
		}
	case 14:
		if p.FlagOfShowing == "showing" {
			if p.FlagOfOrder == "score" {
				//Score
				sqlStr = fmt.Sprintf("select * from movie_info where %v >= date and %v < down_time and zone = %v order by score desc limit %v, %v", date, date, p.FlagOfZone, startIdx, p.Num)
			} else if p.FlagOfOrder == "boxoffice" {
				//Boxoffice
				sqlStr = fmt.Sprintf("select * from movie_info where %v >= date and %v < down_time and zone = %v order by pf desc limit %v, %v", date, date, p.FlagOfZone, startIdx, p.Num)
			} else if p.FlagOfOrder == "date" {
				//Date
				sqlStr = fmt.Sprintf("select * from movie_info where %v >= date and %v < down_time and zone = %v order by date desc limit %v, %v", date, date, p.FlagOfZone, startIdx, p.Num)
			}
		} else if p.FlagOfShowing == "coming" {
			if p.FlagOfOrder == "score" {
				//Score
				sqlStr = fmt.Sprintf("select * from movie_info where %v < date and zone = %v order by score desc limit %v, %v", date, p.FlagOfZone, startIdx, p.Num)
			} else if p.FlagOfOrder == "boxoffice" {
				//Boxoffice
				sqlStr = fmt.Sprintf("select * from movie_info where %v < date and zone = %v order by pf desc limit %v, %v", date, p.FlagOfZone, startIdx, p.Num)
			} else if p.FlagOfOrder == "date" {
				//Date
				sqlStr = fmt.Sprintf("select * from movie_info where %v < date and zone = %v order by date desc limit %v, %v", date, p.FlagOfZone, startIdx, p.Num)
			}
		}
	case 15:
		if p.FlagOfShowing == "showing" {
			if p.FlagOfOrder == "score" {
				//Score
				sqlStr = fmt.Sprintf("select * from movie_info where %v >= date and %v < down_time and zone = %v and type = %v order by score desc limit %v, %v", date, date, p.FlagOfZone, p.FlagOfType, startIdx, p.Num)
			} else if p.FlagOfOrder == "boxoffice" {
				//Boxoffice
				sqlStr = fmt.Sprintf("select * from movie_info where %v >= date and %v < down_time and zone = %v and type = %v order by pf desc limit %v, %v", date, date, p.FlagOfZone, p.FlagOfType, startIdx, p.Num)
			} else if p.FlagOfOrder == "date" {
				//Date
				sqlStr = fmt.Sprintf("select * from movie_info where %v >= date and %v < down_time and zone = %v and type = %v order by date desc limit %v, %v", date, date, p.FlagOfZone, p.FlagOfType, startIdx, p.Num)
			}
		} else if p.FlagOfShowing == "coming" {
			if p.FlagOfOrder == "score" {
				//Score
				sqlStr = fmt.Sprintf("select * from movie_info where %v < date and zone = %v and type = %v order by score desc limit %v, %v", date, p.FlagOfZone, p.FlagOfType, startIdx, p.Num)
			} else if p.FlagOfOrder == "boxffice" {
				//Boxoffice
				sqlStr = fmt.Sprintf("select * from movie_info where %v < date and zone = %v and type = %v order by pf desc limit %v, %v", date, p.FlagOfZone, p.FlagOfType, startIdx, p.Num)
			} else if p.FlagOfOrder == "date" {
				//Date
				sqlStr = fmt.Sprintf("select * from movie_info where %v < date and zone = %v and type = %v order by date desc limit %v, %v", date, p.FlagOfZone, p.FlagOfType, startIdx, p.Num)
			}
		}
	}
	err := db.Select(&movie.MovieList, sqlStr)
	if err != nil {
		zap.L().Error(sqlStr)
		return nil, err
	}
	return movie, err
}

func GetMovieByName()(*models.MovieList, error){
	movie := new(models.MovieList)
	return movie, nil
}
