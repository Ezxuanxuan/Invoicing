package models

import (
	"fmt"
	"github.com/go-xorm/xorm"
)

type OrderNoExplains struct {
	OrderNo   string `xorm:"VARCHAR(30)"`
	OrderType int64  `xorm:"INT(11)"`
	Tag       string `xorm:"VARCHAR(255)"`
}

//创建某单，以及备注
func CreateOrder(No string, Type int64, Tag string) (int64, error) {
	order := new(OrderNoExplains)
	order.OrderNo = No
	order.OrderType = Type
	order.Tag = Tag
	fmt.Println(err)
	return engine.Insert(order)
}

//查询是否存在该单号
func IsExistOrderNo(no string, order_type int64) (bool, error) {
	order := new(OrderNoExplains)
	return engine.Where("order_no = ? and order_type = ?", no, order_type).Exist(order)
}

//连续创建流程单
func CreateAllOrder(no string, tag string) (err error) {
	err = nil
	//创建事务
	session := engine.NewSession()
	defer session.Close()
	//新建采购单
	_, err = CreateOrderSession(no, PURCHASE, tag, session)
	if err != nil {
		session.Rollback()
		return
	}
	//新建投产单
	_, err = CreateOrderSession(no, PRODUCT, tag, session)
	if err != nil {
		session.Rollback()
		return
	}
	//新建质检单
	_, err = CreateOrderSession(no, QUALITY, tag, session)
	if err != nil {
		session.Rollback()
		return
	}
	//新建销毁单
	_, err = CreateOrderSession(no, DESTROY, tag, session)
	if err != nil {
		session.Rollback()
		return
	}
	//新建完成单
	_, err = CreateOrderSession(no, CARRY, tag, session)
	if err != nil {
		session.Rollback()
		return
	}
	//新建入库单
	_, err = CreateOrderSession(no, IN, tag, session)
	if err != nil {
		session.Rollback()
		return
	}
	//新建出库单
	_, err = CreateOrderSession(no, OUT, tag, session)
	if err != nil {
		session.Rollback()
		return
	}
	session.Commit()
	return
}

//创建某单，以及备注
func CreateOrderSession(No string, Type int64, Tag string, session *xorm.Session) (int64, error) {
	order := new(OrderNoExplains)
	order.OrderNo = No
	order.OrderType = Type
	order.Tag = Tag
	return session.Insert(&order)
}

//获取所有订单信息
func GetAllOrder() ([]OrderNoExplains, error) {
	orders := make([]OrderNoExplains, 0)
	err := engine.Find(&orders)
	return orders, err
}
