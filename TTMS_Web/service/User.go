package service

import (
	"TTMS/dao/mysql"
	"TTMS/models"
	"TTMS/pkg/jwt"
	"TTMS/pkg/snowflake"
	"errors"
)
//有关用户的业务逻辑代码

// Login 登录功能
func Login(p *models.ParamsLogin) (*models.User, error) {

	p1, err := mysql.GetUserByUsername(p.Username)
	if err != nil {
		return nil, err
	}
	if p.Username == p1.Username && p.Password == p1.Password {
		Token, err := jwt.GenToken(p1.ID, p1.Username)
		if err == nil {
			p1.Token = Token

		} else {
			return nil, errors.New("请求失败")
		}
		return p1, nil
	}

	return nil, errors.New("登录失败")
}

// Register 注册功能
func Register(p *models.ParamsRegister) error {
	User := new(models.User)
	User.Username = p.Username
	User.Password = p.Password
	User.Email = p.Email
	User.PhoneNumber = p.PhoneNumber
	User.ID = snowflake.GenID()
	err := mysql.InsertUser(User)
	if err != nil {
		return err
	}
	return nil
}
