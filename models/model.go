package models

import  (
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"fmt"
)

var Omer orm.Ormer

func init(){
	// 注册驱动
	fmt.Println("model init")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456789@rain@/pirain?charset=utf8")

	orm.RegisterModel(new (User))
	orm.RegisterModel(new (Transc))
	Omer = orm.NewOrm()
	orm.Debug = true

	//Qb, _ := orm.NewQueryBuilder("mysql")
}
