package model

type User struct {
	Id       int    `gorm:"id" json:"id"`
	Username string `gorm:"username" json:"username"`
	Password string `gorm:"password" json:"password"`
}

func (*User) TableName() string {
	return "user"
}

