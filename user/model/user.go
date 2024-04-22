package model

import (
	"errors"
	"fmt"
	"time"
	"user/global"
	"user/service"

	"gorm.io/gorm"
)

type User struct {
	ID        uint       `gorm:"primarykey;unique;column:id"`
	UserName  string     `gorm:"unique;column:user_name"`
	Password  string     `gorm:"column:password"`
	CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt *time.Time `gorm:"column:deleted_at;default:null"`
}

func InitUser() {
	migartion()
}

func migartion() {
	global.DB.Set("gorm:table_options", fmt.Sprintf("charset=%s", global.MySqlConfig.Charset)).AutoMigrate(&User{})
}

func (u *User) CreateUser(req *service.UserRequest) (*User, error) {
	exist := u.CheckUserIsExist(req.UserName)
	if exist {
		return nil, errors.New("用户已经存在")
	}

	user := User{
		UserName: req.UserName,
		Password: req.Password,
	}
	if err := global.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) GetUser(req *service.UserRequest) (*User, error) {
	var user User
	if exist := u.CheckUserIsExist(req.UserName); !exist {
		return nil, errors.New("用户不存在")
	}
	if err := global.DB.Where("user_name = ? AND password = ?", req.UserName, req.Password).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) CheckUserIsExist(userName string) bool {
	if err := global.DB.Model(&User{}).Where("user_name = ?", userName).First(&u).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}
