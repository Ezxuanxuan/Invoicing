package models

import (
	"time"
)

type Permissions struct {
	Id        int64     `xorm:"not null pk autoincr INT(11)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"index TIMESTAMP"`
	StaffId   int64     `xorm:"INT(11)"`
	Context   string    `xorm:"CHAR(11) default '00000000000'"` //LogPer,InOrderPer,OutOrderPer,PermissionPer,WarePer,ProductPer,DestroyPer,QualityPer,CarryPer,InWarePer,SalePer
}

//创建员工权限
func CreatePermission(permission Permissions) (int64, error) {
	affeced, err := engine.Insert(&permission)
	return affeced, err
}

//初始化用户的权限
func InitPermission(staff_id int64) (int64, error) {
	permission := new(Permissions)
	permission.StaffId = staff_id
	permission.Context = "00000000000"
	return engine.Insert(permission)
}

//更新员工权限
func UpdatePermission(staff_id int64, permission Permissions) (int64, error) {
	affected, err := engine.Where("staff_id = ?", staff_id).Update(&permission)
	return affected, err
}

//通过id获取员工权限
func GetPermissionById(id int64) (string, error) {
	permission := new(Permissions)
	_, err := engine.Where("id = ?", id).Get(&permission)
	return permission.Context, err
}

//通过员工id获取员工权限
func GetPermissionByStaff(staff_id int64) (string, error) {
	permission := new(Permissions)
	_, err := engine.Where("staff_id = ?", staff_id).Get(permission)
	if err != nil {
		return "", err
	}
	return permission.Context, err
}

//获取所有用户权限
func GetPermissions() ([]Permissions, error) {
	permissions := make([]Permissions, 0)
	err := engine.Find(&permissions)
	return permissions, err
}

func IsExistPermissionByStaffId(staff_id int64) (bool, error) {
	return engine.Where("staff_id = ?", staff_id).Exist()
}
