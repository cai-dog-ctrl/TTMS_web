package controllers

import (
	"TTMS/models"
	"TTMS/pkg/utils"
	"TTMS/service"
	"errors"
	"fmt"
	"os"

	// "path/filepath"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//有关电影的controller代码

func GetFrontPage(c *gin.Context) {
	p := new(models.ParamsFrontPage)
	p.CarouselNum = utils.ShiftToNum(c.Query("CarouselNum"))
	p.ShowingNum = utils.ShiftToNum(c.Query("ShowingNum"))
	p.ComingNum = utils.ShiftToNum(c.Query("ComingNum"))
	p.ScoreRankingNum = utils.ShiftToNum(c.Query("ScoreRankingNum"))
	p.BoxOfficeRankingNum = utils.ShiftToNum(c.Query("BoxOfficeRankingNum"))
	front_page, err := service.GetFrontPage(p)
	if err != nil {
		zap.L().Error("service.GetFrontPage Error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, front_page)
}

func GetMovieInfoByID(c *gin.Context) {
	p := new(models.ParamsMovie)
	p.Id = c.Param("Id")
	movie, err := service.GetMovieInfoByID(p)
	if err != nil {
		zap.L().Error("service.GetMovieInfoByID ERROR", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	relevantMovies, err := service.GetRelevantMovies(movie.Tag)
	if err != nil {
		zap.L().Error("service.GetRelevantMovies ERROR", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, gin.H{
		"movie":          movie,
		"relevantMovies": relevantMovies,
	})
}

func GetShowingMovies(c *gin.Context) {
	p := new(models.ParamsGetShowingMovies)
	p.Num = utils.ShiftToNum(c.Query("Num"))
	p.Page_num = utils.ShiftToNum(c.Query("Page_num"))
	movie, err := service.GetAllShowingMovies(p)
	if err != nil {
		zap.L().Error("service.GetAllShowingMovies Error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, movie)
}

func GetComingMovies(c *gin.Context) {
	p := new(models.ParamsGetComingMovies)
	p.Num = utils.ShiftToNum(c.Query("Num"))
	p.Page_num = utils.ShiftToNum(c.Query("Page_num"))
	movie, err := service.GetAllComingMovies(p)
	if err != nil {
		zap.L().Error("service.GetAllComingMovies Error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, movie)
}

func GetScoreRankingMovies(c *gin.Context) {
	p := new(models.ParamsGetScoreRankingMovies)
	p.Num = utils.ShiftToNum(c.Query("Num"))
	p.Page_num = utils.ShiftToNum(c.Query("Page_num"))
	movie, err := service.GetAllScoreRankingMovies(p)
	if err != nil {
		zap.L().Error("service.GetAllScoreRankingMovie Error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, movie)
}

func GetBoxOfficeRankingMovies(c *gin.Context) {
	p := new(models.ParamsGetBoxOfficeRankingMovies)
	p.Num = utils.ShiftToNum(c.Query("Num"))
	p.Page_num = utils.ShiftToNum(c.Query("Page_num"))
	movie, err := service.GetAllBoxOfficeRankingMovies(p)
	if err != nil {
		zap.L().Error("service.GetAllBoxOfficeRankingMovies Error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, movie)
}

func AddNewMovie(c *gin.Context) {
	p := new(models.ParamsAddNewMovie)
	form, _ := c.MultipartForm()
	files := form.File["img"]
	for _, file := range files {
		pwd := GetCurrentPath()
		dst := fmt.Sprintf("%v/img/%v", pwd, file.Filename)
		if file.Filename[1] == 'o' {
			p.CoverImgPath = file.Filename
		} else if file.Filename[1] == 'a' {
			p.CarouselImgPath = file.Filename
		}
		c.SaveUploadedFile(file, dst)
	}
	p.Name = c.PostForm("name")
	p.Description = c.PostForm("description")
	p.Tag = c.PostForm("tag")
	p.Duration = utils.ShiftToNum(c.PostForm("duration"))
	p.Up_time = c.PostForm("up_time")
	p.Down_time = c.PostForm("down_time")
	p.Zone = c.PostForm("zone")
	err := service.AddNewMovie(p)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		zap.L().Error("service.AddNewMovie Error", zap.Error(err))
		return
	}
	ResponseSuccess(c, "add new movie successful.")
}

func ModifyMovieByID(c *gin.Context) {
	p := new(models.ParamsModifyMovie)
	err := c.ShouldBind(&p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("ModifyMovieByID ShouldBind Error", zap.Error(err))
		return
	}
	err = service.ModifyMovieByID(p)
	if err != nil {
		if err == errors.New("schedules exist, delete failed") {
			ResponseErrorWithMsg(c, CodeServerBusy, "schedules exist, delete failed")
		} else {
			ResponseError(c, CodeServerBusy)
			zap.L().Error("service.ModifyMovieByID Error", zap.Error(err))
		}
		return
	}
	ResponseSuccess(c, "modify movie successful.")
}

func GetAllMovies(c *gin.Context) {
	p := new(models.ParamsGetAllMovies)
	p.Num = utils.ShiftToNum(c.Query("Num"))
	p.Page_num = utils.ShiftToNum(c.Query("Page_num"))
	p.FlagOfType = c.Query("FlagOfType")
	p.FlagOfZone = c.Query("FlagOfZone")
	p.FlagOfShowing = c.Query("FlagOfShowing")
	p.FlagOfOrder = c.Query("FlagOfOrder")
	movie, err := service.GetAllMovies(p)
	if err != nil {
		zap.L().Error("service.GetAllMovies Error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, movie)
}

func GetMovieByName(c *gin.Context) {
	p := new(models.ParamsGetMovieByName)
	p.Num = utils.ShiftToNum(c.Query("Num"))
	p.Page_num = utils.ShiftToNum(c.Query("Page_num"))
	p.Name = c.Query("Name")
	movie, err := service.GetMovieByName(p)
	if err != nil {
		zap.L().Error("service.GetMovieByName Error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, movie)
}

func GetMovieSort(c *gin.Context) {
	tag := []string{"全部", "爱情", "喜剧", "动画", "剧情", "恐怖", "惊悚", "科幻",
		"动作", "悬疑", "犯罪", "冒险", "战争", "奇幻", "运动", "家庭", "古装", "武侠",
		"西部", "历史", "传记", "歌舞", "黑色电影", "短片", "纪录片", "戏曲", "音乐",
		"灾难", "青春", "儿童", "其他"}
	zone := []string{"全部", "中国大陆", "美国", "韩国", "日本", "中国香港", "中国台湾",
		"泰国", "印度", "法国", "英国", "俄罗斯", "意大利", "西班牙", "德国", "波兰",
		"澳大利亚", "伊朗", "其他"}
	showing := []string{"正在上映", "即将上映", "经典影片"}
	order := []string{"按热门排序", "按时间排序", "按评价排序"}
	ResponseSuccess(c, gin.H{
		"tag":     tag,
		"zone":    zone,
		"showing": showing,
		"order":   order,
	})
}

func UploadPicture(c *gin.Context) {
	file, err := c.FormFile("img")
	if err != nil {
		zap.L().Error("UploadPicture Error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	pwd := GetCurrentPath()
	dst := fmt.Sprintf("%v/img/%v", pwd, file.Filename)
	c.SaveUploadedFile(file, dst)
	ResponseSuccess(c, "upload picture successfully.")
}

func GetCurrentPath() string {
	path, _ := os.Getwd()
	return path
}
