package models

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID        int    `gorm:"primary_key,AUTO_INCREMENT" json:"id"`
	Name      string `json:"name"`
	Password  string `json:"-"`
	Mobile    string `json:"mobile"`
	Uuid      string `json:"uuid"`
	CreatedAt int64  `json:"-"`
	UpdatedAt int64  `json:"-"`
}

func (User) TableName() string {
	return "users"
}

func GetUsers(pageNum int, pageSize int, maps interface{}) (count int, users []User) {
	db.Where(maps).Preload("Shop").Limit(pageSize).Offset(pageSize * (pageNum - 1)).Order("id asc").Find(&users)
	db.Where(maps).Preload("Shop").Find(&users).Count(&count)
	return
}

func CreateUser(name, password, mobile string) (user User, err error) {

	exists := findUserByNameOrMobile(name, mobile)

	if exists == true {
		return User{}, errors.New("用户名与手机号重复")
	}

	user = User{Name: name, Password: password, Mobile: mobile, Uuid: uuid.New().String()}

	res := db.Save(&user)
	fmt.Println(res)

	return user, nil
}

func findUserByNameOrMobile(name, mobile string) bool {
	user := User{}
	db.Where("name = ? or mobile = ? ", name, mobile).First(&user)

	if user.ID > 0 {
		return true
	}

	return false
}

func FindUserByMobile(mobile string) (user User) {
	db.Where("mobile = ?", mobile).First(&user)
	return
}

func Login(name, password string) (user User, err error) {
	db.Where("name = ? and password = ?", name, password).First(&user)
	if user.ID > 0 {
		return user, nil
	}
	err = errors.New("账号不存在或密码错误")
	return User{}, err
}

func FindUserByUuid(uuid string) (user User, err error) {
	db.Where("uuid = ?", uuid).First(&user)

	if user.ID > 0 {
		return user, nil
	}
	err = errors.New("账号不存在")
	return User{}, err
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now().Unix())
	return nil
}

func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now().Unix())
	return nil
}
