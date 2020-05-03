package models

import (
	"errors"
	"fmt"
)


type Transc struct {
	Id int64
	Amount float64
	User int64 //充值人的id
	Charge_man string //充值人名
}

//充值，充值金额必须大于0
func ChargeC(charge *Transc)(error)  {

	Omer.Begin()

	amount := charge.Amount

	if amount < 0 {
		Omer.Rollback()
		return errors.New("charge not validate")
	}

	_,err := Omer.Insert(charge)
	if err != nil {
		Omer.Rollback()
		return err
	}


	user,err := QueryUserById(charge.User)
	fmt.Println("query ---ßß")
	if err != nil{
		Omer.Rollback()
		return err
	}

	user.Balance = user.Balance + amount
	fmt.Println("--------update")
	err = UpdateUser(&user)
	if err != nil {
		Omer.Rollback()
		return err
	}

	Omer.Commit()
	return nil
}

func init() {
}



