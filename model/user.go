package model

import (
	"fmt"
	"gorm.io/gorm"
	"tstj-based-on-AI/db"
	"tstj-based-on-AI/utils"
)

type User struct {
	gorm.Model
	ID       int64  `json:"id"`   // 用户id
	Name     string `json:"name"` // 用户名称
	Password string `json:"password"`
	Role     int64  `json:"role"`  // 权限：0表示管理员，1表示普通用户
	Email    string `json:"email"` // 邮箱
}

// FindUserWithId 通过用户id判断用户是否存在
func FindUserWithId(id int64) bool {
	user := User{}
	db.Db.Mysql.Where("id = ?", id).Find(&user).Limit(1)
	if user.Name == "" {
		return false
	}
	return true
}

// FindUser 通过username查找用户并返回其id
func FindUser(username string) int64 {
	tbUser := User{}
	db.Db.Mysql.Where("name = ?", username).First(&tbUser)
	return tbUser.ID
}

// FindUserWithName 通过username查找用户并返回其所有信息
func FindUserWithName(username string) User {
	tbUser := User{}
	db.Db.Mysql.Where("name = ?", username).First(&tbUser)
	return tbUser
}

// GetUserData 通过传入的id从数据库获取用户信息
func GetUserData(id int64) User {
	user := User{}
	db.Db.Mysql.Where("id=?", id).First(&user)
	return user
}

// Login 用户登录
func Login(username, password string) bool {
	userData := User{}
	var res bool
	if FindUser(username) == 0 {
		res = false
	} else {
		db.Db.Mysql.Where("name=?", username).First(&userData)
		if userData.Password == password {
			res = true
		} else {
			res = false
		}
	}
	return res
}

// Register 用户注册
func Register(username, password, email string) int64 {

	if FindUser(username) == 0 {
		db.Db.Mysql.Model(&User{}).Create(map[string]interface{}{"name": username, "password": utils.Md5(password), "email": email})
		fmt.Println("用户创建成功")
		return 0
	}
	fmt.Println("用户已存在，创建失败")
	return 1
}
