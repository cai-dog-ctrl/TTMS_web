package models

//存放有关用户模型

type User struct {
	ID          int64  `db:"id" json:"id"`
	Username    string `db:"username" json:"username"`
	Password    string `db:"password" json:"password"`
	Email       string `db:"email" json:"email"`
	PhoneNumber int64  `db:"phone_number" json:"phone_number"`
	Token       string
}
