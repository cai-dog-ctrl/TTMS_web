package models

//存放有关用户模型

type User struct {
	ID          int64  `db:"id" json:"id,string"`
	Username    string `db:"username" json:"username"`
	Password    string `db:"password" json:"password"`
	Email       string `db:"email" json:"email"`
	PhoneNumber string  `db:"phone_number" json:"phone_number"`
	IsLogin     int    `db:"is_login" json:"is_login"`
	Identity    int    `db:"identity" json:"identity"`
	IsDelete    int    `db:"is_delete" json:"is_delete"`
	Token       string
}

type Data struct {
	ID          int64  `db:"id" json:"id,string"`
	Username    string `db:"username" json:"username"`
	Password    string `db:"password" json:"password"`
	Email       string `db:"email" json:"email"`
	PhoneNumber int64  `db:"phone_number" json:"phone_number"`
	IsLogin     int    `db:"is_login" json:"is_login"`
	Identity    int    `db:"identity" json:"identity"`
}

type UserList struct {
	Total int 	`json:"Total"`
	List []Data `json:"list"`
}
