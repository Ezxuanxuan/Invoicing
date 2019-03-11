package models

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
)

type Ins struct {
	Id          int64     `xorm:"not null pk autoincr INT(11)" json:"in_id"`
	CreatedAt   time.Time `xorm:"created"`
	UpdatedAt   time.Time `xorm:"updated"`
	DeletedAt   time.Time `xorm:"index TIMESTAMP"`
	OrderNo     string    `xorm:"VARCHAR(30)"`
	ComponentId int64     `xorm:"INT(11)" json:"component_id"`
	Quantity    int64     `xorm:"INT(11)" json:"in_quantity"`
	Status      int64     `xorm:"INT(11)"`
}

//插入一条新零件
func InsertInComponet(order_no string, component_id int64, quantity int64, status int64) (int64, error) {
	in2 := new(Ins)
	//查看该零件是否已存在
	has, err := engine.Where("order_no = ? and component_id = ?", order_no, component_id).Get(in2)
	if err != nil {
		return 0, err
	}
	//如果该入库单中已经存在该id
	if has {
		in2.Quantity = quantity + in2.Quantity
		_, err := engine.Update(in2)
		if err != nil {
			return 0, err
		}
		return 1, nil
	}
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
	session := engine.NewSession()
	err1 := session.Begin()
	if err1 != nil {
		return 0, err1
	}
	defer session.Close()
	for i := 0; i < len(component_ids); i++ {
		in := new(Ins)
		has, err := session.Where("order_no = ? and component_id = ?", order_no, component_ids[i]).Get(in)
		if err != nil {
			session.Rollback()
			return 0, err
		}
		//如果该入库单中已经存在该id
		if has {
			in.Quantity = quantity + in.Quantity
			_, err := session.Update(in)
			if err != nil {
				session.Rollback()
				return 0, err
			}
		} else {
			ins[i] = new(Ins)
			ins[i].OrderNo = order_no
			ins[i].ComponentId = component_ids[i]
			ins[i].Quantity = quantity
			ins[i].Status = status
		}
	}
	affected, err2 := session.Insert(ins)
	if err2 != nil {
		session.Rollback()
		return 0, err2
	}
	if err3 := session.Commit(); err3 != nil {
		return 0, err3
	}
	return affected, err2
}

//更新某条记录的审核状态
func UpdateInStatusById(id int64, status int64) (bool, error) {
	session := engine.NewSession()
	err := session.Begin()
	defer session.Close()

	in2 := new(Ins)
	//查询该条记录的零件和变更数量
	has, err := session.Where("id = ?", id).Get(in2)
	if err != nil || !has {
		session.Rollback()
		return false, err

	}
	if in2.Status != 0 {
		return false, errors.New("not unverb")
	}
	in2.Status = status
	sql := "update ins set status =? where id =?;"
	_, err = session.Exec(sql, status, id)
	if err != nil {
		fmt.Println("1232")
		session.Rollback()
		return false, err
	}
	//如果status为已经审核
	if status == 1 {
		//更改零件表中的库存数量
		sql := "update components set quantity = quantity+? where id=?;"
		_, err = session.Exec(sql, in2.Quantity, in2.ComponentId)
		if err != nil {
			fmt.Println("000")
			session.Rollback()
			return false, err
		}
	}
	er := session.Commit()
	if er != nil {
		return false, nil
	}
	return true, nil

}

//更新某个order的记录审核状态,只更改待审核订单的状态
func UpdateInStatusByOrderNo(order_no string, status int64) error {
	session := engine.NewSession()
	defer session.Close()
	err := session.Begin()
	//更新当前order中全部的状态
	_, err = session.Table(new(Ins)).Where("order_no = ?",
		order_no).Update(map[string]interface{}{"status": status})
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
		sql := "update components set quantity = quantity + ? where id  = ?"
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
	err := engine.Table("ins").Join("INNER", "components",
		"components.id = ins.component_id").Where("order_no = ?", order_no).Find(&ins)
	return ins, err
}

func GetInByOrderNoByStatus(order_no string, status int64) ([]ComponentIns, error) {
	ins := make([]ComponentIns, 0)
	err := engine.Table("ins").Join("INNER", "components",
		"components.id = ins.component_id").Where("order_no = ? and status = ?",
		order_no, status).Find(&ins)
	return ins, err
}

func DelInById(id int64) (bool, error) {
	in := new(Ins)
	_, err := engine.Where("id = ?", id).Get(in)
	if err != nil {
		return false, err
	}
	//只可以删除未审核的
	if in.Status == 0 && in.OrderNo != "" {
		_, err = engine.Where("id = ?", id).Delete(new(Ins))
		return true, err
	}
	return false, err
}

func UpdateInQuantityById(id int64, quantity int64) (bool, error) {
	in := new(Ins)
	has, err := engine.Where("id = ?", id).Get(in)
	if err != nil {
		return false, err
	}
	if !has {
		return false, errors.New("n")
	}
	if in.Status != 0 {
		return false, nil
	}
	_, err = engine.Table(new(Ins)).Where("id = ?", id).Update(map[string]interface{}{"quantity": quantity})
	return true, err
}

func GetInQuantityById(id int64) (int64, error) {
	in := new(Ins)
	_, err := engine.Id(id).Get(in)
	return in.Quantity, err
}

func IsExistInId(id int64) (bool, error) {
	in := new(Ins)
	return engine.Where("id = ?", id).Exist(in)
}
