package models

import (
	"errors"
	"time"
)

type Outs struct {
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
func InsertOutComponet(order_no string, component_id int64, quantity int64, status int64) (int64, error) {
	out := new(Outs)
	out.OrderNo = order_no
	out.ComponentId = component_id
	out.Quantity = quantity
	out.Status = status
	return engine.InsertOne(out)
}

//插入多个新零件
func InsertOutComponents(order_no string, component_ids []int64, quantity int64, status int64) (int64, error) {
	outs := make([]*Outs, 1)
	for i := 0; i < len(component_ids); i++ {
		outs[i] = new(Outs)
		outs[i].OrderNo = order_no
		outs[i].ComponentId = component_ids[i]
		outs[i].Quantity = quantity
		outs[i].Status = status
	}
	return engine.Insert(outs)
}

//更新某条记录的审核状态
func UpdateOutStatusById(id int64, status int64) error {
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	out := new(Outs)
	out.Status = status
	_, err = session.Where("id = ?", id).Update(out)
	if err != nil {
		session.Rollback()
		return err
	}

	//如果status为已经审核
	if status == 1 {
		out2 := new(Outs)
		//查询该条记录的零件和变更数量
		has, err := session.Where("id = ?", id).Get(out2)
		if err != nil || !has {
			session.Rollback()
			return err

		}
		component := new(Components)
		_, err = session.Where("id = ?", out2.ComponentId).Get(component)
		if err != nil {
			session.Rollback()
			return err
		}
		if component.Quantity < out2.Quantity {
			session.Rollback()
			return errors.New("a pointer to a pointer is not allowed")
		}
		//更改零件表中的库存数量
		sql := "update 'outs' set quantity = quantity - ? where id  = ? and quantity >= ?"
		_, err = session.Exec(sql, out2.Quantity, out2.ComponentId, out2.Quantity)
		if err != nil {
			session.Rollback()
			return err
		}
	}

	return nil

}

//更新某个order的记录审核状态,只更改待审核订单的状态
func UpdateOutStatusByOrderNo(order_no string, status int64) error {
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	//更新当前order中全部的状态
	_, err = session.Table(new(Outs)).Where("order_no = ?", order_no).Update(map[string]interface{}{"status": status})
	if err != nil {
		session.Rollback()
		return err
	}

	//如果status为已经审核
	if status == 1 {
		outs := make([]*Outs, 0)
		//查询该条记录的零件和变更数量
		has, err := session.Where("order_no = ?", order_no).Get(outs)
		if err != nil || !has {
			session.Rollback()
			return err
		}
		//更改零件表中的库存数量，前提是结果数量大于0
		sql := "update 'outs' set quantity = quantity - ? where id  = ? and quantity >= ?"
		for i := 0; i < len(outs); i++ {
			_, err = session.Exec(sql, outs[i].Quantity, outs[i].ComponentId, outs[i].Quantity)
			if err != nil {
				session.Rollback()
				return err
			}

		}
	}

	return nil
}

type ComponentOuts struct {
	Outs       `xorm:"extends"`
	Components `xorm:"extends"`
}

//查询某给单号下的所有零件，包含零件信息
func GetOutByOrderNo(order_no string) ([]ComponentOuts, error) {
	outs := make([]ComponentOuts, 0)
	err := engine.Table("outs").Join("INNER", "components", "components.id = outs.component_id").Where("order_no = ?", order_no).Find(&outs)
	return outs, err
}

func GetOutByOrderNoByStatus(order_no string, status int64) ([]ComponentOuts, error) {
	outs := make([]ComponentOuts, 0)
	err := engine.Table("outs").Join("INNER", "components", "components.id = outs.component_id").Where("order_no = ? and status = ?", order_no, status).Find(&outs)
	return outs, err
}

func DelOutById(id int64) error {
	out := new(Outs)
	_, err := engine.Where("id = ?", id).Get(out)
	if err != nil {
		return err
	}
	//只可以删除未审核的
	if out.Status == 0 {
		_, err = engine.Where("id = ?", id).Delete(new(Outs))
	}
	return err
}

func UpdateOutQuantityById(id int64, quantity int64) error {
	_, err := engine.Table(new(Outs)).Where("id = ?", id).Update(map[string]interface{}{"quantity": quantity})
	return err
}

func GetOutQuantityById(id int64) (int64, error) {
	out := new(Outs)
	_, err := engine.Id(id).Get(out)
	return out.Quantity, err
}
