package mysql

import (
	"TTMS/models"
	"TTMS/pkg/utils"
)

// GetUserByUsername 获取所有信息
func GetUserByUsername(username string) (*models.User, error) {
	user := new(models.User)
	sqlStr := "select id, username, password, email, phone_number, identity, is_login, is_delete from user where username = ?"
	err := db.Get(user, sqlStr, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserById(Id string) (*models.User, error) {
	user := new(models.User)
	sqlStr := "select id, username, password, email, phone_number, identity, is_login, is_delete from user where id = ?"

	err := db.Get(user, sqlStr, utils.ShiftToNum64(Id))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetAllUser() (*models.UserList, error) {
	users := new(models.UserList)

	sqlStr := "select id, username, password, phone_number, email, identity, avatar, is_login from user where is_delete = 0"
	err := db.Select(&users.List, sqlStr)
	if err != nil {
		return nil, err
	}
	return users, nil

}

// InsertUser 添加一个普通用户
func InsertUser(p *models.User) error {
	sqlStr := "insert into user (id, username, password, email, phone_number)values (?, ?, ?, ?, ?)"
	_, err := db.Exec(sqlStr, p.ID, p.Username, p.Password, p.Email, p.PhoneNumber)
	return err
}

// InsertAdmin 添加一个管理员
func InsertAdmin(p *models.User) error {
	sqlStr := "insert into user (id, username, password, email, phone_number, identity)values (?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(sqlStr, p.ID, p.Username, p.Password, p.Email, p.PhoneNumber, p.Identity)
	return err
}

// UpdateMsg 修改信息---删除，拉黑，修改...
func UpdateMsg(p *models.User) error {
	sqlStr := "update user set username=?, password=? , phone_number=?, email=?, is_delete=?, identity=?, is_login=? where id = ?"
	_, err := db.Exec(sqlStr, p.Username, p.Password, p.PhoneNumber, p.Email, p.IsDelete, p.Identity, p.IsLogin, p.ID)
	return err
}
