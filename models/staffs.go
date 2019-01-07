package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Staffs struct {
	Id          int64     `xorm:"not null pk autoincr INT(11)"`
	CreatedAt   time.Time `xorm:"created"`
	UpdatedAt   time.Time `xorm:"updated"`
	DeletedAt   time.Time `xorm:"index TIMESTAMP"`
	Name        string    `xorm:"VARCHAR(30)"`
	EnglishName string    `xorm:"VARCHAR(30)"`
	Password    string    `xorm:"CHAR(32)"`
	Birthday    time.Time `xorm:"TIMESTAMP"`
	IdCard      string    `xorm:"VARCHAR(18)"`
	Telephone   string    `xorm:"VARCHAR(20)"`
}

//根据用户名获取用户数量
func GetUserCountbyUsername(username string) (int64, error) {

	staff := new(Staffs)
	//查询该用户名是否有账号
	total, err := engine.Where("english_name = ?", username).Count(staff)
	if err != nil {
		return 0, err
	}
	return total, nil
	//	.Model(&Staff{}).Where("english_name = ?", username).Count(&count)
}

//根据用户名获取密码
func GetPasswordbyUsername(username string) (string, error) {
	staff := new(Staffs)
	_, err := engine.Where("english_name = ?", username).Get(staff)
	return staff.Password, err
}

//根据用户名获取id
func GetIdbyUsername(username string) (int64, error) {
	staff := new(Staffs)
	_, err := engine.Where("english_name = ?", username).Get(staff)
	return staff.Id, err

}

func IsExitName(Name string) (bool, error) {
	staff := new(Staffs)
	has, err := engine.Where("name = ?", Name).Exist(staff)
	if err != nil {
		return false, err
	}
	return has, nil
}

func IsExitEnglishName(EnglishName string) (bool, error) {
	staff := new(Staffs)
	has, err := engine.Where("english_name = ?", EnglishName).Exist(staff)
	return has, err
}

//创建用户,返回受影响行数
func CreateUser(staff Staffs) (int64, error) {
	affected, err := engine.Insert(&staff)
	return affected, err
}

//获取所有用户信息
func GetAllStaff() ([]Staffs, error) {
	staffs := make([]Staffs, 0)
	err := engine.Find(&staffs)
	return staffs, err
}

//查询该用户id是否存在
func GetStaffById(id int64) (bool, error) {
	staff := new(Staffs)
	has, err := engine.Where("id = ?", id).Exist(staff)
	return has, err
}

//修改密码
func UpdatePassword(id int64, password string) (int64, error) {
	staff := new(Staffs)
	staff.Password = password
	affected, err := engine.Where("id = ?", id).Cols("password").Update(&staff)
	return affected, err
}

//修改电话
func UpdateTelephone(id int64, telephone string) (int64, error) {
	staff := new(Staffs)
	staff.Telephone = telephone
	affected, err := engine.Id(id).Cols("telephone").Update(&staff)
	return affected, err
}
