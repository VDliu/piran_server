package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"fmt"
)

//充值记录
type Charge struct {
	Id int64
	Amount float64
	User int64 //充值人的id
	Charge_man string //充值人名
}


//转账记录
type Transition struct {
	Id int64
	FromId int64
	FromInvite string
	ToId int64
	ToInvite string
	Amount float64
	Status int8  //0 ,无状态状态 ，1 发出状态，2领取状态 3.领取中 4 领取完毕状态 5 过期状态 6异常状态
}

//充值接口，充值金额必须大于0
func ChargeC(charge *Charge)(error)  {
	o := orm.NewOrm();

	o.Begin()

	amount := charge.Amount

	if amount < 0 {
		o.Rollback()
		return errors.New("charge not validate")
	}

	_,err := o.Insert(charge)
	if err != nil {
		o.Rollback()
		return err
	}


	user,err := QueryUserById(charge.User)
	if err != nil{
		o.Rollback()
		return err
	}

	user.Balance = user.Balance + amount
	err = UpdateUser(&user,o)
	if err != nil {
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}

//转账功能，一对多 ，一对一
func Transtion(trans []Transition) (error) {
	o := orm.NewOrm()
	fmt.Printf("tt %v -----------------------ooo1",o)
	o.Begin()
	for _,tran := range trans{
		err := P2pTrans(&tran,o)
		if err != nil{
			o.Rollback()
			return err
		}
	}
	o.Commit()
	return nil

}

//一对一
func P2pTrans(tran *Transition,o orm.Ormer) (error) {
	fmt.Printf("tt %v o2######################",o)
	_,err := o.Insert(tran) //save to trans table
	if err != nil {
		fmt.Println("1111")
		return err
	}

	amount := tran.Amount

	if (amount < 0) {
		return errors.New("不能转为负数")
	}

	//从发送人那里减去
	user,err := QueryUserById(tran.FromId)
	if err != nil{
		return err
	}

	if user.Balance < amount {
		return errors.New("余额不足")
	}

	user.Balance = user.Balance - amount
	err = UpdateUser(&user,o)

	//添加给接收者
	user,err = QueryUserById(tran.ToId)
	if err != nil{
		return err
	}

	user.Balance = user.Balance + amount
	err = UpdateUser(&user,o)
	return  err;
}

//红包记录


func init() {
}



