package mysql

import "TTMS/models"

//有关用户的持久化代码

func GetUserByUsername(username string)(*models.User ,error){
	user:=new(models.User)
	sqlStr:="select id, username, password, email, phone_number from user where username = ? and is_delete = 0"
	err:=db.Get(user,sqlStr,username)
	if err!=nil{
		return nil, err
	}
	return user,nil
}

func InsertUser(p *models.User) error {
	sqlStr:="insert into user (id, username, password, email, phone_number)values (?, ?, ?, ?, ?)"
	_,err:=db.Exec(sqlStr,p.ID,p.Username,p.Password,p.Email,p.PhoneNumber)
	return err
}