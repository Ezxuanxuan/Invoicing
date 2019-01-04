package models

import "github.com/jinzhu/gorm"

//7.仓库表
type Ware struct {
	gorm.Model
	ComponentId int
	Quantity    int
}

//根据零件id查询数量
func GetQuantityByComponent(component_id int) (int, error) {
	var ware Ware
	err := db.Where("component_id = ?", component_id).First(&ware).Error
	if err != nil {
		return -1, err
	}
	return ware.Quantity, nil
}

type ReturnWare struct {
	ComponentId int    //零件id
	No          string //零件编码
	Name        string //零件名
	Material    string //材质
	Quality     int    //单品质量
	Quantity    int    //库存数量
}

func GetAllWare() ([]ReturnWare, error) {
	var returnWare []ReturnWare
	err := db.Exec("SELECT components.no, components.name, components.material, wares.component_id, wares.quantity from components  INNER JOIN wares on components.id = wares.component_id ").Scan(&returnWare).Error

	if err != nil {
		return nil, err
	}
	return returnWare, nil
}
