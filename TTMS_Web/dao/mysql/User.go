package mysql

import (
	"TTMS/models"
	"TTMS/pkg/utils"
	//"github.com/gin-gonic/gin"
)

// GetUserByUsername 获取所有信息
func GetUserByUsername(username string) (*models.User, error) {
	user := new(models.User)
	user.ID = 0
	sqlStr := "select id, username, password, email, phone_number, identity, is_login, is_delete from user where username = ?"
	err := db.Get(user, sqlStr, username)
	if err != nil {
		return user, err
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

func GetAllUser(page_num, page_size int, key_word string) (*models.UserList, error) {
	users := new(models.UserList)
	sqlStr := "select id, username, password, email, phone_number, is_login, identity from user where is_delete = -1 and username like ? limit ?,?"
	err := db.Select(&users.List, sqlStr,"%"+key_word+"%",(page_num-1)*page_size, page_size)
	
	if err != nil {
		return nil, err
	}
	
	sqlStr1 := "select count(id) from user where is_delete = -1 and username like ?"

	err = db.Get(&users.Total, sqlStr1, "%" + key_word + "%")

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
	sqlStr := "insert into user (id, username, password, email, phone_number, identity)values (?, ?, ?, ?, ?, 2)"
	_, err := db.Exec(sqlStr, p.ID, p.Username, p.Password, p.Email, p.PhoneNumber)
	return err
}

// UpdateMsg 修改信息---删除，拉黑，修改...
func UpdateMsg(p *models.User) error {
	sqlStr := "update user set username=?, password=? , phone_number=?, email=?, is_delete=?, identity=?, is_login=? where id = ?"
	_, err := db.Exec(sqlStr, p.Username, p.Password, p.PhoneNumber, p.Email, p.IsDelete, p.Identity, p.IsLogin, p.ID)
	return err
}
