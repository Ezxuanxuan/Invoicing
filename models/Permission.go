package models

import "github.com/jinzhu/gorm"

//2.权限表
type Permission struct {
	gorm.Model
	StaffId int    //StaffTable one-to-one
	Context string `gorm:"char:11"` //"{LogPer,InOrderPer,OutOrderPer,PermissionPer,WarePer,ProductPer,DestroyPer,QualityPer,CarryPer,InWarePer,SalePer}"
}

//创建员工权限
func CreatePermission(permission Permission) bool {
	err := db.Create(&permission).Error
	if err != nil {
		return false
	}
	return true
}

//更新员工权限
func UpdatePermission(permission Permission) bool {
	err := db.Where("staff_id = ?", permission.StaffId).Update(permission).Error
	if err != nil {
		return false
	}
	return true
}

//通过id获取员工权限
func GetPermissionById(id int) (string, error) {
	var permission Permission
	err := db.Where("id = ?", id).Select(&permission).Error
	if err != nil {
		return "", err
	}
	return permission.Context, nil
}

//通过员工id获取员工权限
func GetPermissionByStaff(staff_id int) (string, error) {
	var permission Permission
	err := db.Where("staff_id = ?", staff_id).Select(&permission).Error
	if err != nil {
		return "", err
	}
	return permission.Context, nil
}

//获取所有用户权限
func GetPermissions() ([]Permission, error) {
	var permissions []Permission
	err := db.Select(&permissions).Error
	if err != nil {
		return nil, err
	}
	return permissions, nil
}
