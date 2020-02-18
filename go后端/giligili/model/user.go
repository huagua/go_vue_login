package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User 用户模型
type User struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	Nickname       string
	Status         string
	Avatar         string `gorm:"size:1000"`
}

//火车票模型
type Train struct {
	gorm.Model
	TrainNum       string
	From           string
	To             string
	StartTime      string
	EndTime        string
	Duration       string
	Date           string
}

//空闲火车票模型
type FreeTicket struct {
	gorm.Model
	TrainNum       string
	Date           string
	Carriage       int
	Row            int
	Which          string
	Rank           string
}

type Ticket struct {
	Carriage int
	Row      int
	Which    string
	Money    int
}

type SearchTicket struct {
	From     string
	To       string
	TrainNum string
	Super    interface{}
	First    interface{}
	Second   interface{}
	StartTime interface{}
	Duration  interface{}
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
)

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
