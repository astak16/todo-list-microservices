package model

import (
	"fmt"
	"user/global"
	"user/service"
)

type User struct {
	UserID   string `gorm:"primarykey;unique;column:user_id"`
	UserName string `gorm:"unique;column:user_name"`
	Password string `gorm:"column:password"`
}

func InitUser() {
	migartion()
}

func migartion() {
	global.DB.Set("gorm:table_options", fmt.Sprintf("charset=%s", global.MySqlConfig.Charset)).AutoMigrate(&User{})
}

func (u *User) Create(req *service.UserRequest) User {
	// CheckUserIsExist(req.UserName)
	return User{}
}

func CheckUserIsExist(userName string) bool {
	// a := global.DB.Model(&User{}).Where("user_name = ?", userName).First(&User{})
	// fmt.Println(a)
	return true
}
