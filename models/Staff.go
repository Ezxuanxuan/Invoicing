package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

//1.员工表
type Staff struct {
	gorm.Model
	Name        string `gorm:"type:varchar(20)" `
	EnglishName string `gorm:"type:varchar(20)" `
	Password    string `gorm:"type:char(32)" `
	Birthday    time.Time
	IdCard      string `gorm:"type:varchar(18)" ` //身份证号码
	Telephone   string `gorm:"type:varchar(20)"`  //电话号码
}

//根据用户名获取用户数量
func GetUserCountbyUsername(username string) int {

	var count int
	//查询该用户名是否有账号
	db.Model(&Staff{}).Where("english_name = ?", username).Count(&count)
	return count
}

//根据用户名获取密码
func GetPasswordbyUsername(username string) string {
	var staff Staff
	db.Where("english_name = ?", username).First(&staff)
	return staff.Password
}

//根据用户名获取密码
func GetIdbyUsername(username string) int {
	var staff Staff
	db.Where("english_name = ?", username).First(&staff)
	return staff.ID
}

func IsExitName(Name string) bool {
	var count int
	//查询该用户名是否有账号
	db.Model(&Staff{}).Where("name = ?", Name).Count(&count)
	if count < 1 {
		return false
	}
	return true
}

func IsExitEnglishName(EnglishName string) bool {
	var count int
	//查询该用户名是否有账号
	db.Model(&Staff{}).Where("name = ?", EnglishName).Count(&count)
	if count < 1 {
		return false
	}
	return true
}

//创建用户
func CreateUser(staff Staff) bool {
	db.Create(&staff)
	if db.NewRecord(staff) {
		return false
	}
	return true
}

//获取所有用户信息
func GetAllStaff() []Staff {
	var staffs []Staff
	db.Find(&staffs)

	return staffs

}

//查询该用户id是否存在
func GetStaffById(id int) bool {
	var count int = 0
	//查询该用户名是否有账号
	db.Model(&Staff{}).Where("id = ?", id).Count(&count)
	if count > 0 {
		return true
	} else {
		return false
	}
}

//修改密码
func UpdatePassword(id int, password string) bool {
	var staff Staff
	db.Model(&staff).Where("id = ?", id).Update("password", password)
	return true
}

//修改电话
func UpdateTelephone(id int, telephone string) bool {
	var staff Staff
	db.Model(&staff).Where("id = ?", id).Update("telephone", telephone)
	return true
}
