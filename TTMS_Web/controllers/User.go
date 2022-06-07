package controllers

import (
	"TTMS/models"
	"TTMS/pkg/utils"
	"TTMS/service"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//有关用户的controller代码
// const RootPath="./img/"
func Login(c *gin.Context){
	p:=new(models.ParamsLogin)
<<<<<<< HEAD
	fmt.Println(p)
	err := c.ShouldBind(&p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
=======
	//fmt.Println(p)
	err:=c.ShouldBind(&p)
	if err!=nil{
		ResponseError(c,CodeInvalidParams)
>>>>>>> 6d4bad3732f772e7a3c44b5350feb6e790a2e8cf
		zap.L().Error("Login ShouldBind Error", zap.Error(err))
		return
	}
	p1, err := service.Login(p)
	if err != nil {
		if err == errors.New("登录失败") {
			ResponseErrorWithMsg(c, CodeInvalidPassword, "登录失败")
			return
		}else if p1.ID == 0 {
			ResponseErrorWithMsg(c, CodeInvalidPassword, "无此用户")
			return
		} else {
			zap.L().Error("Login Service error", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
	}
	if p1.IsLogin == -1 && p1.IsDelete == 1 {
		ResponseErrorWithMsg(c, CodeInvalidPassword, "登录失败")
		return
	}
	ResponseSuccess(c, gin.H{
		"username": p1.Username,
		"token":    p1.Token,
		"identity": p1.Identity,
		"userid": 	utils.ShiftToStringFromInt64(p1.ID),
	})
}

// Register 注册
func Register(c *gin.Context) {
	p := new(models.ParamsRegister)
	err := c.ShouldBind(&p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("Register ShouldBind Error", zap.Error(err))
		return
	}
	err = service.Register(p)
	if err != nil {
		zap.L().Error("service.Register error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, "注册成功")
}

// AddAdmin 添加管理员
func AddAdmin(c *gin.Context) {
	p := new(models.ParamsAdminmsg)
	err := c.ShouldBind(&p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("AddAdmin ShouldBind Error", zap.Error(err))
		return
	}
	err = service.AddAdmin(p)
	if err != nil {
		zap.L().Error("service.AddAdmin error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, "添加成功")
}

// GetAllMsg 获取所有用户信息
func GetAllMsg(c *gin.Context) {

	pageNum := utils.ShiftToNum(c.Query("page_num"))
	pageSize := utils.ShiftToNum(c.Query("page_size"))
	key_word := c.Query("key_word")

	p1, err := service.GetAllMsg(pageNum, pageSize, key_word)

	if err != nil {
		zap.L().Error("service.GetAllMsg error", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, "获取失败")

		return
	}
	ResponseSuccess(c, p1)
}

// GetUserMsgById 获取所有信息
func GetUserMsgById(c *gin.Context) {
	p := c.Param("id")
	if p == "" {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("GetUserMsgById getid Error")
		return
	}
	p1, err := service.GetMsgById(p)

	if err != nil {
		zap.L().Error("service.GetMsgById error", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, "获取失败")
		return
	}
	id := utils.ShiftToStringFromInt64(p1.ID)
	ResponseSuccess(c, gin.H{
		"id":           id,
		"username":     p1.Username,
		"token":        p1.Token,
		"password":     p1.Password,
		"phone_number": p1.PhoneNumber,
		"email":        p1.Email,
		"is_delete":    p1.IsDelete,
		"identity":     p1.Identity,
		"is_login":     p1.IsLogin,
	})
}

// UpdateMsg 更新信息
func UpdateMsg(c *gin.Context) {
	p := new(models.ParamsUpdateMsg)
	err := c.ShouldBind(&p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("UpdateMsg ShouldBind Error", zap.Error(err))
		return
	}
	err = service.UpdateMsg(p)
	if err != nil {
		zap.L().Error("service.UpdateMsg error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, "修改成功")
}

func GetPictureByFileName(c *gin.Context) {
	img := c.Param("img")
	if img == "" {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("GetFile Vaild param")
		return
	}
	RootPath := GetCurrentPath()
	Path := fmt.Sprintf("%v/img/%v", RootPath, img)
	c.File(Path)
}
