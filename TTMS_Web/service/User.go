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

// AddAdmin 添加管理员
func AddAdmin(p *models.ParamsAdminmsg) error {
	User := new(models.User)
	User.Username = p.Username
	User.Password = p.Password
	User.Email = p.Email
	User.PhoneNumber = p.PhoneNumber

	User.ID = snowflake.GenID()
	err := mysql.InsertAdmin(User)
	if err != nil {
		return err
	}
	return nil
}

// GetAllMsg 获取所有用户信息
func GetAllMsg() (p *models.UserList, err error) {
	p1, err := mysql.GetAllUser()
	if err != nil {
		return nil, err
	}
	return p1, nil
}

// GetMsgById 根据Id获取信息
func GetMsgById(userId string) (p *models.User, err error) {
	p1, err := mysql.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return p1, nil
}

// UpdateMsg 更新信息
func UpdateMsg(p *models.ParamsUpdateMsg) error {
	User := new(models.User)
	User.Username = p.Username
	User.Password = p.Password
	User.Email = p.Email
	User.PhoneNumber = p.PhoneNumber
	User.IsDelete = p.IsDelete
	User.Identity = p.Identity
	User.IsLogin = p.IsLogin
	User.ID = p.Id
	err := mysql.UpdateMsg(User)
	if err != nil {
		return err
	}
	return nil
}
