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
func InsertInComponet(order_no string, component_id int64, quantity int64, status int64) (int64, error) {
	in := new(Ins)
	in.OrderNo = order_no
	in.ComponentId = component_id
	in.Quantity = quantity
	in.Status = status

	return engine.InsertOne(in)
}

//插入多个新零件
func InsertInComponents(order_no string, component_ids []int64, quantity int64, status int64) (int64, error) {
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
func UpdateInStatusById(id int64, status int64) error {
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	in := new(Ins)
	in.Status = status
	_, err = session.Where("id = ?", id).Update(in)
	if err != nil {
		session.Rollback()
		return err
	}

	//如果status为已经审核
	if status == 1 {
		in2 := new(Ins)
		//查询该条记录的零件和变更数量
		has, err := session.Where("id = ?", id).Get(in2)
		if err != nil || !has {
			session.Rollback()
			return err

		}
		//更改零件表中的库存数量
		sql := "update 'ins' set quantity = quantity + ? where id  = ?"
		_, err = session.Exec(sql, in2.Quantity, in2.ComponentId)
		if err != nil {
			session.Rollback()
			return err
		}
	}

	return nil

}

//更新某个order的记录审核状态,只更改待审核订单的状态
func UpdateInStatusByOrderNo(order_no string, status int64) error {
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	//更新当前order中全部的状态
	_, err = session.Table(new(Ins)).Where("order_no = ?", order_no).Update(map[string]interface{}{"status": status})
	if err != nil {
		session.Rollback()
		return err
	}

	//如果status为已经审核
	if status == 1 {
		ins := make([]*Ins, 0)
		//查询该条记录的零件和变更数量
		has, err := session.Where("order_no = ?", order_no).Get(ins)
		if err != nil || !has {
			session.Rollback()
			return err
		}
		//更改零件表中的库存数量
		sql := "update 'ins' set quantity = quantity + ? where id  = ?"
		for i := 0; i < len(ins); i++ {
			_, err = session.Exec(sql, ins[i].Quantity, ins[i].ComponentId)
			if err != nil {
				session.Rollback()
				return err
			}

		}
	}

	return nil
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

func DelInById(id int64) error {
	in := new(Ins)
	_, err := engine.Where("id = ?", id).Get(in)
	if err != nil {
		return err
	}
	//只可以删除未审核的
	if in.Status == 0 {
		_, err = engine.Where("id = ?", id).Delete(new(Ins))
	}
	return err
}

func UpdateInQuantityById(id int64, quantity int64) error {
	_, err := engine.Table(new(Ins)).Where("id = ?", id).Update(map[string]interface{}{"quantity": quantity})
	return err
}

func GetInQuantityById(id int64) (int64, error) {
	in := new(Ins)
	_, err := engine.Id(id).Get(in)
	return in.Quantity, err
}
