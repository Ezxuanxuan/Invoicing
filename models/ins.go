package models

import (
	"time"
)

type Ins struct {
	Id          int64     `xorm:"not null pk autoincr INT(11)"`
	CreatedAt   time.Time `xorm:"created"`
	UpdatedAt   time.Time `xorm:"updated"`
	DeletedAt   time.Time `xorm:"index TIMESTAMP"`
	OrderNo     string    `xorm:"VARCHAR(30)"`
	ComponentId int64     `xorm:"INT(11)"`
	Quantity    int64     `xorm:"INT(11)"`
	Status      int64     `xorm:"INT(11)"`
}

//插入一条新零件
func InsertComponet(order_no string, component_id int64, quantity int64, status int64) (int64, error) {
	in := new(Ins)
	in.OrderNo = order_no
	in.ComponentId = component_id
	in.Quantity = quantity
	in.Status = status

	return engine.InsertOne(in)
}

//插入多个新零件
func InsertComponents(order_no string, component_ids []int64, quantity int64, status int64) (int64, error) {
	ins := make([]*Ins, 1)
	for i := 0; i < len(component_ids); i++ {
		ins[i] = new(Ins)
		ins[i].OrderNo = order_no
		ins[i].ComponentId = component_ids[i]
		ins[i].Quantity = quantity
		ins[i].Status = status
	}
	return engine.Insert(ins)
}

//更新某条记录的审核状态
func UpdateInStatusById(id int64, status int64) (int64, error) {
	in := new(Ins)
	in.Status = status
	return engine.Where("id = ?", id).Update(in)
}

//更新某个order的记录审核状态,智能更改待审核订单的状态
func UpdateInStatusByOrderNo(order_no string, status int64) error {
	sql := "update 'ins' set status =? where order_no = ? and status = 0"
	_, err := engine.Exec(sql, status, order_no)
	return err
}

type ComponentIns struct {
	Ins        `xorm:"extends"`
	Components `xorm:"extends"`
}

//查询某给单号下的所有零件，包含零件信息
func GetInByOrderNo(order_no string) ([]ComponentIns, error) {
	ins := make([]ComponentIns, 0)
	err := engine.Table("ins").Join("INNER", "components", "components.id = ins.component_id").Where("order_no = ?", order_no).Find(&ins)
	return ins, err
}

func GetInByOrderNoByStatus(order_no string, status int64) ([]ComponentIns, error) {
	ins := make([]ComponentIns, 0)
	err := engine.Table("ins").Join("INNER", "components", "components.id = ins.component_id").Where("order_no = ? and status = ?", order_no, status).Find(&ins)
	return ins, err
}
