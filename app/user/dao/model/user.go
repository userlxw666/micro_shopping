package model

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(30)"`
	Password string `gorm:"type:varchar(100)"`
	Salt     string `gorm:"type:varchar(100)"`
	Token    string `gorm:"type:varchar(500)"`
}

func NewUser(username string, password string) *User {
	return &User{
		UserName: username,
		Password: password,
	}
}

var charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seedRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func (u *User) CreateSalt() {
	b := make([]byte, bcrypt.MaxCost)
	for i := range b {
		b[i] = charset[seedRand.Intn(len(charset))]
	}
	u.Salt = string(b)
}

func (u *User) CreateHashPassword() {
	password := u.Password + u.Salt
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("密码加密失败！", err)
		return
	}
	u.Password = string(hashpassword)
}

func (u *User) CheckPassword(password string) bool {
	hashpassword := u.Password
	err := bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password+u.Salt))
	if err != nil {
		fmt.Println("密码验证失败！", err)
		return false
	}
	return true
}
