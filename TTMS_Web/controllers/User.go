package controllers

import (
	"TTMS/models"
	"TTMS/service"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//有关用户的controller代码
const RootPath="./img/"
func Login(c *gin.Context){
	p:=new(models.ParamsLogin)
	err:=c.ShouldBind(&p)
	if err!=nil{
		ResponseError(c,CodeInvalidParams)
		zap.L().Error("Login ShouldBind Error")
		return
	}
	p1,err:=service.Login(p)
	if err != nil {
		if err==errors.New("登录失败"){
			ResponseErrorWithMsg(c,CodeInvalidPassword,"登录失败")
			return
		}else {
			zap.L().Error("Login Service error",zap.Error(err))
			ResponseError(c,CodeServerBusy)
			return
		}
		return
	}
	ResponseSuccess(c,gin.H{
		"username":p1.Username,
		"token":p1.Token,
	})
}
func Register(c *gin.Context){
	p:=new(models.ParamsRegister)
	err:=c.ShouldBind(&p)
	if err!=nil{
		ResponseError(c,CodeInvalidParams)
		zap.L().Error("Register ShouldBind Error")
		return
	}
	err=service.Register(p)
	if err != nil {
		zap.L().Error("service.Register error",zap.Error(err))
		ResponseError(c,CodeServerBusy)
		return
	}
	ResponseSuccess(c,"注册成功")
}
func GetPictureByFileName(c *gin.Context){
	img:=c.Param("img")
	if img==""{
		ResponseError(c,CodeInvalidParams)
		zap.L().Error("GetFile Vaild param")
		return
	}
	c.File(RootPath+img)
}